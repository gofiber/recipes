package handler

import "github.com/gofiber/fiber/v3"

// Hello handle api status
func Hello(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello I'm ok!", "data": nil})
}
