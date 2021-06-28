package app

import (
	"fmt"
    "strings"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func init() {
	app = fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Health check ✅")
	})

	group := app.Group("api")

	group.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World 🚀")
	})

	group.Get("/ola", func(c *fiber.Ctx) error {
		return c.SendString("Olá Mundo 🚀")
	})
}

// Start start Fiber app with normal interface
func Start(addr string) error {
    if -1 == strings.IndexByte(addr, ':') {
        addr = ":" + addr
    }
    
    return app.Listen(addr)
}

// MyCloudFunction Exported http.HandlerFunc to be deployed to as a Cloud Function
func MyCloudFunction(w http.ResponseWriter, r *http.Request) {
	err := CloudFunctionRouteToFiber(app, w, r)
	if err != nil {
		fmt.Fprintf(w, "err : %v", err)
		return
	}
}
