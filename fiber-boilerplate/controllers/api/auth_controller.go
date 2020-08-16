package controllers

import (
	"github.com/gofiber/fiber" //nolint:goimports
	"github.com/gookit/validate"
	"github.com/itsursujit/fiber-boilerplate/auth"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/models"
)

func OAuthToken(c *fiber.Ctx) { //nolint:wsl
	var login models.Login
	if err := c.BodyParser(&login); err != nil {
		c.SendStatus(401)
		c.JSON(fiber.Map{
			"error":   true,
			"message": "Invalid Credentials",
		})
		return
	}

	v := validate.Struct(login)
	if !v.Validate() {
		c.SendStatus(401)
		c.JSON(fiber.Map{
			"error":   true,
			"message": v.Errors.All(),
		})
		return
	}
	user, err := login.CheckLogin() //nolint:wsl
	if err != nil {
		c.SendStatus(401)
		c.JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	token, err := auth.Login(c, user.ID, config.AuthConfig.Api_Jwt_Secret) //nolint:wsl
	if err != nil {
		c.SendStatus(401)
		c.JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	c.JSON(fiber.Map{
		"token":      token.Hash,
		"expires_in": token.Expire,
	})
}

func RefreshOauthToken(c *fiber.Ctx) {

}
