// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Post("/", func(c *fiber.Ctx) {
		// Get first file from form field "document":
		file, err := c.FormFile("document")

		// Check for errors:
		if err == nil {
			// Save file to root directory:
			c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
		}
	})

	// Start server
	log.Fatal(app.Listen(3000))
}
