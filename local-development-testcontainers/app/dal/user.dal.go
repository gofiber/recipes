package dal

import (
	"gorm.io/gorm"
)

// User struct defines the user
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Todos    []Todo `gorm:"foreignKey:User"`
}

// CreateUser create a user entry in the user's table
func CreateUser(db *gorm.DB, user *User) *gorm.DB {
	return db.Create(user)
}

// FindUser searches the user's table with the condition given
func FindUser(db *gorm.DB, dest interface{}, conds ...interface{}) *gorm.DB {
	return db.Model(&User{}).Take(dest, conds...)
}

// FindUserByEmail searches the user's table with the email given
func FindUserByEmail(db *gorm.DB, dest interface{}, email string) *gorm.DB {
	return FindUser(db, dest, "email = ?", email)
}
