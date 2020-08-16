package models

import (
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/jinzhu/gorm"
	"strconv"
)

type User struct {
	*gorm.Model
	FirstName     string  `json:"FirstName" gorm:"first_name"` //nolint:gofmt
	LastName      string  `json:"LastName" gorm:"last_name"`   //nolint:gofmt
	Email         string  `json:"Email" gorm:"email"`
	Password      string  `json:"-" gorm:"password"`
	Balance       float32 `json:"balance" gorm:"balance"`
	EmailVerified bool    `json:"email_verified" gorm:"email_verified"`
	Currency      string  `json:"Currency" gorm:"curreny"`
}

func AllUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := DB.Where(&User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil //nolint:wsl
}

func GetUserById(id interface{}) (*User, error) {
	var user User
	if err := DB.Where("id = ? ", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) AddAmount(amount string) {
	value, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		// do something sensible
	}
	u.Balance += float32(value)
	DB.Save(&u)
}

func (u *User) Files() (*UserFile, error) {
	var uf UserFile
	if err := DB.Where(UserFile{UserID: u.ID}).First(&uf).Error; err != nil {
		return nil, err
	}
	return &uf, nil
}
