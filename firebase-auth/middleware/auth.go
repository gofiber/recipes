package middleware

import (
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v3"
)

// FirebaseAuth returns a middleware that validates Firebase ID tokens from the
// Authorization: Bearer <token> header and stores the token claims in c.Locals("user").
func FirebaseAuth(authClient *auth.Client) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid Authorization header")
		}

		idToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if idToken == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing bearer token")
		}

		token, err := authClient.VerifyIDToken(c.Context(), idToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
		}

		c.Locals("user", token.Claims)
		return c.Next()
	}
}
