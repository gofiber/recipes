package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/itsursujit/fiber-boilerplate/auth"
)

func Landing(c *fiber.Ctx) {
	user, _ := auth.User(c)
	layout := "layouts/main"
	view := "index"
	if user == nil {
		layout = "layouts/landing"
		view = "landing"
	}

	if err := c.Render(view, fiber.Map{
		"auth": user != nil,
		"user": user,
	}, layout); err != nil {
		panic(err.Error())
	}
}
