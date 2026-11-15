package api

import (
	"net/url"

	"catalog/domain"

	"github.com/gofiber/fiber/v3"
)

type handler struct {
	productService domain.Service
}

// productCode returns the :code route parameter. Fiber leaves it percent-encoded, so a code
// such as "a b" arrives as "a%20b" and has to be decoded before it is used as a key.
func productCode(ctx fiber.Ctx) (string, error) {
	return url.PathUnescape(ctx.Params("code"))
}

func (h *handler) Get(ctx fiber.Ctx) error {
	code, err := productCode(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}
	p, err := h.productService.Find(code)
	if err != nil {
		return ctx.Status(404).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Post(ctx fiber.Ctx) error {
	p := &domain.Product{}
	if err := ctx.Bind().Body(&p); err != nil {
		return ctx.Status(500).JSON(nil)
	}
	err := h.productService.Store(p)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Put(ctx fiber.Ctx) error {
	p := &domain.Product{}
	if err := ctx.Bind().Body(&p); err != nil {
		return ctx.Status(500).JSON(nil)
	}
	err := h.productService.Update(p)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Delete(ctx fiber.Ctx) error {
	code, err := productCode(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(nil)
	}
	err = h.productService.Delete(code)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.Status(200).JSON(fiber.Map{"message": "ok"})
}

func (h *handler) GetAll(ctx fiber.Ctx) error {
	p, err := h.productService.FindAll()
	if err != nil {
		return ctx.Status(404).JSON(nil)
	}
	return ctx.JSON(&p)
}

// NewHandler  New handler instantiates a http handler for our product service
func NewHandler(productService domain.Service) *handler {
	return &handler{productService: productService}
}
