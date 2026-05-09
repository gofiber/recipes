package main

import (
	"log"

	"validation/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

type User struct {
	ID        uint   `json:"id"        validate:"required"`
	Firstname string `json:"firstname" validate:"required"`
	Password  string `json:"-"         validate:"gte=10"` // gte = Greater than or equal
}

func main() {
	// Load .env variables before starting the server.
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system environment")
	}

	app := fiber.New()
	validate := validator.New()

	// Use Cors
	app.Use(cors.New())

	app.Post("/test", func(ctx fiber.Ctx) error {
		var user User

		if err := ctx.Bind().Body(&user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := validate.Struct(user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(user)
	})

	log.Fatal(app.Listen(config.Config("PORT")))
}
