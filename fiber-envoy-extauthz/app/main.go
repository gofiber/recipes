package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Healthy")
	})

	api := app.Group("/api")

	api.Get("/resource", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Some Resource API")
	})

	log.Fatal(app.Listen(":3000"))
}
