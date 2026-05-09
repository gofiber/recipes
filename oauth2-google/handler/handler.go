package handler

import (
	"crypto/rand"
	"encoding/hex"

	"fiber-oauth-google/auth"

	"github.com/gofiber/fiber/v3"
)

// Auth fiber handler
func Auth(c fiber.Ctx) error {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	state := hex.EncodeToString(b)

	// Store state in a cookie to validate in the callback (CSRF protection).
	c.Cookie(&fiber.Cookie{
		Name:     "oauth_state",
		Value:    state,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	url := auth.ConfigGoogle().AuthCodeURL(state)
	return c.Redirect().To(url)
}

// Callback to receive google's response
func Callback(c fiber.Ctx) error {
	// Validate state parameter against cookie to prevent CSRF attacks.
	cookieState := c.Cookies("oauth_state")
	queryState := c.Query("state")
	if cookieState == "" || cookieState != queryState {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "invalid OAuth state"})
	}

	token, err := auth.ConfigGoogle().Exchange(c.Context(), c.FormValue("code"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	email, err := auth.GetEmail(token.AccessToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"email": email, "login": true})
}
