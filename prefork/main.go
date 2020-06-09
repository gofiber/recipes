// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Enable prefork 🚀
	app.Settings.Prefork = true

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(3000))

	// Run the following command to see all processes sharing port 3000:
	// sudo lsof -i -P -n | grep LISTEN
}

// Handler
func hello(c *fiber.Ctx) {
	c.Send("Hello, World!")
}
