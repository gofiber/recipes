package route_handlers

import "github.com/gofiber/fiber/v2"

// RouteHandler Must be implemented by route handler services.
type RouteHandler interface {
	Register(app *fiber.App)
}
