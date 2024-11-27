package main

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
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

	store := session.New(session.Config{
		Storage:    storage,
		Expiration: 5 * time.Minute,
		KeyLookup:  "cookie:myapp_session",
	})

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Render index page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// Handle login API
	app.Post("/api/login", func(c *fiber.Ctx) error {
		req := struct {
			UID string `json:"uid"`
		}{}
		if err := c.BodyParser(&req); err != nil {
			log.Println(err)
		}

		// For semplicity I used directly the user ID passed by the front-end
		// you should instead check email/password here and then get your user ID.

		// Get or create session
		s, _ := store.Get(c)

		// If this is a new session
		if s.Fresh() {
			// Get session ID
			sid := s.ID()

			// Get user ID
			uid := req.UID

			// Save session data
			s.Set("uid", uid)
			s.Set("sid", sid)
			s.Set("ip", c.Context().RemoteIP().String())
			s.Set("login", time.Unix(time.Now().Unix(), 0).UTC().String())
			s.Set("ua", string(c.Request().Header.UserAgent()))

			err := s.Save()
			if err != nil {
				log.Println(err)
			}

			// Save user reference
			stmt, err := db.Prepare(`UPDATE sessions SET u = ? WHERE k = ?`)
			if err != nil {
				log.Println(err)
			}

			_, err = stmt.Exec(uid, sid)
			if err != nil {
				log.Println(err)
			}
		}

		return c.JSON(nil)
	})

	// Handle logout API
	app.Post("/api/logout", func(c *fiber.Ctx) error {
		req := struct {
			SID string `json:"sid"`
		}{}
		if err := c.BodyParser(&req); err != nil {
			log.Println(err)
		}

		// Get current session
		s, _ := store.Get(c)

		// Check session ID
		if len(req.SID) > 0 {
			// Get requested session
			data, err := store.Storage.Get(req.SID)
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
			if s.Get("uid").(string) == dm["uid"] {
				store.Storage.Delete(req.SID)
			}
		} else {
			// Destroy current session
			s.Destroy()
		}

		return c.JSON(nil)
	})

	// Handle account API
	app.Get("/api/account", func(c *fiber.Ctx) error {
		// Get current session
		s, _ := store.Get(c)

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
