package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
)

const defaultUpstreamURL = "https://dummyjson.com/products/1"

var client = http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	app := fiber.New()

	// GET /proxy?url=<upstream-url>
	// Falls back to dummyjson products endpoint when no url query param is given.
	app.Get("/proxy", func(c fiber.Ctx) error {
		target := c.Query("url", defaultUpstreamURL)

		resp, err := client.Get(target) //nolint:noctx // simple recipe, context omitted for brevity
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return c.Status(resp.StatusCode).JSON(&fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("upstream returned status %d", resp.StatusCode),
			})
		}

		c.Set(fiber.HeaderContentType, resp.Header.Get(fiber.HeaderContentType))

		if _, err := io.Copy(c.Response().BodyWriter(), resp.Body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
