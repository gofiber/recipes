package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

//region struct definitions

// User contains the login user information
type User struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

// TransferItem contains the information of the transfer item
type TransferItem struct {
	To     string `json:"to" xml:"to" form:"to"`
	Amount int    `json:"amount" xml:"amount" form:"amount"`
}

//endregion

var sessionStore = session.New()
var csrfActivated = true

func init() {
	sessionStore.RegisterType(fiber.Map{})
	// this mean, csrf is activated
	csrfActivated = len(os.Args) > 1 && os.Args[1] == "withoutCsrf"
}

// Add CSRF protection middleware.
// Should be done AFTER session middleware.
var csrfProtection = csrf.New(csrf.Config{
	// only to control the switch whether csrf is activated or not
	Next: func(c *fiber.Ctx) bool {
		return csrfActivated
	},
	KeyLookup:      "form:_csrf",
	CookieName:     "csrf_",
	CookieSameSite: "Strict",
	Expiration:     1 * time.Hour,
	KeyGenerator:   utils.UUID,
	ContextKey:     "token",
})

// RegisterRoutes registers the routes and middlewares necessary for the server
func RegisterRoutes(app *fiber.App) {
	// Super simple login system.
	// This is not how real login systems should work.
	validLogins := []User{
		{Username: "bob", Password: "test"},
		{Username: "alice", Password: "test"},
	}
	// Simple accounts ledger.
	// This information would normally be stored in a database like MySQL, PostgreSQL, etc.
	var accounts = map[string]int{
		"bob":   500,
		"alice": 500,
	}

	app.Use(recover.New())

	app.Get("/", requireLogin, csrfProtection, func(c *fiber.Ctx) error {
		currSession, err := sessionStore.Get(c)
		if err != nil {
			return err
		}
		sessionUser := currSession.Get("User").(fiber.Map)
		// release the currSession
		err = currSession.Save()
		if err != nil {
			return err
		}

		if sessionUser["Name"] == "" {
			return c.Status(fiber.StatusBadRequest).SendString("User is empty")
		}
		username := sessionUser["Name"].(string)

		return c.Render("views/home", fiber.Map{
			"username":  username,
			"balance":   accounts[username],
			"csrfToken": c.Locals("token"),
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("views/login", fiber.Map{})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		user := &User{}
		err := c.BodyParser(user)
		if err != nil {
			return err
		}

		if user.Username == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Username is required.")
		}

		if user.Password == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Password is required.")
		}

		if !findUser(validLogins, user) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid username or password.")
		}

		// Valid login.
		// Create a new currSession and save their user data in the currSession.
		currSession, err := sessionStore.Get(c)
		defer currSession.Save()
		if err != nil {
			return err
		}
		err = currSession.Regenerate()
		if err != nil {
			return err
		}
		currSession.Set("User", fiber.Map{"Name": user.Username})

		return c.Redirect("/")
	})

	// Funds transfer with HTTP POST request.
	// This is safer than using a GET request, but is still vulnerable to some attack vectors.
	app.Post("/transfer", requireLogin, csrfProtection, func(c *fiber.Ctx) error {
		transfer := &TransferItem{}
		err := c.BodyParser(transfer)
		if err != nil {
			return err
		}

		if transfer.To == "" {
			return c.Status(fiber.StatusBadRequest).SendString("\"To account\" is required.")
		}
		currSession, err := sessionStore.Get(c)
		if err != nil {
			return err
		}
		sessionUser := currSession.Get("User").(fiber.Map)
		// release the currSession
		err = currSession.Save()
		if err != nil {
			return err
		}

		if sessionUser["Name"] == "" {
			return c.Status(fiber.StatusBadRequest).SendString("\"From account\" is required.")
		}
		username := sessionUser["Name"].(string)

		if _, ok := accounts[transfer.To]; !ok {
			return c.Status(fiber.StatusBadRequest).SendString("Cannot transfer funds to non-existent account (\"" + transfer.To + "\").")
		}

		if _, ok := accounts[username]; !ok {
			return c.Status(fiber.StatusBadRequest).SendString("Cannot transfer funds from non-existent account (\"" + transfer.To + "\").")
		}

		if transfer.Amount == 0 {
			return c.Status(fiber.StatusBadRequest).SendString("\"Amount\" is required.")
		}

		fmt.Printf("Transferring funds (%s) from \"%s\" to \"%s\"\n", transfer.Amount, username, transfer.To)

		accounts[transfer.To] = accounts[transfer.To] + transfer.Amount
		accounts[username] = accounts[username] - transfer.Amount

		fmt.Printf("New account balances:\n%+v \n", accounts)
		// Successfully transferred funds.
		// Redirect the sessionUser back to the home page.
		return c.Redirect("/")
	})
}

// Create a helper function to require login for some routes.
func requireLogin(c *fiber.Ctx) error {
	currSession, err := sessionStore.Get(c)
	if err != nil {
		return err
	}
	user := currSession.Get("User")
	defer currSession.Save()

	if user == nil {
		// This request is from a user that is not logged in.
		// Send them to the login page.
		return c.Redirect("/login")
	}

	// If we got this far, the request is from a logged-in user.
	// Continue on to other middleware or routes.
	return c.Next()
}

func findUser(list []User, compareUser *User) bool {
	for _, item := range list {
		if item.Username == compareUser.Username && item.Password == compareUser.Password {
			return true
		}
	}
	return false
}
