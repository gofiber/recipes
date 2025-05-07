package api

import (
	"aws-ses-sender/config"

	"github.com/gofiber/fiber/v3"
)

// apiKeyAuth API Key Authentication Middleware
// Middleware to check for the API key in the request header
func apiKeyAuth(c fiber.Ctx) error {
	apiKey := c.Get("x-api-key")
	expectedAPIKey := config.GetEnv("API_KEY", "")
	if apiKey != expectedAPIKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	return c.Next()
}
