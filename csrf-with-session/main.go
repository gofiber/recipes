package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the dummy authentication system
type User struct {
	Username string
	Password string
}

// Dummy user database
var users = map[string]User{
	"user1": {Username: "user1"},
	"user2": {Username: "user2"},
}

func main() {
	// In production, run the app on port 443 with TLS enabled
	// or run the app behind a reverse proxy that handles TLS.
	//
	// It is also recommended that the csrf cookie is set to be
	// Secure and HttpOnly and have the SameSite attribute set
	// to Lax or Strict.
	//
	// In this example, we use the "__Host-" prefix for cookie names.
	// This is suggested when your app uses secure connections (TLS).
	// A cookie with this prefix is only accepted if it's secure,
	// comes from a secure source, doesn't have a Domain attribute,
	// and its Path attribute is "/".
	// This makes these cookies "locked" to the domain.
	//
	// See the following for more details:
	// https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html
	//
	// It's recommended to use the "github.com/gofiber/fiber/v2/middleware/helmet"
	// middleware to set headers to help prevent attacks such as XSS, man-in-the-middle,
	// protocol downgrade, cookie hijacking, SSL stripping, clickjacking, etc.

	// Never hardcode passwords in production code
	hashedPassword1, _ := bcrypt.GenerateFromPassword([]byte("password1"), 10)
	hashedPassword2, _ := bcrypt.GenerateFromPassword([]byte("password2"), 10)

	// Used to help prevent timing attacks
	emptyHash, _ := bcrypt.GenerateFromPassword([]byte(""), 10)
	emptyHashString := string(emptyHash)
	users := map[string]User{
		"user1": {Username: "user1", Password: string(hashedPassword1)},
		"user2": {Username: "user2", Password: string(hashedPassword2)},
	}

	// HTML templates
	engine := html.New("./views", ".html")

	// Create a Fiber app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// Initialize a session store
	sessConfig := session.Config{
		Expiration:     30 * time.Minute,        // Expire sessions after 30 minutes of inactivity
		KeyLookup:      "cookie:__Host-session", // Recommended to use the __Host- prefix when serving the app over TLS
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
	}
	store := session.New(sessConfig)

	// CSRF Error handler
	csrfErrorHandler := func(c *fiber.Ctx, err error) error {
		// Log the error so we can track who is trying to perform CSRF attacks
		// customize this to your needs
		fmt.Printf("CSRF Error: %v IP: %v\n", err, c.IP())

		// Don't leak CSRF error info to the client
		return c.Render("error", fiber.Map{
			"Title":     "Error",
			"Error":     "403 Forbidden",
			"ErrorCode": "403",
		})
	}

	// Configure the CSRF middleware
	csrfConfig := csrf.Config{
		Session:        store,
		KeyLookup:      "form:csrf",   // In this example, we will be using a hidden input field to store the CSRF token
		CookieName:     "__Host-csrf", // Recommended to use the __Host- prefix when serving the app over TLS
		CookieSameSite: "Lax",         // Recommended to set this to Lax or Strict
		CookieSecure:   true,          // Recommended to set to true when serving the app over TLS
		CookieHTTPOnly: true,          // Recommended, otherwise if using JS framework recomend: false and KeyLookup: "header:X-CSRF-Token"
		ContextKey:     "csrf",
		ErrorHandler:   csrfErrorHandler,
		Expiration:     30 * time.Minute,
	}
	csrfMiddleware := csrf.New(csrfConfig)

	// Route for the root path
	app.Get("/", func(c *fiber.Ctx) error {
		// render the root page as HTML
		return c.Render("index", fiber.Map{
			"Title": "Index",
		})
	})

	// Route for the login page
	app.Get("/login", csrfMiddleware, func(c *fiber.Ctx) error {
		csrfToken, ok := c.Locals("csrf").(string)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Render("login", fiber.Map{
			"Title": "Login",
			"csrf":  csrfToken,
		})
	})

	// Route for processing the login
	app.Post("/login", csrfMiddleware, func(c *fiber.Ctx) error {
		// Retrieve the submitted form data
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Check if the credentials are valid
		user, exists := users[username]
		var checkPassword string
		if exists {
			checkPassword = user.Password
		} else {
			checkPassword = emptyHashString
		}

		if bcrypt.CompareHashAndPassword([]byte(checkPassword), []byte(password)) != nil {
			// Authentication failed
			csrfToken, ok := c.Locals("csrf").(string)
			if !ok {
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			return c.Render("login", fiber.Map{
				"Title": "Login",
				"csrf":  csrfToken,
				"error": "Invalid credentials",
			})
		}

		// Set a session variable to mark the user as logged in
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if err := session.Reset(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		session.Set("loggedIn", true)
		if err := session.Save(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Redirect to the protected route
		return c.Redirect("/protected")
	})

	// Route for logging out
	app.Get("/logout", func(c *fiber.Ctx) error {
		// Retrieve the session
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Revoke users authentication
		if err := session.Destroy(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Redirect to the login page
		return c.Redirect("/login")
	})

	// Route for the protected content
	app.Get("/protected", csrfMiddleware, func(c *fiber.Ctx) error {
		// Check if the user is logged in
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		loggedIn, _ := session.Get("loggedIn").(bool)

		if !loggedIn {
			// User is not authenticated, redirect to the login page
			return c.Redirect("/login")
		}

		csrfToken, ok := c.Locals("csrf").(string)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Render("protected", fiber.Map{
			"Title": "Protected",
			"csrf":  csrfToken,
		})
	})

	// Route for processing the protected form
	app.Post("/protected", csrfMiddleware, func(c *fiber.Ctx) error {
		// Check if the user is logged in
		session, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		loggedIn, _ := session.Get("loggedIn").(bool)

		if !loggedIn {
			// User is not authenticated, redirect to the login page
			return c.Redirect("/login")
		}

		csrfToken, ok := c.Locals("csrf").(string)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Retrieve the submitted form data
		message := c.FormValue("message")

		return c.Render("protected", fiber.Map{
			"Title":   "Protected",
			"csrf":    csrfToken,
			"message": message,
		})
	})

	certFile := "cert.pem"
	keyFile := "key.pem"

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		fmt.Println("Self-signed certificate not found, generating...")
		if err := generateSelfSignedCert(certFile, keyFile); err != nil {
			panic(err)
		}
		fmt.Println("Self-signed certificate generated successfully")
		fmt.Println("You will need to accept the self-signed certificate in your browser")
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	ln, err := tls.Listen("tcp", "127.0.0.1:8443", config)
	if err != nil {
		panic(err)
	}

	app.Listener(ln)
}

// generateSelfSignedCert generates a self-signed certificate and key
// and saves them to the specified files
//
// This is only for testing purposes and should not be used in production
func generateSelfSignedCert(certFile string, keyFile string) error {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	certOut, err := os.Create(certFile)
	if err != nil {
		return err
	}
	defer certOut.Close()

	_ = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyOut, err := os.Create(keyFile)
	if err != nil {
		return err
	}
	defer keyOut.Close()

	_ = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return nil
}
