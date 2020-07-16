package routes

import (
	"github.com/gofiber/fiber"
	database "fiber-docker-nginx/database/user"

	"net/http"
)

/*GetProfile : Get the user profile*/
func GetProfile(c *fiber.Ctx) {
	ID := c.Query("id")

	if len(ID) < 1 {
		c.Send(ID)
		c.Send("Parameter ID is required")
		c.SendStatus(http.StatusBadRequest)
		return
	}

	profile, err := database.SearchProfile(ID)
	if err != nil {
		c.Send("Error Occurred" + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	if err := c.JSON(profile); err != nil {
		c.Status(500).Send(err)
		return
	}
	c.Accepts("application/json")
	c.SendStatus(http.StatusAccepted)
}
