package routes

import (
	"gorm-mysql/database"
	"gorm-mysql/models"

	"github.com/gofiber/fiber/v2"
)

//Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

//AddBook
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&book)

	return c.Status(200).JSON(book)
}

//getBook
func GetBook(c *fiber.Ctx) error {
	books := []models.Book{}

	database.DBConn.First(&books, c.Params("id"))

	return c.Status(200).JSON(books)
}

//AllBooks
func AllBooks(c *fiber.Ctx) error {
	books := []models.Book{}

	database.DBConn.Find(&books)

	return c.Status(200).JSON(books)
}

//Book
func Book(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DBConn.Where("title = ?", title.Title).Find(&book)
	return c.Status(200).JSON(book)
}

//Update
func Update(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Model(&book).Where("title = ?", title.Title).Update("author", title.Author)

	return c.Status(400).JSON("updated")
}

//Delete
func Delete(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DBConn.Where("title = ?", title.Title).Delete(&book)

	return c.Status(200).JSON("deleted")
}
