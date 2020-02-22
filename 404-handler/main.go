// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import "github.com/gofiber/fiber"

func handler() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.Send("This is a dummy route")
	}
}

func main() {
	// Create new Fiber instance
	app := fiber.New()

	app.Get("/demo", handler())
	app.Get("/list", handler())

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404) // => 404 "Not Found"
	})

	// Start server on http://localhost:3000
	app.Listen(3000)
}
