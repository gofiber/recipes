package handler

import (
	"app/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoHandler struct {
	Client *entity.Client
}

// NewTodoHandler creates a new instance of TodoHandler.
func NewTodoHandler(client *entity.Client) *TodoHandler {
	return &TodoHandler{
		Client: client,
	}
}

// GetAllTodos retrieves all todo items.
//
// [GET] /todos
func (th *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	todos, err := th.Client.Todo.Query().All(c.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "Failed to retrieve todos")
	}
	return c.JSON(todos)
}

// GetTodoByID retrieves a specific todo item.
//
// [GET] /todo/get/:id
func (th *TodoHandler) GetTodoByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}

	todo, err := th.Client.Todo.Get(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo item not found"})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

// CreateTodo adds a new todo item.
//
// POST /todo/create
func (th *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	todo := new(entity.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}
	if todo.ID == uuid.Nil {
		todo.ID = uuid.New()
	}
	todo.Completed = false

	createdTodo, err := th.Client.Todo.Create().
		SetID(todo.ID).
		SetContent(todo.Content).
		SetCompleted(todo.Completed).
		Save(c.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "Failed to create todo")
	}
	return c.Status(fiber.StatusCreated).JSON(createdTodo)
}

// UpdateTodoByID updates a specific todo item by ID.
//
// PUT /todo/update/:id
func (th *TodoHandler) UpdateTodoByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}
	data := new(entity.Todo)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	todo, err := th.Client.Todo.Get(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo item not found"})
	}
	updatedTodo, err := todo.Update().
		SetContent(data.Content).
		SetCompleted(data.Completed).
		Save(c.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, "Failed to update todo")
	}
	return c.JSON(updatedTodo)
}

// DeleteTodoByID deletes a specific todo item by ID.
//
// DELETE /todo/delete/:id
func (th *TodoHandler) DeleteTodoByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}

	err = th.Client.Todo.DeleteOneID(id).Exec(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo item not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Todo item deleted"})
}
