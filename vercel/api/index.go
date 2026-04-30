package handler

import (
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

var (
	app     *fiber.App
	appOnce sync.Once
)

func init() {
	appOnce.Do(func() {
		app = fiber.New(fiber.Config{
			ErrorHandler: func(ctx fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
				if e, ok := err.(*fiber.Error); ok {
					code = e.Code
				}
				return ctx.Status(code).JSON(fiber.Map{
					"error": err.Error(),
				})
			},
		})

		app.Get("/v1", func(ctx fiber.Ctx) error {
			return ctx.JSON(fiber.Map{
				"version": "v1",
			})
		})

		app.Get("/v2", func(ctx fiber.Ctx) error {
			return ctx.JSON(fiber.Map{
				"version": "v2",
			})
		})

		app.Get("/", func(ctx fiber.Ctx) error {
			return ctx.JSON(fiber.Map{
				"uri":  ctx.Request().URI().String(),
				"path": ctx.Path(),
			})
		})

		// 404 catch-all
		app.Use(func(ctx fiber.Ctx) error {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "not found",
			})
		})
	})
}

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `fiber.Ctx`
	r.RequestURI = r.URL.String()

	adaptor.FiberApp(app).ServeHTTP(w, r)
}
