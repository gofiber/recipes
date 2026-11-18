// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
	"log"

	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Serve Single Page Application on "/web".
	// A path that is not a file in dist, such as the client-side route /web/about,
	// gets index.html so client-side routing works. The static middleware has set
	// the status to 404 by then, so reset it to 200.
	app.Get("/web*", static.New("dist", static.Config{
		NotFoundHandler: func(ctx fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).SendFile("./dist/index.html")
		},
	}))

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
