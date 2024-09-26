package routes

import (
	"gorm-mysql/database"
	"gorm-mysql/models"
	"strconv"

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

//Update
func Update(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Model(&models.Book{}).Where("id = ?", id).Update("title", book.Title)

	return c.Status(200).JSON("updated")
}

//Delete
func Delete(c *fiber.Ctx) error {
	book := new(models.Book)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Where("id = ?", id).Delete(&book)

	return c.Status(200).JSON("deleted")
}
