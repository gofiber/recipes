package router

import (
	"fiber-oauth-google/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// setting up the routes for google authentication with gofiber

func Routes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Auth)
	api.Get("/auth/google/callback", handler.Callback)

}
