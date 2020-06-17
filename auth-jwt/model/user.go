package model

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Names    string `json:"names"`
}
