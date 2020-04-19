package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	// Use an external setup function in order
	// to configure the app in tests as well
	app := Setup()

	// start the application on http://localhost:3000
	app.Listen(3000)
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	// Initialize a new app
	app := fiber.New()

	// Register the index route with a simple
	// "OK" response. It should return status
	// code 200
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("OK")
	})

	// Return the configured app
	return app
}
