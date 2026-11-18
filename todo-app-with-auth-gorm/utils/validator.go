package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

// Validate validates the input struct
func Validate(payload any) *fiber.Error {
	err := validate.Struct(payload)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		// The payload is not something the validator can check, for example a non-struct.
		return fiber.ErrBadRequest
	}

	messages := make([]string, 0, len(validationErrors))
	for _, fieldErr := range validationErrors {
		messages = append(
			messages,
			fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", fieldErr.Field(), fieldErr.Value(), fieldErr.Tag()),
		)
	}

	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: strings.Join(messages, ","),
	}
}

// CUSTOM VALIDATION RULES =============================================

// Password validation rule: required, at least 6 and at most 72 bytes.
// bcrypt rejects anything longer than 72 bytes, so a longer password has to be
// refused here, before it reaches password.Generate.
var _ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	l := len(fl.Field().String())

	return l >= 6 && l <= 72
})
