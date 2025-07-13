// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io
package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// Listen on port 8080
	go func() {
		log.Fatal(app.Listen(":8080"))
	}()

	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
