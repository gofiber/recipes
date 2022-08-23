package handler

import (
	"fiber-oauth-google/auth"

	"github.com/gofiber/fiber/v2"
)

// Auth route
func Auth(c *fiber.Ctx) error {
	path := auth.ConfigGoogle()
	url := path.AuthCodeURL("state") // the url which takes to google access
	return c.Redirect(url)

}

// callback route for google after auth
func Callback(c *fiber.Ctx) error {
	token, error := auth.ConfigGoogle().Exchange(c.Context(), c.FormValue("code"))
	if error != nil {
		panic(error)
	}
	email := auth.GetEmail(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{"email": email, "login": true})

}
