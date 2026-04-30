package handler

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// Hello handles GET /api/hello (protected).
func Hello(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

// Salut handles GET /salut (public).
func Salut(c fiber.Ctx) error {
	return c.SendString("Salut, World 👋!")
}

// Ciao handles POST /ciao (public).
func Ciao(c fiber.Ctx) error {
	return c.SendString("Ciao, World 👋!")
}

// Ayubowan handles GET /api/ayubowan (protected).
// It reads the authenticated user claims from the request context.
func Ayubowan(c fiber.Ctx) error {
	claims := c.Locals("user")
	log.Printf("authenticated user claims: %v", claims)
	return c.SendString("Ayubowan👋!")
}

// Salanthe handles GET /salanthe (public).
func Salanthe(c fiber.Ctx) error {
	return c.SendString("Salanthe👋!")
}
