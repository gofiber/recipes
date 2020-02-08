package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send(c.Protocol()) // => "https"
	})

	app.Listen(443, "server.crt", "server.key")
}
