// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) {
		// Get first file from form field "document":
		file, err := c.FormFile("document")

		// Check for errors:
		if err == nil {
			// Save file to root directory:
			c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
		}
	})

	// Start server on http://localhost:3000
	app.Listen(3000)
}
