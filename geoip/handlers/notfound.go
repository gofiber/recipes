package handlers

import "github.com/gofiber/fiber/v3"

// NotFound returns status code 404 along with the given html file
func NotFound(file string) fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile(file)
	}
}
