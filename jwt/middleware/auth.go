package middleware

import (
	"errors"

	"api-fiber-gorm/config"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
)

// Protected protect routes
func Protected() func(fiber.Ctx) error {
	// The same secret signs the token in handler.Login. Refuse to start without one,
	// rather than falling back to a value that everyone who reads the code knows.
	secret := config.Config("SECRET")
	if secret == "" {
		panic("SECRET environment variable is required")
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: jwtError,
	})
}

func jwtError(c fiber.Ctx, err error) error {
	// No token in the request at all (the extractor found nothing to read).
	if errors.Is(err, extractors.ErrNotFound) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}

	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
