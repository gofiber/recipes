package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func RateLimiterMiddleware(client *redis.Client, ttl time.Duration, maxRequests int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		key := fmt.Sprintf("rate-limit:%s", c.IP())

		count, err := client.Incr(ctx, key).Result()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   "Redis Error",
				"details": err.Error(),
			})
		}
		if count == 1 {
			client.Expire(ctx, key, ttl)
		}
		if count > int64(maxRequests) {
			ttlVal, _ := client.TTL(ctx, key).Result()
			c.Set("Retry-After", fmt.Sprintf("%.0f", ttlVal.Seconds()))
			return c.Status(429).JSON(fiber.Map{
				"error": "Too Many Requests",
			})
		}
		return c.Next()
	}
}
