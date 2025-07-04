package services

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"local-development/testcontainers/app/dal"
	"local-development/testcontainers/app/types"
	"local-development/testcontainers/config/database"
	"local-development/testcontainers/utils"
)

// CreateTodo is responsible for create todo
func CreateTodo(c fiber.Ctx) error {
	b := new(types.CreateDTO)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	d := &dal.Todo{
		Task: b.Task,
		User: utils.GetUser(c),
	}

	if err := dal.CreateTodo(database.DB, d).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	// Make sure that gorm.Model.ID is uint64, which could happen
	// if the machine compiling the code has multiple versions of gorm.
	uid := uint64(d.ID)

	return c.JSON(&types.TodoCreateResponse{
		Todo: &types.TodoResponse{
			ID:        uid,
			Task:      d.Task,
			Completed: d.Completed,
		},
	})
}

// GetTodos returns the todos list
func GetTodos(c fiber.Ctx) error {
	d := &[]types.TodoResponse{}

	err := dal.FindTodosByUser(database.DB, d, utils.GetUser(c)).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.TodosResponse{
		Todos: d,
	})
}

// GetTodo return a single todo
func GetTodo(c fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	d := &types.TodoResponse{}

	err := dal.FindTodoByUser(database.DB, d, todoID, utils.GetUser(c)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(&types.TodoCreateResponse{})
	}

	return c.JSON(&types.TodoCreateResponse{
		Todo: d,
	})
}

// DeleteTodo deletes a single todo
func DeleteTodo(c fiber.Ctx) error {
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	res := dal.DeleteTodo(database.DB, todoID, utils.GetUser(c))
	if res.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusConflict, "Unable to delete todo")
	}

	err := res.Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully deleted",
	})
}

// CheckTodo TODO
func CheckTodo(c fiber.Ctx) error {
	b := new(types.CheckTodoDTO)
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	err := dal.UpdateTodo(database.DB, todoID, utils.GetUser(c), map[string]interface{}{"completed": b.Completed}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully updated",
	})
}

// UpdateTodoTitle TODO
func UpdateTodoTitle(c fiber.Ctx) error {
	b := new(types.CreateDTO)
	todoID := c.Params("todoID")

	if todoID == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid todoID")
	}

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	err := dal.UpdateTodo(database.DB, todoID, utils.GetUser(c), &dal.Todo{Task: b.Task}).Error
	if err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return c.JSON(&types.MsgResponse{
		Message: "Todo successfully updated",
	})
}
