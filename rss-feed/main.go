package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/mustache/v3"
)

func main() {
	engineXML := mustache.New("./xmls", ".xml")
	if err := engineXML.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/rss", func(c fiber.Ctx) error {
		// Set Content-Type to application/rss+xml
		c.Type("rss")

		published := time.Date(2024, time.January, 2, 15, 4, 5, 0, time.UTC)

		// Set rendered template to body. {{Name}} in the template escapes &, <, > and
		// quotes, so titles and descriptions cannot break the XML. ({{{Name}}} would not.)
		return engineXML.Render(c, "example", fiber.Map{
			"Lang":        "en",
			"Title":       "Fiber recipes & news",
			"Link":        "http://localhost:3000/",
			"Description": "An example RSS 2.0 feed rendered with Mustache",
			"Items": []fiber.Map{
				{
					"Title":       "Hello, RSS",
					"Link":        "http://localhost:3000/posts/1",
					"Description": "The first <post> in the feed",
					"PubDate":     published.Format(time.RFC1123Z),
				},
				{
					"Title":       "Second item",
					"Link":        "http://localhost:3000/posts/2",
					"Description": "Another entry",
					"PubDate":     published.Add(24 * time.Hour).Format(time.RFC1123Z),
				},
			},
		})
	})

	log.Fatal(app.Listen(":3000"))
}
