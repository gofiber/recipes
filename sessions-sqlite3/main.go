package main

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"log"
	"time"

	"github.com/gofiber/fiber/v3/extractors"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/gofiber/template/html/v2"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var users = map[string]user{}

func main() {
	// Init users (usually you have this data stored in a RDBMS and indexed by an id / uuid)
	users["jj"] = user{Email: "john.joe@example.com", Firstname: "John", Lastname: "Joe"}
	users["mm"] = user{Email: "mary.moe@example.com", Firstname: "Mary", Lastname: "Moe"}
	users["dd"] = user{Email: "dale.doe@example.com", Firstname: "Dale", Lastname: "Doe"}

	// Init SQLite3 database
	db, err := sql.Open("sqlite3", "./fiber.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Storage package can create this table for you at init time
	// but for the purpose of this example I created it manually
	// expanding its structure with an "u" column to better query
	// all user-related sessions.
	query := `CREATE TABLE IF NOT EXISTS sessions (
			  k  VARCHAR(64) PRIMARY KEY NOT NULL DEFAULT '',
		      v  BLOB NOT NULL,
			  e  BIGINT NOT NULL DEFAULT '0',
			  u  TEXT);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// Init sessions store
	storage := sqlite3.New(sqlite3.Config{
		Database:        "./fiber.db",
		Table:           "sessions",
		Reset:           false,
		GCInterval:      10 * time.Second,
		MaxOpenConns:    100,
		MaxIdleConns:    100,
		ConnMaxLifetime: 1 * time.Second,
	})

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{Views: engine})

	// Use session middleware (recommended pattern)
	app.Use(session.New(session.Config{
		Storage:     storage,
		IdleTimeout: 5 * time.Minute,
		Extractor:   extractors.FromCookie("myapp_session"),
	}))

	// Render index page
	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// Handle login API
	app.Post("/api/login", func(c fiber.Ctx) error {
		req := struct {
			UID string `json:"uid"`
		}{}
		if err := c.Bind().Body(&req); err != nil {
			log.Println(err)
		}

		// For semplicity I used directly the user ID passed by the front-end
		// you should instead check email/password here and then get your user ID.

		// Get session from context (automatically managed)
		s := session.FromContext(c)

		// Get user ID from request
		uid := req.UID

		// Check if already logged in as a different user
		if existingUID := s.Get("uid"); existingUID != nil {
			if existingUID.(string) == uid {
				// Already logged in as this user, just return success
				return c.JSON(fiber.Map{"status": "already logged in"})
			}
			// Trying to login as different user - require logout first
			return c.Status(400).JSON(fiber.Map{"error": "Please logout first"})
		}

		// Important: Regenerate session ID to prevent session fixation
		// This changes the session ID while preserving existing data (like cart items)
		if err := s.Regenerate(); err != nil {
			log.Println(err)
			return c.Status(500).JSON(fiber.Map{"error": "Session error"})
		}

		// Get new session ID after regeneration
		sid := s.ID()

		// Save session data (automatically saved when handler returns)
		s.Set("uid", uid)
		s.Set("sid", sid)
		s.Set("ip", c.RequestCtx().RemoteIP().String())
		s.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
		s.Set("ua", string(c.Request().Header.UserAgent()))

		// Save user reference
		stmt, err := db.Prepare(`UPDATE sessions SET u = ? WHERE k = ?`)
		if err != nil {
			log.Println(err)
		}

		_, err = stmt.Exec(uid, sid)
		if err != nil {
			log.Println(err)
		}

		return c.JSON(nil)
	})

	// Handle logout API
	app.Post("/api/logout", func(c fiber.Ctx) error {
		req := struct {
			SID string `json:"sid"`
		}{}
		if err := c.Bind().Body(&req); err != nil {
			log.Println(err)
		}

		// Get current session from context (automatically managed)
		s := session.FromContext(c)

		// Check session ID
		if len(req.SID) > 0 {
			// Get requested session from storage
			data, err := storage.Get(req.SID)
			if err != nil {
				log.Println(err)
			}

			// Decode requested session data
			gd := gob.NewDecoder(bytes.NewBuffer(data))
			dm := make(map[string]interface{})
			if err := gd.Decode(&dm); err != nil {
				log.Println(err)
			}

			// If it belongs to current user destroy requested session
			if s.Get("uid") != nil && s.Get("uid").(string) == dm["uid"] {
				storage.Delete(req.SID)
			}
		} else {
			// Reset clears all data and generates new session ID
			if err := s.Reset(); err != nil {
				log.Println(err)
				return c.Status(500).JSON(fiber.Map{"error": "Session error"})
			}
		}

		return c.JSON(nil)
	})

	// Handle account API
	app.Get("/api/account", func(c fiber.Ctx) error {
		// Get current session from context (automatically managed)
		s := session.FromContext(c)

		// If there is a valid session
		if len(s.Keys()) > 0 {
			type session struct {
				SID    string `json:"sid"`
				IP     string `json:"ip"`
				Login  string `json:"login"`
				Expiry string `json:"expiry"`
				UA     string `json:"ua"`
			}
			type account struct {
				Email     string    `json:"email"`
				Firstname string    `json:"firstname"`
				Lastname  string    `json:"lastname"`
				Session   string    `json:"session"`
				Sessions  []session `json:"sessions"`
			}

			// Get profile info
			sid := s.ID()
			uid := s.Get("uid").(string)
			u := account{
				Email:     users[uid].Email,
				Firstname: users[uid].Firstname,
				Lastname:  users[uid].Lastname,
				Session:   sid,
			}

			// Get sessions list
			rows, err := db.Query(`SELECT v, e FROM sessions WHERE u = ?`, uid)
			if err != nil {
				log.Println(err)
			}

			defer rows.Close()

			// Loop through sessions
			for rows.Next() {
				var (
					data       = []byte{}
					exp  int64 = 0
				)
				if err := rows.Scan(&data, &exp); err != nil {
					log.Println(err)
				}

				// If session isn't expired
				if exp > time.Now().Unix() {
					// Decode session data
					gd := gob.NewDecoder(bytes.NewBuffer(data))
					dm := make(map[string]interface{})
					if err := gd.Decode(&dm); err != nil {
						log.Println(err)
					}

					// Append session
					u.Sessions = append(u.Sessions, session{
						SID:    dm["sid"].(string),
						IP:     dm["ip"].(string),
						Login:  dm["login"].(string),
						Expiry: time.Unix(exp, 0).UTC().String(),
						UA:     dm["ua"].(string),
					})
				}
			}

			return c.JSON(u)
		}

		return c.JSON(nil)
	})

	log.Fatal(app.Listen(":3000"))
}
