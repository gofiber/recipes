package controllers

import (
	"fmt"
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/auth"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/libraries"
	"github.com/itsursujit/fiber-boilerplate/models"
	"time"
)

func RegisterGet(c *fiber.Ctx) {
	if err := c.Render("auth/register", fiber.Map{"Title": "Register"}, "layouts/auth"); err != nil { //nolint:wsl
		panic(err.Error())
	}
}

func RegisterPost(c *fiber.Ctx) {
	register := c.Locals("register").(models.RegisterForm)
	user, err := register.Signup()
	if err != nil {
		_ = c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": true, "message": "Error on register request", "data": err.Error()}) //nolint:errcheck
		return
	}
	store := Session.Get(c)       // get/create new session
	store.Set("user_id", user.ID) // save to storage
	_ = store.Save()

	go SendConfirmationEmail(user.Email, c.BaseURL())
	_ = c.JSON(user)
	c.Redirect("/")
	return
}

func VerifyRegisteredEmail(c *fiber.Ctx) {
	c.Redirect("/")
	return
}

func ResendConfirmEmail(c *fiber.Ctx) {
	user, _ := auth.User(c)
	go SendConfirmationEmail(user.Email, c.BaseURL())
	c.Redirect("/")
	return
}

func SendPasswordResetEmail(email string, baseURL string) {
	resetEmail := fmt.Sprintf("%s-reset-%d", email, time.Now().Unix())
	resetLink := GeneratePasswordResetURL(resetEmail, baseURL)
	htmlBody := config.PrepareHtml("emails/password-reset", fiber.Map{
		"reset_link": resetLink,
	})
	config.Send(email, "You asked to reset? Please click here!", htmlBody, "", "")
}

func RequestPasswordReset(c *fiber.Ctx) {
	if err := c.Render("auth/request-password-reset", fiber.Map{"Title": "Reset Password"}, "layouts/auth"); err != nil { //nolint:wsl
		panic(err.Error())
	}
}

func RequestPasswordResetPost(c *fiber.Ctx) {
	email := c.FormValue("email")
	_, err := models.GetUserByEmail(email)
	if err != nil {
		c.Redirect("/request-password-reset")
		return
	}
	go SendPasswordResetEmail(email, c.BaseURL())
}

func PasswordReset(c *fiber.Ctx) {
	token := c.Query("t")
	if err := c.Render("auth/password-reset", fiber.Map{
		"Title": "Password Reset",
		"Token": token,
	}, "layouts/auth"); err != nil { //nolint:wsl
		panic(err.Error())
	}
}

func PasswordResetPost(c *fiber.Ctx) {
	register := c.Locals("register").(models.RegisterForm)
	email := c.Locals("email").(string)
	user, err := models.GetUserByEmail(email)
	if err != nil {
		c.SendStatus(401)
		c.Send("Invalid Password Reset Token")
		return
	}
	register.ID = user.ID
	_, err = register.ResetPassword()
	if err != nil {
		c.SendStatus(400)
		c.Send("Oops!! Can't update password at the moment")
		return
	}
	c.Redirect("/login")
	return
}

func SendConfirmationEmail(email string, baseURL string) {
	confirmLink := GenerateConfirmURL(email, baseURL)
	htmlBody := config.PrepareHtml("emails/confirm", fiber.Map{
		"confirm_link": confirmLink,
	})
	config.Send(email, "Is it you? Please confirm!", htmlBody, "", "")
}

func GenerateConfirmURL(email string, baseURL string) string {
	token := libraries.Encrypt(email, config.AppConfig.App_Key)
	uri := fmt.Sprintf("%s/do/verify-email?t=%s", baseURL, token)
	return uri
}

func GeneratePasswordResetURL(email string, baseURL string) string {
	token := libraries.Encrypt(email, config.AppConfig.App_Key)
	uri := fmt.Sprintf("%s/reset-password?t=%s", baseURL, token)
	return uri
}
