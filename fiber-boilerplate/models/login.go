package models

import (
	"errors"
	"github.com/gookit/validate"
	. "github.com/itsursujit/fiber-boilerplate/app"
)

type Login struct {
	Email    string `json:"email" gorm:"email" form:"email" validate:"required|email"`
	Password string `json:"password" gorm:"password" form:"password" validate:"required"`
}

// Messages you can custom validator error messages.
func (f Login) Messages() map[string]string {
	return validate.MS{
		"required": "oh! the {field} is required",
		"email":    "Invalid email format",
	}
}

func (Login) TableName() string {
	return "users"
}

func (l *Login) CheckLogin() (*User, error) {

	user, err := GetUserByEmail(l.Email)
	if err != nil {
		return nil, errors.New("invalid Username or Password")
	}
	match, err := Hash.Match(l.Password, user.Password)
	if !match {
		return nil, errors.New("invalid Username or Password")
	}
	return user, nil
}
