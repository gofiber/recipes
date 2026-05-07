package routes

import (
	"time"

	"numtostr/gotodo/app/services"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

// AuthRoutes containes all the auth routes
func AuthRoutes(app fiber.Router) {
	authLimiter := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
	})

	r := app.Group("/auth")
	r.Use(authLimiter)

	r.Post("/signup", services.Signup)
	r.Post("/login", services.Login)
}
