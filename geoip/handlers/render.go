package handlers

import "github.com/gofiber/fiber/v3"

// Render will pass the remove IP value to the template input
func Render() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"IP": c.IP(),
		})
	}
}
