package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/contrib/v3/circuitbreaker"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/addon/retry"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

func setupApp(upstreamURL string, cb *circuitbreaker.CircuitBreaker, backoff *retry.ExponentialBackoff) *fiber.App {
	app := fiber.New()

	app.Get("/circuit/status", cb.HealthHandler())

	app.All("/proxy/*", circuitbreaker.Middleware(cb), func(c fiber.Ctx) error {
		target := upstreamURL + "/" + c.Params("*")
		if q := c.Request().URI().QueryString(); len(q) > 0 {
			target += "?" + string(q)
		}
		attempt := func() error {
			if err := proxy.Do(c, target); err != nil {
				return err
			}
			if c.Response().StatusCode() >= fiber.StatusInternalServerError {
				return fmt.Errorf("upstream server error: %d", c.Response().StatusCode())
			}
			return nil
		}

		var err error
		// Retrying a POST could run it twice upstream.
		if fiber.IsMethodIdempotent(c.Method()) {
			err = backoff.Retry(attempt)
		} else {
			err = attempt()
		}
		if err != nil {
			// Drop the upstream headers, such as Content-Encoding, before writing our own body.
			c.Response().Reset()
			// The breaker counts this 502 as a failure.
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
		}
		return nil
	})

	return app
}

func main() {
	cb := circuitbreaker.New(circuitbreaker.Config{
		FailureThreshold: 3,
		Timeout:          5 * time.Second,
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
