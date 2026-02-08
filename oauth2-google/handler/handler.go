package handler

import (
	"fiber-oauth-google/auth"

	"github.com/gofiber/fiber/v3"
)

// Auth fiber handler
func Auth(c fiber.Ctx) error {
	path := auth.ConfigGoogle()
	url := path.AuthCodeURL("state")
	return c.Redirect().To(url)
}

// Callback to receive google's response
func Callback(c fiber.Ctx) error {
	token, error := auth.ConfigGoogle().Exchange(c.RequestCtx(), c.FormValue("code"))
	if error != nil {
		panic(error)
	}
	email := auth.GetEmail(token.AccessToken)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"email": email, "login": true})
}

// fiber:context-methods migrated
