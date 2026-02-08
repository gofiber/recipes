package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"username" gorm:"unique;" validate:"required,email,min=6,max=32"`
	Password string `json:"-" gorm:"type:text;" validate:"required,min=6"`
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
