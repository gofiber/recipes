package handler

import (
	"fiber-oauth-google/auth"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	path := auth.ConfigGoogle()
	url := path.AuthCodeURL("state")
	return c.Redirect(url)

}

func Callback(c *fiber.Ctx) error {
	token, error := auth.ConfigGoogle().Exchange(c.Context(), c.FormValue("code"))
	if error != nil {
		panic(error)
	}
	email := auth.GetEmail(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{"email": email, "login": true})

}
