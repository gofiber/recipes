package controller

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"fiber-sqlc/database"
	"fiber-sqlc/database/sqlc"

	"github.com/gofiber/fiber/v3"
)

func GetAuthors(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	authors, err := sqlc.New(database.DB).GetAuthors(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(authors)
}

func GetAuthor(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	authorId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	author, err := sqlc.New(database.DB).GetAuthor(ctx, int32(authorId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString("author not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(author)
}

func NewAuthor(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var author sqlc.NewAuthorParams
	if err := c.Bind().Body(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	newAuthor, err := sqlc.New(database.DB).NewAuthor(ctx, author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(newAuthor)
}

func DeleteAuthor(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	authorId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = sqlc.New(database.DB).DeleteAuthor(ctx, int32(authorId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Author successfully deleted")
}

func UpdateAuthor(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	authorId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var author sqlc.UpdateAuthorParams
	author.ID = int32(authorId)
	if err := c.Bind().Body(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updatedAuthor, err := sqlc.New(database.DB).UpdateAuthor(ctx, author)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString("author not found")
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(updatedAuthor)
}
