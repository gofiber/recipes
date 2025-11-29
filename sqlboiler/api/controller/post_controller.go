package controller

import (
	"context"
	"net/http"
	"strconv"

	"fiber-sqlboiler/database"
	"fiber-sqlboiler/models"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetPosts(c *fiber.Ctx) error {
	posts, err := models.Posts().All(context.Background(), database.DB)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	post, err := models.FindPost(context.Background(), database.DB, postId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(post)
}

func NewPost(c *fiber.Ctx) error {
	post := models.Post{}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	if err := post.Insert(context.Background(), database.DB, boil.Infer()); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	post, err := models.FindPost(context.Background(), database.DB, postId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}
	if _, err := post.Delete(context.Background(), database.DB); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return c.SendStatus(http.StatusOK)
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newPost := models.Post{}
	if err := c.BodyParser(&newPost); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	post, err := models.FindPost(context.Background(), database.DB, postId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	post.Title = newPost.Title
	post.Content = newPost.Content
	if _, err := post.Update(context.Background(), database.DB, boil.Infer()); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(post)
}
