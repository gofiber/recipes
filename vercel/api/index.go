package handler

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberHandlerFunc(func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})(w, r)
}
