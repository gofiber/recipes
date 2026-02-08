package book

import (
	"fiber-gorm/database"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(fiber.StatusNotFound).SendString("No Book Found with ID")
	}
	db.Delete(&book)
	return c.SendString("Book Successfully deleted")
}
