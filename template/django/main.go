package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

func main() {

	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		log.Println(err)
	}
	cwd := filepath.Dir(path)

	// Create a new engine
	engine := django.New(cwd+"/views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".django"))

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render with and extends
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/embed", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("embed", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main2")
	})

	log.Fatal(app.Listen(":3000"))
}
