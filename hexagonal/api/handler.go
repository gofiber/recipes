package api

import (
	"github.com/gofiber/fiber/v3"
)

// ProductHandler  an interface with operations to be implemented by a specific handler, ie http, gRCP
type ProductHandler interface {
	Get(ctx fiber.Ctx) error
	Post(ctx fiber.Ctx) error
	Put(ctx fiber.Ctx) error
	Delete(ctx fiber.Ctx) error
	GetAll(ctx fiber.Ctx) error
}
