package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func main() {
	// Initialize the application
	app := fiber.New()

	// Hello, World!
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	// Listen and Server in 0.0.0.0:$PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	err := app.Listen(":" + port)
	log.Panic(err)
}
