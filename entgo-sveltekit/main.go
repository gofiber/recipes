package main

import (
	"app/config/database"
	"app/fixtures"
	"app/handler"
	"app/template"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"

)

const (
	databaseName = ".sqlite"
	appName      = "Go Fiber, Ent ORM and Sveltekit FSA"
	apiVersion   = "v1"
	port         = ":3000"
)

func main() {
	// Load fixtures
	if err := fixtures.CheckAndLoadFixtures(databaseName); err != nil {
		panic(err)
	}

	// Connect to the database
	client, err := database.Connect(databaseName)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: appName,
	})
	defer app.Shutdown()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Routes
	v1 := app.Group("/api/" + apiVersion)

	todoHandler := handler.NewTodoHandler(client)

	todo := v1.Group("/todo")
	todo.Get("/list", todoHandler.GetAllTodos)
	todo.Get("/get/:id", todoHandler.GetTodoByID)
	todo.Post("/create", todoHandler.CreateTodo)
	todo.Put("/update/:id", todoHandler.UpdateTodoByID)
	todo.Delete("/delete/:id", todoHandler.DeleteTodoByID)

	// Serve static files
	app.All("/*", static.New("", static.Config{
		FS:           os.DirFS(template.Dist()),
		NotFoundFile: "index.html",
		IndexNames:   []string{"index.html"},
	}))

	// Start the server
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
