package dal

import (
	"numtostr/gotodo/config/database"

	"gorm.io/gorm"
)

// Todo struct defines the Todo Model
type Todo struct {
	gorm.Model
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	User      *uint  `gorm:"index,not null"`
	// this is a pointer because int == 0,
}

// CreateTodo create a todo entry in the todo's table
func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Create(todo)
}

// FindTodo finds a todo with given condition
func FindTodo(dest any, conds ...any) *gorm.DB {
	return database.DB.Model(&Todo{}).Take(dest, conds...)
}

// FindTodoByUser finds a todo with given todo and user identifier
func FindTodoByUser(dest any, todoIden any, userIden any) *gorm.DB {
	return FindTodo(dest, "id = ? AND user = ?", todoIden, userIden)
}

// FindTodosByUser finds the todos with user's identifier given
func FindTodosByUser(dest any, userIden any) *gorm.DB {
	return database.DB.Model(&Todo{}).Find(dest, "user = ?", userIden)
}

// DeleteTodo deletes a todo from todos' table with the given todo and user identifier
func DeleteTodo(todoIden any, userIden any) *gorm.DB {
	return database.DB.Unscoped().Delete(&Todo{}, "id = ? AND user = ?", todoIden, userIden)
}

// UpdateTodo allows to update the todo with the given todoID and userID
func UpdateTodo(todoIden any, userIden any, data any) *gorm.DB {
	return database.DB.Model(&Todo{}).Where("id = ? AND user = ?", todoIden, userIden).Updates(data)
}
