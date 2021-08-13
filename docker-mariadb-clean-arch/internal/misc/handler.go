package misc

import "github.com/gofiber/fiber/v2"

// Represents a new handler.
func NewMiscHandler(miscRoute fiber.Router) {
	miscRoute.Get("", healthCheck)
}

// Check for the health of the API.
func healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Hello World!",
	})
}
