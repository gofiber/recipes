// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Fiber instance
	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})

	// Routes
	app.Post("/", func(c fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}
		// Sanitize filename to prevent path traversal attacks:
		filename := filepath.Base(file.Filename)
		// Save file to root directory:
		if err := c.SaveFile(file, fmt.Sprintf("./%s", filename)); err != nil {
			return err
		}
		return c.JSON(fiber.Map{"message": "File uploaded successfully"})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
