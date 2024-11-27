package route_handlers

import (
	"strconv"

	"parsley-app/internal/services"

	"github.com/gofiber/fiber/v2"
)

type greeterRouteHandler struct {
	greeter services.Greeter
}

const defaultPoliteFlag = "true"

// Register Registers all greeterRouteHandler route handlers.
func (h *greeterRouteHandler) Register(app *fiber.App) {
	app.Get("/say-hello", h.HandleSayHelloRequest)
}

// HandleSayHelloRequest Handles "GET /say-hello" requests.
func (h *greeterRouteHandler) HandleSayHelloRequest(ctx *fiber.Ctx) error {
	name := ctx.Query("name")

	politeFlag := ctx.Query("polite", defaultPoliteFlag)
	polite, err := strconv.ParseBool(politeFlag)
	if err != nil {
		polite = true
	}

	msg := h.greeter.SayHello(name, polite)
	return ctx.Status(fiber.StatusOK).Send([]byte(msg))
}

var _ RouteHandler = &greeterRouteHandler{}

// NewGreeterRouteHandler Activates the route handler for the /say-hello endpoint.
func NewGreeterRouteHandler(greeter services.Greeter) RouteHandler {
	return &greeterRouteHandler{
		greeter: greeter,
	}
}
