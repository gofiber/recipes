package controllers

import (
	"github.com/gofiber/fiber" //nolint:goimports
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/auth"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/models"
)

func LoginGet(c *fiber.Ctx) {
	Flash.Get(c)
	if err := c.Render("auth/login", Flash.Data, "layouts/auth"); err != nil { //nolint:wsl
		panic(err.Error())
	}
}

func LoginPost(c *fiber.Ctx) { //nolint:wsl
	user := c.Locals("user").(*models.User)
	auth.Login(c, user.ID, config.AuthConfig.App_Jwt_Secret) //nolint:wsl
	c.Redirect("/")
	return
}

func LogoutPost(c *fiber.Ctx) { //nolint:nolintlint,wsl
	if auth.IsLoggedIn(c) {
		auth.Logout(c)
	}
	c.Redirect("/login")
	return
}
