package routes

import (
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

	database.DB.Db.Create(&book)

	return c.Status(fiber.StatusOK).JSON(book)
}

// AllBooks
func AllBooks(c fiber.Ctx) error {
	books := []models.Book{}
	database.DB.Db.Find(&books)

	return c.Status(fiber.StatusOK).JSON(books)
}

// Book
func Book(c fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.Bind().Body(title); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", title.Title).Find(&book)
	return c.Status(fiber.StatusOK).JSON(book)
}

// Update
func Update(c fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.Bind().Body(title); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.DB.Db.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

	return c.Status(fiber.StatusOK).JSON("updated")
}

// Delete
func Delete(c fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.Bind().Body(title); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", title.Title).Delete(&book)

	return c.Status(fiber.StatusOK).JSON("deleted")
}
