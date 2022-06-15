package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Renders the home view
func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

// Renders the about view
func About(c *fiber.Ctx) error {
	return c.Render("about", nil)
}

// Renders the 404 view
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).Render("404", nil)
}
