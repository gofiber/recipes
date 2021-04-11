package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterEvilRoutes(evilApp *fiber.App) {
	evilApp.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/evil-examples", fiber.Map{})
	})

	evilApp.Get("/malicious-form", func(c *fiber.Ctx) error {
		return c.Render("views/malicious-form", fiber.Map{})
	})
}
