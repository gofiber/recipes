package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello Heroku")
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
		log.Print("$PORT == 3000")
	}

	app.Listen(port)
}
