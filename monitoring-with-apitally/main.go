package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"os"

	apitally "github.com/apitally/apitally-go/fiber"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

type apiKeyInfo struct {
	userID   string
	userName string
	group    string
}

type Book struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

// API keys for demo purposes only
// WARNING: In production applications:
// - Never hardcode API keys in source code
// - Use a secure database and only store hashed keys
var apiKeys = map[string]apiKeyInfo{
	"d7e123f5a2b9c4e8d6a7b2c1f5e9d3a4": {userID: "user1", userName: "Alice", group: "admin"},
	"8f4e2d1c9b7a5f3e2d8c6b4a9f7e3d1c": {userID: "user2", userName: "Bob", group: "developer"},
	"3a9b8c7d6e5f4a2b1c9d8e7f6a5b4c3d": {userID: "user3", userName: "Charlie", group: "reader"},
}

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	hashedKey := sha256.Sum256([]byte(key))

	for apiKey, info := range apiKeys {
		hashedAPIKey := sha256.Sum256([]byte(apiKey))
		if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
			// Set the consumer for Apitally
			consumer := apitally.Consumer{
				Identifier: info.userID,
				Name:       info.userName,
				Group:      info.group,
			}
			apitally.SetConsumer(c, consumer)
			return true, nil
		}
	}

	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func main() {
	app := fiber.New()
	validate := validator.New()

	// Monitoring and request logging with Apitally
	cfg := apitally.NewConfig(os.Getenv("APITALLY_CLIENT_ID"))
	cfg.Env = "dev"
	cfg.RequestLogging.Enabled = true
	cfg.RequestLogging.LogRequestHeaders = true
	cfg.RequestLogging.LogRequestBody = true
	cfg.RequestLogging.LogResponseBody = true

	app.Use(apitally.Middleware(app, cfg))

	// API key authentication
	app.Use(keyauth.New(keyauth.Config{
		KeyLookup:  "header:Authorization",
		AuthScheme: "Bearer",
		Validator:  validateAPIKey,
	}))

	// Routes
	app.Get("/v1/books", func(c *fiber.Ctx) error {
		books := []Book{
			{Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
			{Title: "Clean Code", Author: "Robert C. Martin"},
		}
		return c.JSON(books)
	})

	app.Post("/v1/books", func(c *fiber.Ctx) error {
		// Parse and validate the request body
		var req Book
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if err := validate.Struct(req); err != nil {
			// Capture validation errors in Apitally
			apitally.CaptureValidationError(c, err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// Logic to create a post goes here ...

		return c.Status(fiber.StatusCreated).Send(nil)
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
