package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Creates a new Fiber instance.
	app := fiber.New()

	// Prepare a static middleware to serve the built React files.

	// If you serve Single Page Application on "/web", make sure to add basename on BrowserRouter
	// app.Get("/web*", static.New("./web/build"))

	app.Get("/*", static.New("./web/build"))

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.

	// app.Get("/web/**", static.New("./web/build/index.html"))

	app.Get("*", static.New("./web/build/index.html"))

	// Listen to port 8080.
	log.Fatal(app.Listen(":8080"))
}
