package routes

import (
	"strconv"

	"gorm-mysql/database"
	"gorm-mysql/models"

	"github.com/gofiber/fiber/v3"
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

	if result := database.DBConn.Create(&book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func GetBook(c fiber.Ctx) error {
	book := models.Book{}

	if result := database.DBConn.First(&book, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

// AllBooks
func AllBooks(c fiber.Ctx) error {
	books := []models.Book{}

	if result := database.DBConn.Find(&books); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

// Update
func Update(c fiber.Ctx) error {
	book := new(models.Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	if result := database.DBConn.Model(&models.Book{}).Where("id = ?", id).Updates(book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON("updated")
}

// Delete
func Delete(c fiber.Ctx) error {
	book := new(models.Book)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	if result := database.DBConn.Where("id = ?", id).Delete(&book); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON("deleted")
}
