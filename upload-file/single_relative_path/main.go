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
		// Save file inside uploads folder under current working directory:
		if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename)); err != nil {
			return err
		}
		return c.JSON(fiber.Map{"message": "File uploaded successfully"})
	})

	app.Post("/temp", func(c fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}
		// Sanitize filename to prevent path traversal attacks:
		filename := filepath.Base(file.Filename)
		// uploads_relative folder must be created beforehand:
		// Save file to a temp uploads directory:
		if err := c.SaveFile(file, fmt.Sprintf("./uploads_relative/%s", filename)); err != nil {
			return err
		}
		return c.JSON(fiber.Map{"message": "File uploaded successfully"})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
