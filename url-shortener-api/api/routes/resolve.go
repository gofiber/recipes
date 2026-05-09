package routes

import (
	"log"

	"github.com/amalshaji/stoyle/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
)

// ResolveURL resolves a short URL to its original and redirects the client.
func ResolveURL(c fiber.Ctx) error {
	url := c.Params("url")
	ctx := c.Context()

	value, err := database.DB0.Get(ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	// Increment the global redirect counter. Log the error but don't fail the
	// redirect — the counter is non-critical.
	if err := database.DB1.Incr(ctx, "counter").Err(); err != nil {
		log.Printf("warn: failed to increment redirect counter: %v", err)
	}

	return c.Redirect().Status(301).To(value)
}
