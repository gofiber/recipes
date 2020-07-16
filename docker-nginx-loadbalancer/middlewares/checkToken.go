package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber"

	"fiber-docker-nginx/utilities"
)

/*CheckToken : Check the validate of the jwt*/
func CheckToken(c *fiber.Ctx) {
	_, _, _, err := utilities.ProcessToken(c.Get("Authorization"))
	if err != nil {
		c.Send("Error in the jwt !" + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return
	}
	c.Next()
}
