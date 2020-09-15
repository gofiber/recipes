// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Post("/", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("document")

		// Check for errors:
		if err == nil {
			//Save file inside uploads folder under current working directory:
			c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	
		}
	})

	app.Post("/temp", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("document")

		// Check for errors:
		if err == nil {
			//(uploads_relative)folder must be created before hand:
			//Save file using a relative path:
			c.SaveFile(file, fmt.Sprintf("/tmp/uploads_relative/%s", file.Filename))
		}
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}