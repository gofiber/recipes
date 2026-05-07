package book

import (
	"fiber-gorm/database"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(books)
}

func GetBook(c fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString("No Book Found with ID")
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(book)
}

func NewBook(c fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := db.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(book)
}

func DeleteBook(c fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	if err := db.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).SendString("No Book Found with ID")
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if err := db.Delete(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("Book Successfully deleted")
}
