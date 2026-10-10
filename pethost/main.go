package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Pethost finds the port the application listens on, so any port works.
	log.Fatal(app.Listen(":3000"))
}
