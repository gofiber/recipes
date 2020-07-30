package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/django"
)

func main() {
	// Create a new engine
	engine := django.New("./views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".django"))

	// Pass the engine to the Views
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) {
		// Render with and extends
		c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/embed", func(c *fiber.Ctx) {
		// Render index within layouts/main
		c.Render("embed", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main2")
	})

	app.Listen(3000)
}
