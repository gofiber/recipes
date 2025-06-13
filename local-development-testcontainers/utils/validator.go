package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

// Validate validates the input struct
func Validate(payload interface{}) *fiber.Error {
	err := validate.Struct(payload)
	if err != nil {
		var targetErr *fiber.Error
		if errors.As(err, &targetErr) {
			return &fiber.Error{
				Code:    fiber.StatusBadRequest,
				Message: targetErr.Error(),
			}
		}

		return nil
	}

	return nil
}

// CUSTOM VALIDATION RULES =============================================

// Password validation rule: required,min=6,max=100
var _ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	l := len(fl.Field().String())

	return l >= 6 && l < 100
})
