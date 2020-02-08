package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) {
		if form, err := c.MultipartForm(); err != nil {
			files := form.File["documents"]
			for _, file := range files {
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"
				c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
				// Saves the file to disk
			}
		}
	})

	app.Listen(8080)
}
