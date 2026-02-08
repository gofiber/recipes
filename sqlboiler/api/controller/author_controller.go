package controller

import (
	"context"
	"strconv"

	"fiber-sqlboiler/database"
	"fiber-sqlboiler/models"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gofiber/fiber/v3"
)

func GetAuthors(c fiber.Ctx) error {
	authors, err := models.Authors().All(context.Background(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(authors)
}

func GetAuthor(c fiber.Ctx) error {
	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	author, err := models.FindAuthor(context.Background(), database.DB, authorId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(author)
}

func NewAuthor(c fiber.Ctx) error {
	author := models.Author{}
	if err := c.Bind().Body(&author); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	if err := author.Insert(context.Background(), database.DB, boil.Infer()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(author)
}

func DeleteAuthor(c fiber.Ctx) error {
	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	author, err := models.FindAuthor(context.Background(), database.DB, authorId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	if _, err := author.Delete(context.Background(), database.DB); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func UpdateAuthor(c fiber.Ctx) error {
	id := c.Params("id")
	authorId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	newAuthor := models.Author{}
	if err := c.Bind().Body(&newAuthor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	author, err := models.FindAuthor(context.Background(), database.DB, authorId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	author.Name = newAuthor.Name
	if _, err := author.Update(context.Background(), database.DB, boil.Infer()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(author)
}
