package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
	"github.com/zeimedee/go-postgres/models"
)

// Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

// AddBook
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	database.DB.Db.Create(&book)

	return c.Status(http.StatusOK).JSON(book)
}

// AllBooks
func AllBooks(c *fiber.Ctx) error {
	books := []models.Book{}
	database.DB.Db.Find(&books)

	return c.Status(http.StatusOK).JSON(books)
}

// Book
func Book(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", title.Title).Find(&book)
	return c.Status(http.StatusOK).JSON(book)
}

// Update
func Update(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	database.DB.Db.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

	return c.Status(http.StatusOK).JSON("updated")
}

// Delete
func Delete(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", title.Title).Delete(&book)

	return c.Status(http.StatusOK).JSON("deleted")
}
