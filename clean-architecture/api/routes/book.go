package routes

import (
	"clean-architecture/pkg/book"
	"clean-architecture/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", getBooks(service))
	app.Post("/books", addBook(service))
	app.Put("/books", updateBook(service))
	app.Delete("/books", removeBook(service))
}

func addBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.InsertBook(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func updateBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.UpdateBook(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func removeBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)
		bookID := requestBody.ID
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		dberr := service.RemoveBook(bookID)
		if dberr != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  false,
			"message": "updated successfully",
		})
	}
}

func getBooks(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchBooks()
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"books":  fetched,
		})
	}
}
