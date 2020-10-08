package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

// Cache will return a caching middleware
func Cache(exp time.Duration) fiber.Handler {
	return cache.New(cache.Config{
		Expiration:   exp,
		CacheControl: true,
	})
}
