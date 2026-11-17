package controller

import (
	"context"
	"time"

	"fiber-sqlc/database"
	"fiber-sqlc/database/sqlc"

	"github.com/gofiber/fiber/v3"
)

func GetPosts(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	posts, err := sqlc.New(database.DB).GetPosts(ctx)
	if err != nil {
		return internalError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(posts)
}

func GetPost(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	postId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid post id")
	}

	post, err := sqlc.New(database.DB).GetPost(ctx, postId)
	if err != nil {
		return lookupFailed(c, err, "post")
	}

	return c.Status(fiber.StatusOK).JSON(post)
}

func NewPost(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var post sqlc.NewPostParams
	if err := c.Bind().Body(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	newPost, err := sqlc.New(database.DB).NewPost(ctx, post)
	if err != nil {
		if isForeignKeyViolation(err) {
			return c.Status(fiber.StatusUnprocessableEntity).SendString("author does not exist")
		}
		return internalError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(newPost)
}

func DeletePost(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	postId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid post id")
	}

	if err := sqlc.New(database.DB).DeletePost(ctx, postId); err != nil {
		return internalError(c, err)
	}

	return c.Status(fiber.StatusOK).SendString("Post deleted")
}

func UpdatePost(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	postId, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid post id")
	}

	var post sqlc.UpdatePostParams
	if err := c.Bind().Body(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	post.ID = postId
	if _, err := sqlc.New(database.DB).UpdatePost(ctx, post); err != nil {
		if isForeignKeyViolation(err) {
			return c.Status(fiber.StatusUnprocessableEntity).SendString("author does not exist")
		}
		return lookupFailed(c, err, "post")
	}

	return c.Status(fiber.StatusOK).SendString("Post updated")
}
