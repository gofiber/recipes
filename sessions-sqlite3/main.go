package main

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3/extractors"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/gofiber/template/html/v3"

	_ "github.com/mattn/go-sqlite3"
)

// sessionKeyPrefix is how Fiber's session store prefixes a session ID in the storage key,
// which is the "k" column. This recipe maintains the extra "u" column with plain SQL, so it
// has to build the same key. Everything else goes through the store (GetByID, Delete),
// which also takes care of decoding the session data.
const sessionKeyPrefix = "session:"

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
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

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

	store := session.NewStore(session.Config{
		Storage:        storage,
		IdleTimeout:    5 * time.Minute,
		Extractor:      extractors.FromCookie("myapp_session"),
		CookieHTTPOnly: true, // the front end never needs to read the session cookie
	})

	// Create a new engine
	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{Views: engine})

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

		// Get or create session
		s, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
		defer s.Release() // Important: Manual cleanup required

		// If this is a new session
		if s.Fresh() {
			// Get session ID
			sid := s.ID()

			// Get user ID
			uid := req.UID

			// Save session data
			s.Set("uid", uid)
			s.Set("sid", sid)
			s.Set("ip", c.RequestCtx().RemoteIP().String())
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
			} else {
				defer func() { _ = stmt.Close() }()
				_, err = stmt.Exec(uid, sessionKeyPrefix+sid)
				if err != nil {
					log.Println(err)
				}
			}
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

		// Get current session
		s, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
		defer s.Release() // Important: Manual cleanup required

		// Check session ID
		if len(req.SID) > 0 {
			// Only a logged-in user may delete one of their sessions
			uid, ok := s.Get("uid").(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON("invalid session")
			}

			// Get the requested session through the store, which knows how sessions are
			// keyed and encoded in the table
			target, err := store.GetByID(c.Context(), req.SID)
			if err != nil {
				log.Println(err)
			} else {
				defer target.Release()

				// If it belongs to the current user destroy requested session
				if targetUID, _ := target.Get("uid").(string); targetUID == uid {
					if err := store.Delete(c.Context(), req.SID); err != nil {
						log.Println(err)
					}
				}
			}
		} else {
			// Destroy current session
			if err := s.Destroy(); err != nil {
				log.Println(err)
			}
		}

		return c.JSON(nil)
	})

	// Handle account API
	app.Get("/api/account", func(c fiber.Ctx) error {
		// Get current session
		s, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
		defer s.Release() // Important: Manual cleanup required

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
			uid, ok := s.Get("uid").(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON("invalid session")
			}
			u := account{
				Email:     users[uid].Email,
				Firstname: users[uid].Firstname,
				Lastname:  users[uid].Lastname,
				Session:   sid,
				Sessions:  []session{},
			}

			// Get sessions list. The "k" column holds the storage key (the session ID with
			// Fiber's prefix), and the store loads and decodes a session from its ID.
			rows, err := db.Query(`SELECT k, e FROM sessions WHERE u = ?`, uid)
			if err != nil {
				log.Println(err)
				return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
			}

			defer func() { _ = rows.Close() }()

			// Loop through sessions
			for rows.Next() {
				var (
					key string
					exp int64
				)
				if err := rows.Scan(&key, &exp); err != nil {
					log.Println(err)
					continue
				}

				// Skip expired sessions
				if exp <= time.Now().Unix() {
					continue
				}

				other, err := store.GetByID(c.Context(), strings.TrimPrefix(key, sessionKeyPrefix))
				if err != nil {
					log.Println(err)
					continue
				}

				// Checked assertions leave a field empty instead of panicking
				sid, _ := other.Get("sid").(string)
				ip, _ := other.Get("ip").(string)
				login, _ := other.Get("login").(string)
				ua, _ := other.Get("ua").(string)
				other.Release()

				// Append session
				u.Sessions = append(u.Sessions, session{
					SID:    sid,
					IP:     ip,
					Login:  login,
					Expiry: time.Unix(exp, 0).UTC().String(),
					UA:     ua,
				})
			}

			return c.JSON(u)
		}

		return c.JSON(nil)
	})

	log.Fatal(app.Listen(":3000"))
}

// fiber:context-methods migrated
