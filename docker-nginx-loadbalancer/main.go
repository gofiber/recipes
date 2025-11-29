package main

import (
	"log"

	"github.com/gofiber/fiber/v3"                    // Importing the fiber package for handling HTTP requests
	"github.com/gofiber/fiber/v3/middleware/cors"    // Middleware for handling Cross-Origin Resource Sharing (CORS)
	"github.com/gofiber/fiber/v3/middleware/favicon" // Middleware for serving favicon
	"github.com/gofiber/fiber/v3/middleware/logger"  // Middleware for logging HTTP requests
	// Package for logging errors
)

func main() {
	app := fiber.New() // Initialize a new Fiber instance
	// register middlewares
	app.Use(favicon.New()) // Use favicon middleware to serve favicon
	app.Use(cors.New())    // Use CORS middleware to allow cross-origin requests
	app.Use(logger.New())  // Use logger middleware to log HTTP requests

	// Define a GET route for the path '/hello'
	app.Get("/hello", func(c fiber.Ctx) error {
		return c.SendString("World!") // Send a response when the route is accessed
	})

	log.Fatal(app.Listen(":5000")) // Start the server on port 5000 and log any errors
}
