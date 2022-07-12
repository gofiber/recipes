package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	return c.Status(200).JSON("Welcome to Fiber!")
}
