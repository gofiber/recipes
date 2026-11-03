package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/addon/retry"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

func setupApp(upstreamURL string, cb *CircuitBreaker, backoff *retry.ExponentialBackoff) *fiber.App {
	app := fiber.New()

	app.Get("/circuit/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"state": cb.State().String(),
		})
	})

	app.All("/proxy/*", func(c fiber.Ctx) error {
		path := c.Params("*")
		target := fmt.Sprintf("%s/%s", upstreamURL, path)

		err := cb.Execute(func() error {
			return backoff.Retry(func() error {
				if proxyErr := proxy.Do(c, target); proxyErr != nil {
					return proxyErr
				}
				// Treat upstream 5xx errors as retriable server failures
				if c.Response().StatusCode() >= fiber.StatusInternalServerError {
					return fmt.Errorf("upstream server error: %d", c.Response().StatusCode())
				}
				return nil
			})
		})

		if err != nil {
			if errors.Is(err, ErrCircuitOpen) {
				return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
					"error":   "upstream unavailable",
					"circuit": cb.State().String(),
				})
			}
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error":   err.Error(),
				"circuit": cb.State().String(),
			})
		}

		return nil
	})

	return app
}

func main() {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		MaxFailures: 3,
		Timeout:     5 * time.Second,
	})

	backoff := retry.NewExponentialBackoff(retry.Config{
		InitialInterval: 100 * time.Millisecond,
		MaxBackoffTime:  1 * time.Second,
		Multiplier:      2.0,
		MaxRetryCount:   3,
	})

	app := setupApp("https://httpbin.org", cb, backoff)
	log.Fatal(app.Listen(":3000"))
}
