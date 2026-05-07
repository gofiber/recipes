package controller

import (
	"strconv"

	"fiber-sqlboiler/database"
	"fiber-sqlboiler/models"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gofiber/fiber/v3"
)

func GetAuthors(c fiber.Ctx) error {
	authors, err := models.Authors().All(c.Context(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(authors)
}

func GetAuthor(c fiber.Ctx) error {
	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	author, err := models.FindAuthor(c.Context(), database.DB, authorId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(author)
}

func NewAuthor(c fiber.Ctx) error {
	author := models.Author{}
	if err := c.Bind().Body(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := author.Insert(c.Context(), database.DB, boil.Infer()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(author)
}

func DeleteAuthor(c fiber.Ctx) error {
	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	author, err := models.FindAuthor(c.Context(), database.DB, authorId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	if _, err := author.Delete(c.Context(), database.DB); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func UpdateAuthor(c fiber.Ctx) error {
	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	newAuthor := models.Author{}
	if err := c.Bind().Body(&newAuthor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	author, err := models.FindAuthor(c.Context(), database.DB, authorId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	author.Name = newAuthor.Name
	if _, err := author.Update(c.Context(), database.DB, boil.Infer()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(author)
}
