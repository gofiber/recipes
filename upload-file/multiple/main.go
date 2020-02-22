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
		// Parse the multipart form:
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form

			// Get all files from "documents" key:
			files := form.File["documents"]
			// => []*multipart.FileHeader

			// Loop through files:
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				// Save the files to disk:
				c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
			}
		}
	})

	// Start server with https/ssl enabled on http://localhost:443
	app.Listen(3000)
}
