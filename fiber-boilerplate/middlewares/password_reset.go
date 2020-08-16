package middlewares

import (
	"errors"
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/libraries"
	"strconv"
	"strings"
	"time"
)

func ValidatePasswordReset(c *fiber.Ctx) {
	token := c.Query("t")
	err := _validatePasswordReset(c, token)
	if err != nil {
		Flash.WithError(c, fiber.Map{
			"message": err.Error(),
		}).Redirect("/login")
		return
	}
	c.Next()
}

func ValidatePasswordResetPost(c *fiber.Ctx) {
	token := c.Params("token")
	err := _validatePasswordReset(c, token)
	if err != nil {
		Flash.WithError(c, fiber.Map{
			"message": err.Error(),
		}).Redirect("/login")
	}
	c.Next()
}

func _validatePasswordReset(c *fiber.Ctx, t string) error {
	t = libraries.Decrypt(t, config.AppConfig.App_Key)
	emailParts := strings.Split(t, "-reset-")
	if len(emailParts) != 2 {
		return errors.New("Invalid Password Reset Token")
	}

	tokenTS, err := strconv.ParseInt(emailParts[1], 10, 64)
	if err != nil {
		return errors.New("Invalid Password Reset Token")
	}
	now := time.Now().Unix()
	diff := now - tokenTS
	if diff > (5 * 60) {
		return errors.New("Password Reset Token has expired!")
	} else if diff < 0 {
		return errors.New("Invalid Password Reset Token")
	}
	c.Locals("email", emailParts[0])
	c.Locals("token", t)
	return nil
}
