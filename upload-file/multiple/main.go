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

	app.Post("/", func(c fiber.Ctx) error {
		// Parse the multipart form:
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		// => *multipart.Form

		// Get all files from "documents" key:
		files := form.File["documents"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			// Sanitize filename to prevent path traversal attacks:
			filename := filepath.Base(file.Filename)
			fmt.Println(filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./%s", filename))
			// Check for errors
			if err != nil {
				return err
			}
		}
		return c.JSON(fiber.Map{"message": "File uploaded successfully"})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
