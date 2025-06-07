package dal

import (
	"gorm.io/gorm"
)

// Todo struct defines the Todo Model
type Todo struct {
	gorm.Model
	Task      string  `gorm:"not null"`
	Completed bool    `gorm:"default:false"`
	User      *uint64 `gorm:"index,not null"`
	// this is a pointer because int == 0,
}

// CreateTodo create a todo entry in the todo's table
func CreateTodo(db *gorm.DB, todo *Todo) *gorm.DB {
	return db.Create(todo)
}

// FindTodo finds a todo with given condition
func FindTodo(db *gorm.DB, dest interface{}, conds ...interface{}) *gorm.DB {
	return db.Model(&Todo{}).Take(dest, conds...)
}

// FindTodoByUser finds a todo with given todo and user identifier
func FindTodoByUser(db *gorm.DB, dest interface{}, todoIden interface{}, userIden interface{}) *gorm.DB {
	return FindTodo(db, dest, "todos.id = ? AND todos.user = ?", todoIden, userIden)
}

// FindTodosByUser finds the todos with user's identifier given
func FindTodosByUser(db *gorm.DB, dest interface{}, userIden interface{}) *gorm.DB {
	return db.Model(&Todo{}).Find(dest, "todos.user = ?", userIden)
}

// DeleteTodo deletes a todo from todos' table with the given todo and user identifier
func DeleteTodo(db *gorm.DB, todoIden interface{}, userIden interface{}) *gorm.DB {
	return db.Unscoped().Delete(&Todo{}, "todos.id = ? AND todos.user = ?", todoIden, userIden)
}

// UpdateTodo allows to update the todo with the given todoID and userID
func UpdateTodo(db *gorm.DB, todoIden interface{}, userIden interface{}, data interface{}) *gorm.DB {
	return db.Model(&Todo{}).Where("todos.id = ? AND todos.user = ?", todoIden, userIden).Updates(data)
}
