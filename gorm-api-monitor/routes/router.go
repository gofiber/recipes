package routes

import (
	"robot-monitoreo/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {

	// Middleware
	api := app.Group("/api")

	//Index endpoint
	api.Get("/", handlers.Welcome)
	api.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))

	//Dogs
	api.Get("/dogs", handlers.GetDogs)
	api.Get("/dogs/:id", handlers.GetDog)
	api.Post("/dogs", handlers.AddDog)
	api.Put("/dogs/:id", handlers.UpdateDog)
	api.Delete("/dogs/:id", handlers.RemoveDog)

}
