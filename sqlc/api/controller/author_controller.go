package controller

import (
	"context"
	"strconv"
	"time"

	"fiber-sqlc/database"
	"fiber-sqlc/database/sqlc"

	"github.com/gofiber/fiber/v2"
)

func GetAuthors(c *fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	authors, err := sqlc.New(database.DB).GetAuthors(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(authors)
}

func GetAuthor(c *fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	author, err := sqlc.New(database.DB).GetAuthor(ctx, int32(authorId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(author)
}

func NewAuthor(c *fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	var author sqlc.NewAuthorParams
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	newAuthor, err := sqlc.New(database.DB).NewAuthor(ctx, author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(newAuthor)
}

func DeleteAuthor(c *fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = sqlc.New(database.DB).DeleteAuthor(ctx, int32(authorId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Author successfully deleted")
}

func UpdateAuthor(c *fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var author sqlc.UpdateAuthorParams
	author.ID = int32(authorId)
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updatedAuthor, err := sqlc.New(database.DB).UpdateAuthor(ctx, author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(updatedAuthor)
}
