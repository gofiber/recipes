package routes

import (
	"example.com/GofiberFirebaseBoilerplate/src/database"
	"example.com/GofiberFirebaseBoilerplate/src/models"
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

func (r *Routes) Setup(app *fiber.App) {
	app.Post("message", r.insertMessage)
}

func (r *Routes) insertMessage(c fiber.Ctx) error {
	var body models.MessageInputBody
	if err := c.Bind().JSON(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := r.mainRepository.InsertMessage(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": body})
}
