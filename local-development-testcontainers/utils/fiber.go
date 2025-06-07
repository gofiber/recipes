package utils

import (
	"github.com/gofiber/fiber/v3"
)

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(ctx fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.Bind().Body(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(ctx fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.Bind().Body(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return Validate(body)
}

// GetUser is helper function for getting authenticated user's id
func GetUser(c fiber.Ctx) *uint64 {
	id, _ := c.Locals("USER").(uint64)
	return &id
}
