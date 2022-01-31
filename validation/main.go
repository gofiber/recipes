package main

import (
	"log"
	"validation/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	validate := validator.New() // Create Validate for using.

	// Use Cors
	app.Use(cors.New())

	app.Get("/test", func(ctx *fiber.Ctx) error {
		type User struct {
			ID        uint   `validate:"required,omitempty"`
			Firstname string `validate:"required"`
			Password  string `validate:"gte=10"` // gte = Greater than or equal
		}

		user := User{
			ID:        1,
			Firstname: "Fiber",
			/*
				if you delete Firstname field
				you'll get response like this: Error:Field validation for 'Firstname' failed on the 'required' tag"
			*/
			Password: "FiberPassword123",
			/*
				if you enter "Fiber" in Password
				you'll get response like this: Error:Field validation for 'Password' failed on the 'gte' tag"
			*/
		}

		if err := validate.Struct(user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("success time")
	})

	// .env Variables validation
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	log.Fatal(app.Listen(config.Config("PORT")))
}
