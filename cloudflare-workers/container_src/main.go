package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	// Create Fiber app
	app := fiber.New(fiber.Config{
		// Production settings
		AppName:     "Cloudflare Container Worker",
		ProxyHeader: fiber.HeaderXForwardedFor,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Root handler
	app.Get("/", func(c fiber.Ctx) error {
		message := os.Getenv("MESSAGE")
		instanceID := os.Getenv("CLOUDFLARE_DEPLOYMENT_ID")

		return c.JSON(fiber.Map{
			"message":     message,
			"instance_id": instanceID,
			"framework":   "gofiber/v3",
			"status":      "ok",
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
