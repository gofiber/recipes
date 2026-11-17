package controller

import (
	"context"
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
		return internalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(authors)
}

func GetAuthor(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	authorId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid author id")
	}

	author, err := sqlc.New(database.DB).GetAuthor(ctx, authorId)
	if err != nil {
		return lookupFailed(c, err, "author")
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
		return internalError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(newAuthor)
}

func DeleteAuthor(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	authorId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid author id")
	}

	err = sqlc.New(database.DB).DeleteAuthor(ctx, authorId)
	if err != nil {
		if isForeignKeyViolation(err) {
			return c.Status(fiber.StatusConflict).SendString("author still has posts, delete them first")
		}
		return internalError(c, err)
	}

	return c.Status(fiber.StatusOK).SendString("Author successfully deleted")
}

func UpdateAuthor(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	authorId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid author id")
	}

	var author sqlc.UpdateAuthorParams
	if err := c.Bind().Body(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// Set the id after binding, so an "id" in the request body cannot override the URL.
	author.ID = authorId

	updatedAuthor, err := sqlc.New(database.DB).UpdateAuthor(ctx, author)
	if err != nil {
		return lookupFailed(c, err, "author")
	}

	return c.Status(fiber.StatusOK).JSON(updatedAuthor)
}
