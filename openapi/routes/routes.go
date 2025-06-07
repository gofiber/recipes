package routes

import (
	"openapi/handlers"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	api := humafiber.New(app, huma.DefaultConfig("Book API", "1.0.0"))
	group := huma.NewGroup(api, "/v1")

	huma.Get(group, "/books", handlers.GetAllBooks)
	huma.Get(group, "/books/{id}", handlers.GetBookByID)

	// or more details to the endpoint
	// huma.Register(api, huma.Operation{
	// 	OperationID: "get-book-by-id",
	// 	Method:      http.MethodGet,
	// 	Path:        "/book/{id}",
	// 	Summary:     "Get a book",
	// 	Description: "Get a book by book ID.",
	// 	Tags:        []string{"Books"},
	// }, handlers.GetBookByID)

	huma.Post(group, "/books", handlers.RegisterBook)
	huma.Patch(group, "/books/{id}", handlers.UpdateBook)
	huma.Delete(group, "/books/{id}", handlers.DeleteBook)

	return app
}
