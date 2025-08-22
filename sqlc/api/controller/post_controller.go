package controller

import (
	"context"
	"strconv"
	"time"

	"fiber-sqlc/database"
	"fiber-sqlc/database/sqlc"

	"github.com/gofiber/fiber/v3"
)

func GetPosts(c fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	posts, err := sqlc.New(database.DB).GetPosts(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(posts)
}

func GetPost(c fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	post, err := sqlc.New(database.DB).GetPost(ctx, int32(postId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

func NewPost(c fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	var post sqlc.NewPostParams
	if err := c.Bind().Body(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	newPost, err := sqlc.New(database.DB).NewPost(ctx, post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(newPost)
}

func DeletePost(c fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := sqlc.New(database.DB).DeletePost(ctx, int32(postId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Post deleted")
}

func UpdatePost(c fiber.Ctx) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	id := c.Params("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var post sqlc.UpdatePostParams
	if err := c.Bind().Body(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	post.ID = int32(postId)
	if _, err := sqlc.New(database.DB).UpdatePost(ctx, post); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Post updated")
}
