package controller

import (
	"fiber-sqlboiler/database"
	"fiber-sqlboiler/models"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gofiber/fiber/v3"
)

func GetPosts(c fiber.Ctx) error {
	posts, err := models.Posts().All(c.Context(), database.DB)
	if err != nil {
		return internalError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(posts)
}

func GetPost(c fiber.Ctx) error {
	postId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid post id")
	}
	post, err := models.FindPost(c.Context(), database.DB, postId)
	if err != nil {
		return lookupFailed(c, err, "post")
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

func NewPost(c fiber.Ctx) error {
	post := models.Post{}
	if err := c.Bind().Body(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := post.Insert(c.Context(), database.DB, boil.Infer()); err != nil {
		if isForeignKeyViolation(err) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON("author does not exist")
		}
		return internalError(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(post)
}

func DeletePost(c fiber.Ctx) error {
	postId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid post id")
	}
	post, err := models.FindPost(c.Context(), database.DB, postId)
	if err != nil {
		return lookupFailed(c, err, "post")
	}
	if _, err := post.Delete(c.Context(), database.DB); err != nil {
		return internalError(c, err)
	}
	return c.SendStatus(fiber.StatusOK)
}

func UpdatePost(c fiber.Ctx) error {
	postId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid post id")
	}

	newPost := models.Post{}
	if err := c.Bind().Body(&newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	post, err := models.FindPost(c.Context(), database.DB, postId)
	if err != nil {
		return lookupFailed(c, err, "post")
	}

	post.Title = newPost.Title
	post.Content = newPost.Content
	if _, err := post.Update(c.Context(), database.DB, boil.Infer()); err != nil {
		return internalError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(post)
}
