package handlers

import (
	"log/slog"

	"app/server/domain"
	"app/server/services"

	"github.com/gofiber/fiber/v2"
)

// GetBooks returns a handler function that retrieves all books
func GetBooks(service services.BooksService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		books, err := service.GetBooks(c.UserContext())
		if err != nil {
			slog.Error("GetBooks failed", "error", err)
			return sendError(c, fiber.StatusInternalServerError, "internal error")
		}

		return c.JSON(domain.BooksResponse{
			Books: books,
		})
	}
}

// AddBook returns a handler function that adds a book
func AddBook(service services.BooksService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var book domain.Book
		if err := c.BodyParser(&book); err != nil {
			slog.Warn("AddBook request parsing failed", "error", err)
			return sendError(c, fiber.StatusBadRequest, "invalid request")
		}
		// For production use add proper validation here

		err := service.SaveBook(c.UserContext(), book)
		if err != nil {
			slog.Error("AddBook failed", "error", err)
			return sendError(c, fiber.StatusInternalServerError, "internal error")
		}
		return c.SendStatus(fiber.StatusCreated)
	}
}

func sendError(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(domain.ErrorResponse{
		Error: message,
	})
}
