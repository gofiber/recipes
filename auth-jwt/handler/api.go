package handler

import "github.com/gofiber/fiber"

// Hello hanlde api status
func Hello(c *fiber.Ctx) {
	c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
