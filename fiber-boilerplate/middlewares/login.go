package middlewares

import (
	"github.com/gofiber/fiber"
	"github.com/gookit/validate"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/auth"
	"github.com/itsursujit/fiber-boilerplate/models"
)

func RedirectToHomePageOnLogin(c *fiber.Ctx) {
	if auth.IsLoggedIn(c) {
		c.Redirect("/")
		return
	}
	c.Next()
}

func ValidateLoginPost(c *fiber.Ctx) {
	var login models.Login
	if err := c.BodyParser(&login); err != nil {
		Flash.WithError(c, fiber.Map{
			"message": err.Error(),
		}).Redirect("/login")
		return
	}
	v := validate.Struct(login)
	if !v.Validate() {
		Flash.WithError(c, fiber.Map{
			"message": v.Errors.One(),
		}).Redirect("/login")
		return
	}
	user, err := login.CheckLogin() //nolint:wsl

	if err != nil {
		Flash.WithError(c, fiber.Map{
			"message": err.Error(),
		}).Redirect("/login")
		return
	}
	c.Locals("user", user)
	c.Next()
}
