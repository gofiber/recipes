package main

import (
	"fmt"
	"log"
	"validation/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Use Cors
	app.Use(cors.New())

	app.Get("/test", func(ctx *fiber.Ctx) error {
		type User struct {
			ID        uint   `validate:"required,omitempty"`
			Firstname string `validate:"required"`
			Password  string `validate:"gte=10"`
		}

		user := User{
			ID:        1,
			Firstname: "Fiber",
			Password:  "FiberPassword123",
		}

		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			fmt.Println(err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		} else {
			return ctx.Status(fiber.StatusOK).JSON("success time")
		}
	})

	// .env Variables validation
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	log.Fatal(app.Listen(config.Config("PORT")))
}
