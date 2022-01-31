package routes

import (
	"clean-architecture/api/handlers"
	"clean-architecture/pkg/book"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.AddBook(service))
	app.Put("/books", handlers.UpdateBook(service))
	app.Delete("/books", handlers.RemoveBook(service))
}
