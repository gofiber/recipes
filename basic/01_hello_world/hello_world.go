// 🚀 Fiber is an Express.js inspired web framework written in Go with ❤️.
// 📌 Please open an issue if you got suggestions or found a bug!
// 🖥 https://github.com/gofiber/fiber

// 👤 Authors: Fiber Community (https://fiber.wiki)
// 📚 Docs: https://github.com/gofiber/recipes/blob/master/basic/01_hello_world/README.md

package main

import "github.com/gofiber/fiber"

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Create new GET route on path "/hello"
	app.Get("/hello", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	// Start server on http://localhost:8080
	app.Listen(8080)
}
