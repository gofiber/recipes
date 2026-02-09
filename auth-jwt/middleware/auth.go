package middleware

import (
	"os"
	"strings"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
)

// Protected protect routes
func Protected() fiber.Handler {
	secret := os.Getenv("SECRET")
	if secret == "" {
		panic("SECRET environment variable is required")
	}

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(c fiber.Ctx, err error) error {
			status := fiber.StatusUnauthorized
			message := "Invalid or expired JWT"

			if strings.Contains(strings.ToLower(err.Error()), "missing or malformed jwt") {
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
