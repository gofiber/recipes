package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/mustache/v2"
)

func main() {
	engineXML := mustache.New("./xmls", ".xml")
	if err := engineXML.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/rss", func(c *fiber.Ctx) error {
		// Set Content-Type to application/rss+xml
		c.Type("rss")

		// Set rendered template to body
		return engineXML.Render(c, "example", fiber.Map{
			"Lang":      "en",
			"Title":     "hello-rss",
			"Greetings": "Hello World",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
