// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Fiber instance
	app := fiber.New(fiber.Config{
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	c.Status(fiber.StatusInternalServerError)
		// 	return c.SendString(err.Error())
		// },
	})

	// Middleware
	app.Use(recover.New())

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
	panic("No worries, I won't crash! 🙏")
}
