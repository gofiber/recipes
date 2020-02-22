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

	// Enable prefork
	app.Settings.Prefork = true

	// Create new GET route on path "/hello"
	app.Get("/hello", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	// Listen on port 3000
	app.Listen(3000)

	// Run the following command to see all processes sharing port 3000
	// sudo lsof -i -P -n | grep LISTEN
}
