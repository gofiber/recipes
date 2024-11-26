package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	authKey  = "apiKey"
	authName = "x-api-key"
	authSrc  = "header"
)

var (
	authList   = []string{"valid-key"}
	errMissing = &fiber.Error{
		Code:    403000,
		Message: "Missing API key",
	}
	errInvalid = &fiber.Error{
		Code:    403001,
		Message: "Invalid API key",
	}
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(keyauth.New(keyauth.Config{
		SuccessHandler: successHandler,
		ErrorHandler:   errHandler,
		KeyLookup:      strings.Join([]string{authSrc, authName}, ":"),
		Validator:      validator,
		ContextKey:     authKey,
	}))

	log.Fatal(app.Listen(":1337"))
}

func successHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func errHandler(ctx *fiber.Ctx, err error) error {
	ctx.Status(fiber.StatusForbidden)

	if err == errMissing {
		return ctx.JSON(errMissing)
	}

	return ctx.JSON(errInvalid)
}

func validator(ctx *fiber.Ctx, s string) (bool, error) {
	if s == "" {
		return false, errMissing
	}

	for _, val := range authList {
		if s == val {
			return true, nil
		}
	}

	return false, errInvalid
}
