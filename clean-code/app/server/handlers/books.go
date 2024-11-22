package handlers

import (
	"log"

	"app/server/domain"
	"app/server/services"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(service services.BooksService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		books, err := service.GetBooks(c.UserContext())
		if err != nil {
			log.Printf("GetBooks failed: %v", err)
			return sendError(c, fiber.StatusInternalServerError, "internal error")
		}

		return c.JSON(domain.BooksResponse{
			Books: books,
		})
	}
}

func AddBook(service services.BooksService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var book domain.Book
		if err := c.BodyParser(&book); err != nil {
			log.Printf("AddBook request parsing failed: %v", err)
			return sendError(c, fiber.StatusBadRequest, "invalid request")
		}

		err := service.SaveBook(c.UserContext(), book)
		if err != nil {
			log.Printf("AddBook failed: %v", err)
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
