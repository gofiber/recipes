package routes

import (
	"github.com/gofiber/fiber"
	database "fiber-docker-nginx/database/user"

	"net/http"

	"fiber-docker-nginx/models"
	"fiber-docker-nginx/utilities"
)

/*UpdateProfile : Edit a user profile*/
func UpdateProfile(c *fiber.Ctx) {
	var user models.User
	var isUpdated bool

	if err := c.BodyParser(&user); err != nil {
		c.Send("Bad request" + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}

	isUpdated, err := database.ModifyProfile(user, utilities.UserID)
	if err != nil {
		c.Send("Error occurred" + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	if !isUpdated {
		c.Send("Can't modify the user register")
		c.SendStatus(http.StatusBadRequest)
		return
	}
	c.SendStatus(http.StatusCreated)

}
