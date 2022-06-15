package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func About(c *fiber.Ctx) error {
	return c.Render("about", nil)
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).Render("404", nil)
}
