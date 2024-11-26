package internal

import (
	"context"

	"parsley-app/internal/route_handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

type parsleyApplication struct {
	app *fiber.App
}

var _ bootstrap.Application = &parsleyApplication{}

// NewApp Creates the main application service instance. This constructor function gets invoked by Parsley; add parameters for all required services.
func NewApp(app *fiber.App, routeHandlers []route_handlers.RouteHandler) bootstrap.Application {
	// Register RouteHandler services with the resolved Fiber instance.
	for _, routeHandler := range routeHandlers {
		routeHandler.Register(app)
	}

	return &parsleyApplication{
		app: app,
	}
}

// Run The entrypoint for the Parsley application.
func (a *parsleyApplication) Run(_ context.Context) error {
	return a.app.Listen(":5502")
}
