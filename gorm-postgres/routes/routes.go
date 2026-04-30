package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm-postgres/database"
	"gorm-postgres/models"
)

// Hello
func Hello(c fiber.Ctx) error {
	return c.SendString("fiber")
}

// AddBook
func AddBook(c fiber.Ctx) error {
	book := new(models.Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if result := database.DB.Db.Create(&book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

// AllBooks
func AllBooks(c fiber.Ctx) error {
	books := []models.Book{}

	if result := database.DB.Db.Find(&books); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

// Book returns a single book by ID
func Book(c fiber.Ctx) error {
	book := models.Book{}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	if result := database.DB.Db.First(&book, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

// Update updates a book's fields by ID
func Update(c fiber.Ctx) error {
	book := new(models.Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	if result := database.DB.Db.Model(&models.Book{}).Where("id = ?", id).Updates(book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON("updated")
}

// Delete removes a book by ID
func Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	if result := database.DB.Db.Delete(&models.Book{}, id); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON("deleted")
}
