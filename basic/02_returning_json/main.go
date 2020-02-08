package main

import "github.com/gofiber/fiber"

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := fiber.New()

	app.Get("/json", func(c *fiber.Ctx) {
		data := Data{
			Name: "John",
			Age:  20,
		}
		err := c.JSON(data)
		if err != nil {
			c.SendStatus(500)
		}
	})

	app.Listen(8080)
}
