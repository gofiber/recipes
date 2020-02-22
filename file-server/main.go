// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import "github.com/gofiber/fiber"

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Serve files from "files" directory
	app.Static("/", "./files")
	// http://localhost:3000/hello.txt
	// http://localhost:3000/gopher.gif

	// Start server on http://localhost:3000
	app.Listen(3000)
}
