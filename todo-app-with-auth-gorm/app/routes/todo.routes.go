package routes

import (
	"numtostr/gotodo/app/services"
	"numtostr/gotodo/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

// TodoRoutes contains all routes relative to /todo
func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo").Use(middleware.Auth)

	r.Post("/create", services.CreateTodo)
	r.Get("/list", services.GetTodos)
	r.Get("/:todoID", services.GetTodo)
	r.Patch("/:todoID", services.UpdateTodoTitle)
	r.Patch("/:todoID/check", services.CheckTodo)
	r.Delete("/:todoID", services.DeleteTodo)
}
