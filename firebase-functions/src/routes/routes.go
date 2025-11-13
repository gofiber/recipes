package routes

import (
	"example.com/GofiberFirebaseBoilerplate/src/database"
	"example.com/GofiberFirebaseBoilerplate/src/repositories"

	"github.com/gofiber/fiber/v3"
)

type Routes struct {
	mainRepository *repositories.MainRepository
}

func New() *Routes {
	db := database.NewConnection()
	return &Routes{mainRepository: &repositories.MainRepository{DB: db}}
}

func (self *Routes) Setup(app *fiber.App) {
	app.Post("message", self.insertMessage)
}

func (self *Routes) insertMessage(c fiber.Ctx) error {
	return c.SendString("ok")
}
