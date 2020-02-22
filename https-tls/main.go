// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import "github.com/gofiber/fiber"

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Create new GET route on path "/hello"
	app.Get("/protocol", func(c *fiber.Ctx) {
		c.Send(c.Protocol())
		// => https
	})

	// Start server with https/ssl enabled on http://localhost:443
	app.Listen(443, "./certs/ssl.cert", "./certs/ssl.key")
}
