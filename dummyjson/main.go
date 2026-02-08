package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
)

var client = http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		resp, err := client.Get("https://dummyjson.com/products/1")
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
				"error":   err.Error(),
			})
		}

		if _, err := io.Copy(c.Response().BodyWriter(), resp.Body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusOK)
	})

	log.Fatal(app.Listen(":3000"))
}
