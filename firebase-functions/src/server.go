package src

import (
	"example.com/GofiberFirebaseBoilerplate/src/routes"

	"github.com/gofiber/fiber/v3"
)

func CreateServer() *fiber.App {
	version := "v1.0.0"

	app := fiber.New(fiber.Config{
		ServerHeader: "Gofiber Firebase Boilerplate",
		AppName:      "Gofiber Firebase Boilerplate " + version,
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Gofiber Firebase Boilerplate " + version)
	})

	routes.New().Setup(app)

	return app
}
