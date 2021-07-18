package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Creates a new Fiber instance.
	app := fiber.New()

	// Prepare a static middleware to serve the built React files.
	app.Static("/", "./web/build")

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.
	app.Static("*", "./web/build/index.html")

	// Listen to port 8080.
	log.Fatal(app.Listen(":8080"))
}
