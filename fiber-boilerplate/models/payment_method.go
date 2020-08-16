package models

import (
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/jinzhu/gorm"
	"strings"
)

type PaymentMethod struct {
	*gorm.Model
	Name     string `json:"Name" gorm:"name"`          //nolint:gofmt
	Slug     string `json:"Slug" gorm:"slug"`          //nolint:gofmt
	IsActive bool   `json:"IsActive" gorm:"is_active"` //nolint:gofmt
	Currency string `json:"Currency" gorm:"curreny"`
}

func GetPaymentMethodBySlug(slug string) (*PaymentMethod, error) {
	var pm PaymentMethod
	if err := DB.Where(&PaymentMethod{Slug: slug}).FirstOrCreate(&pm).Error; err != nil {
		return nil, err
	}
	return &pm, nil
}

func (u *PaymentMethod) BeforeCreate() (err error) {
	if u.Name == "" {
		u.Name = strings.ToTitle(u.Slug)
	}
	if u.Currency == "" {
		u.Currency = "USD"
	}
	return
}
