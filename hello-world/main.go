// 🚀 Fiber is an Express.js inspired web framework written in Go with ❤️.
// 📌 Please open an issue if you got suggestions or found a bug!
// 🖥 https://github.com/gofiber/fiber

package main

import "github.com/gofiber/fiber"

// Handler function
func hello(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Create new GET route on path "/hello"
	app.Get("/hello", hello)

	// Start server on http://localhost:3000
	app.Listen(3000)
}
