package router

import (
	"fiber-oauth-google/handler"

	"github.com/gofiber/fiber/v2"
)

// setting up the routes for google authentication with gofiber

func Routes(app *fiber.App) {
	app.Get("/", handler.Auth)
	app.Get("/auth/google/callback", handler.Callback)

}
