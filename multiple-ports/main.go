// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Create new GET route on path "/hello"
	app.Get("/hello", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	// Listen on port 8080
	go func() {
		app.Listen(8080)
	}()

	// Shortcut for listening on port 8081
	go app.Listen(8081)

	// Listen on port 3000
	app.Listen(3000)
}
