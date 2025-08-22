// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// serve Single Page application on "/web"
	// assume static file at dist folder
	app.Get("/web*", static.New("dist"))

	app.Get("/web/*", func(ctx fiber.Ctx) error {
		return ctx.SendFile("./dist/index.html")
	})

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
