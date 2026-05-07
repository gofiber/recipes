package handlers

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// Errors will process all errors returned to fiber
func Errors(file string) fiber.ErrorHandler {
	return func(c fiber.Ctx, err error) error {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).SendFile(file)
	}
}
