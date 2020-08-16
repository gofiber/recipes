package auth

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/config"
	"github.com/itsursujit/fiber-boilerplate/models"
)

func User(c *fiber.Ctx) (*models.User, error) {
	store := Session.Get(c)
	userID := store.Get("user_id")
	if userID == nil {
		return nil, errors.New("User Not Logged In")
	}
	return models.GetUserById(userID)
}

func UserID(c *fiber.Ctx) uint {
	store := Session.Get(c)
	return store.Get("user_id").(uint)
}

func IsLoggedIn(c *fiber.Ctx) bool {
	store := Session.Get(c)
	userID := store.Get("user_id")
	tokenHash := store.Get("user_token")
	if userID != nil {
		token := c.Cookies("fiber-boilerplate-Token")
		if token == "" {
			c.Cookie(&fiber.Cookie{
				Name:     "fiber-boilerplate-Token",
				Value:    fmt.Sprintf("%s", tokenHash),
				Secure:   false,
				HTTPOnly: true,
			})
		}

		return true
	}
	return false
}

func Login(c *fiber.Ctx, userID uint, secret string) (config.Token, error) {
	store := Session.Get(c)      // get/create new session
	store.Set("user_id", userID) // save to storage
	token, err := config.CreateToken(c, userID, secret)
	if err == nil {
		store.Set("user_token", token.Hash)
		store.Set("token_expiry", token.Expire)
	}
	store.Save()
	return token, err
}

func Logout(c *fiber.Ctx) {
	store := Session.Get(c)
	store.Delete("user_id")
	err := store.Save()
	if err != nil {
		panic(err)
	}
	c.ClearCookie()
	c.Send("You are now logged out.")
	return
}

func AuthCookie(c *fiber.Ctx) {
	IsLoggedIn(c)
	c.Next()
}
