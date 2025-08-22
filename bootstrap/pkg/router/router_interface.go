package router

import "github.com/gofiber/fiber/v3"

type Router interface {
	InstallRouter(app *fiber.App)
}
