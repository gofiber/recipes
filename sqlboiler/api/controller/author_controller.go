package controller

import (
	"fiber-sqlboiler/database"
	"fiber-sqlboiler/models"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gofiber/fiber/v3"
)

func GetAuthors(c fiber.Ctx) error {
	authors, err := models.Authors().All(c.Context(), database.DB)
	if err != nil {
		return internalError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(authors)
}

func GetAuthor(c fiber.Ctx) error {
	authorId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid author id")
	}
	author, err := models.FindAuthor(c.Context(), database.DB, authorId)
	if err != nil {
		return lookupFailed(c, err, "author")
	}
	return c.Status(fiber.StatusOK).JSON(author)
}

func NewAuthor(c fiber.Ctx) error {
	author := models.Author{}
	if err := c.Bind().Body(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := author.Insert(c.Context(), database.DB, boil.Infer()); err != nil {
		return internalError(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(author)
}

func DeleteAuthor(c fiber.Ctx) error {
	authorId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid author id")
	}
	author, err := models.FindAuthor(c.Context(), database.DB, authorId)
	if err != nil {
		return lookupFailed(c, err, "author")
	}
	if _, err := author.Delete(c.Context(), database.DB); err != nil {
		if isForeignKeyViolation(err) {
			return c.Status(fiber.StatusConflict).JSON("author still has posts, delete them first")
		}
		return internalError(c, err)
	}
	return c.SendStatus(fiber.StatusOK)
}

func UpdateAuthor(c fiber.Ctx) error {
	authorId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid author id")
	}

	newAuthor := models.Author{}
	if err := c.Bind().Body(&newAuthor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	author, err := models.FindAuthor(c.Context(), database.DB, authorId)
	if err != nil {
		return lookupFailed(c, err, "author")
	}

	author.Name = newAuthor.Name
	if _, err := author.Update(c.Context(), database.DB, boil.Infer()); err != nil {
		return internalError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(author)
}
