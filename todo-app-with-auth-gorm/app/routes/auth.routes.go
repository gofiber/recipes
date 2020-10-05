package routes

import (
	"numtostr/gotodo/app/services"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes containes all the auth routes
func AuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", services.Signup)
	r.Post("/login", services.Login)
}
