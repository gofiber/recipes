package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Home renders the home view
func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

// About renders the about view
func About(c *fiber.Ctx) error {
	return c.Render("about", nil)
}

// NoutFound renders the 404 view
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).Render("404", nil)
}
