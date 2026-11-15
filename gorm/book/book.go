package book

import (
	"strconv"

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

// bookID parses the :id route parameter. It must not be passed to GORM as a
// string: GORM treats a non-numeric string as a raw WHERE clause, which would
// let a request such as /book/1=1 inject SQL. The limit is 63 bits because ids are signed
// 64-bit integers in the database, so a larger id cannot exist and cannot be bound as a query
// argument either.
func bookID(c fiber.Ctx) (uint64, error) {
	return strconv.ParseUint(c.Params("id"), 10, 63)
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
	id, err := bookID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid book ID")
	}
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
	id, err := bookID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid book ID")
	}
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
