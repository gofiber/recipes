// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
    "github.com/gofiber/fiber"
    "github.com/gofiber/recover"
)

func main() {
  app := fiber.New()

	app.Use(recover.New(recover.Config{
		// Config is optional
		Handler: func(c *fiber.Ctx, err error) {
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}))
  
  app.Get("/", func(c *fiber.Ctx) {
    panic("Something went wrong!")
  })

  app.Listen(3000)
}
