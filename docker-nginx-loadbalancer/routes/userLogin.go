package routes

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber"
	user2 "fiber-docker-nginx/database/user""

	"fiber-docker-nginx/jwt"
	"fiber-docker-nginx/models"
)

/*Login : user try to login in the app*/
func Login(c *fiber.Ctx) {
	c.Accepts("application/json")

	var user models.User
	if err := c.BodyParser(&user); err !=nil{
		c.Send("User or password invalid"+err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}

	if len(user.Email) == 0 {
		c.Send("Email is required")
		c.SendStatus(http.StatusBadRequest)
		return
	}

	document, exist := user2.TryLogin(user.Email, user.Password)
	if !exist {
		c.Send("User or password invalid")
		c.SendStatus(http.StatusBadRequest)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		c.Send("Error occurred"+err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	res := models.LoginResponse{
		Token: jwtKey,
	}


	c.Accepts("application/json")
	c.SendStatus(http.StatusCreated)

	if err := c.JSON(res); err != nil {
		c.Status(500).Send(err)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = jwtKey
	cookie.Expires = expirationTime
	c.Cookie(cookie)

}
