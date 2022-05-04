package api

import (
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Get(ctx *fiber.Ctx)
	Post(ctx *fiber.Ctx)
	Put(ctx *fiber.Ctx)
	Delete(ctx *fiber.Ctx)
	GetAll(ctx *fiber.Ctx)
}
