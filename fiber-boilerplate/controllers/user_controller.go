package controllers

import (
	"github.com/gofiber/fiber" //nolint:goimports
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/models"
)

func Index(c *fiber.Ctx) {
	var users []models.User
	DB.Find(&users)   //nolint:wsl
	_ = c.JSON(users) //nolint:errcheck
}
