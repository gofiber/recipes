package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: os.Getenv("SECRET")},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			status := fiber.StatusUnauthorized
			message := "Invalid or expired JWT"

			if err.Error() == "Missing or malformed JWT" {
				status = fiber.StatusBadRequest
				message = "Missing or malformed JWT"
			}

			return c.Status(status).JSON(fiber.Map{
				"status":  "error",
				"message": message,
				"data":    nil,
			})
		},
	})
}
