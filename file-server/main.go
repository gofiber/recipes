// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Static file server
	app.Static("/", "./files")
	// => http://localhost:3000/hello.txt
	// => http://localhost:3000/gopher.gif

	// Start server
	log.Fatal(app.Listen(":3000"))
}
