package routes

import (
	"swagger/handlers"

	swagger "github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// New create an instance of Book app routes
func New() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	swaggerCfg := swagger.Config{
		BasePath: "/docs", // swagger ui base path
		FilePath: "./docs/swagger.json",
	}

	app.Use(swagger.New(swaggerCfg))
	api := app.Group("/api")
	v1 := api.Group("/v1", func(c fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "🐣 v1",
		})
		return c.Next()
	})

	v1.Get("/books", handlers.GetAllBooks)
	v1.Get("/books/:id", handlers.GetBookByID)
	v1.Post("/books", handlers.RegisterBook)
	v1.Delete("/books/:id", handlers.DeleteBook)

	return app
}
