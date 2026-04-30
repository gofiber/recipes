package server

import (
	"app/datasources"
	"app/server/handlers"
	"app/server/services"

	"github.com/gofiber/fiber/v3"
)

// NewServer creates a new Fiber app and sets up the routes.
func NewServer(dataSources *datasources.DataSources) *fiber.App {
	app := fiber.New()
	apiRoutes := app.Group("/api")

	booksService := services.NewBooksService(dataSources.DB)

	apiRoutes.Get("/status", func(c fiber.Ctx) error {
		return c.SendString("ok")
	})
	apiRoutes.Get("/v1/books", handlers.GetBooks(booksService))
	apiRoutes.Post("/v1/books", handlers.AddBook(booksService))

	return app
}
