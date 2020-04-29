package app

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber"
)

var app *fiber.App

func init() {
	app = fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Health check ✅")
	})

	group := app.Group("api")

	group.Get("/hello", func(c *fiber.Ctx) {
		c.Send("Hello World 🚀")
	})

	group.Get("/ola", func(c *fiber.Ctx) {
		c.Send("Olá Mundo 🚀")
	})
}

// Start start Fiber app with normal interface
func Start(address interface{}) error {
	return app.Listen(address)
}

// MyCloudFunction Exported http.HandlerFunc to be deployed to as a Cloud Function
func MyCloudFunction(w http.ResponseWriter, r *http.Request) {
	err := CloudFunctionRouteToFiber(app, w, r)
	if err != nil {
		fmt.Fprintf(w, "err : %v", err)
		return
	}
}
