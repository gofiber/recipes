package misc

import "github.com/gofiber/fiber/v2"

func NewMiscHandler(miscRoute fiber.Router) {
	miscRoute.Get("", healthCheck)
}

func healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Hello World!",
	})
}
