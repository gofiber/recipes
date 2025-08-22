package app

import (
	"net/http"
	"strings"

	"example.com/GofiberFirebaseBoilerplate/src"
	"github.com/gofiber/fiber/v3"
)

var app *fiber.App

func init() {
	app = src.CreateServer()
}

// Start start Fiber app with normal interface
func Start(addr string) error {
	if -1 == strings.IndexByte(addr, ':') {
		addr = ":" + addr
	}

	return app.Listen(addr)
}

// ServerFunction Exported http.HandlerFunc to be deployed to as a Cloud Function
func ServerFunction(w http.ResponseWriter, r *http.Request) {
	CloudFunctionRouteToFiber(app, w, r)
}
