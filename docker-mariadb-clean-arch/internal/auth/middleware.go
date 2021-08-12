package auth

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// JWT error message.
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT!",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
		"status":  "error",
		"message": "Invalid or expired JWT!",
	})
}

// Guards a specific endpoint in the API.
func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler:  jwtError,
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS256",
		TokenLookup:   "cookie:jwt",
	})
}
