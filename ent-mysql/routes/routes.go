package routes

import (
	"context"
	"log"

	"ent-mysql/database"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func GetBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	b, err := database.DBConn.Book.
		Get(ctx, id)
	if b == nil {
		return c.Status(fiber.StatusNotFound).JSON("Not found")
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book data: ", b)
	return c.Status(fiber.StatusOK).JSON(b)
}

func GetAllBook(c *fiber.Ctx) error {
	b, err := database.DBConn.Book.
		Query().All(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book data All: ", b)
	return c.Status(fiber.StatusOK).JSON(b)
}

func CreateBook(c *fiber.Ctx) error {
	title := c.Query("title")
	author := c.Query("author")
	if title == "" || author == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Not enough data for creating")
	}
	b, err := database.DBConn.Book.
		Create().
		SetTitle(title).
		SetAuthor(author).
		Save(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book was created: ", b)
	return c.Status(fiber.StatusOK).JSON(b)
}

func DeleteBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := database.DBConn.Book.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book was deleted")
	return c.Status(fiber.StatusOK).JSON(nil)
}

func UpdateBook(c *fiber.Ctx) error {
	title := c.Query("title")
	author := c.Query("author")
	id, _ := c.ParamsInt("id")
	if title == "" || author == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Not enough data for updating")
	}
	b, err := database.DBConn.Book.
		UpdateOneID(id).
		SetTitle(title).
		SetAuthor(author).
		Save(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	log.Println("book was updated: ", b)
	return c.Status(fiber.StatusOK).JSON(b)
}
