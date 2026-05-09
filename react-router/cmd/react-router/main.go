package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	// Creates a new Fiber instance.
	app := fiber.New()

	// Prepare a static middleware to serve the built React files.
	app.Get("/*", static.New("./web/build"))

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.
	app.Get("*", static.New("./web/build/index.html"))

	// Listen to port from PORT env variable, defaulting to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
