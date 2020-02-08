package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Static("./static")

	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
	})

	app.Listen(8080)
}
