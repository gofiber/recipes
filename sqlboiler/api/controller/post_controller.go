package controller

import (
	"context"
	"strconv"

	"fiber-sqlboiler/database"
	"fiber-sqlboiler/models"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gofiber/fiber/v3"
)

func GetPosts(c fiber.Ctx) error {
	posts, err := models.Posts().All(context.Background(), database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(posts)
}

func GetPost(c fiber.Ctx) error {
	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	post, err := models.FindPost(context.Background(), database.DB, postId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

func NewPost(c fiber.Ctx) error {
	post := models.Post{}
	if err := c.Bind().Body(&post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	if err := post.Insert(context.Background(), database.DB, boil.Infer()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

func DeletePost(c fiber.Ctx) error {
	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	post, err := models.FindPost(context.Background(), database.DB, postId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}
	if _, err := post.Delete(context.Background(), database.DB); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

func UpdatePost(c fiber.Ctx) error {
	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	newPost := models.Post{}
	if err := c.Bind().Body(&newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	post, err := models.FindPost(context.Background(), database.DB, postId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	post.Title = newPost.Title
	post.Content = newPost.Content
	if _, err := post.Update(context.Background(), database.DB, boil.Infer()); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(post)
}
