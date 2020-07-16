package routes

import (
	"net/http"

	"github.com/gofiber/fiber"
	user2 "fiber-docker-nginx/database/user""

	"fiber-docker-nginx/models"
)

/*Register : Save a user in the database*/
func Register(c *fiber.Ctx) {

	var user models.User
	if err := c.BodyParser(&user); err !=nil{
		c.Send("Bad request "+err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	/*Validation*/
	err:=user.Validate()
	if err != nil{
		c.Send(err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	_, found, _ := user2.UserExist(user.Email)
	if found {
		c.Send("Email already register")
		c.SendStatus(http.StatusBadRequest)
		return
	}
	_, isCreated, err := user2.UserRegister(user)
	if err != nil {
		c.Send("Error at moment to register in the db "+err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	if !isCreated {
		c.Send("Can't insert the new user")
		c.SendStatus(http.StatusBadRequest)
		return
	}

	c.SendStatus(http.StatusCreated)
}
