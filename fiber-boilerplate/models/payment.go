package models

import (
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/jinzhu/gorm"
	"github.com/plutov/paypal/v3"
)

type Payment struct {
	*gorm.Model
	PaymentMethodID    uint          `json:"PaymentMethodID" gorm:"payment_method_id"` //nolint:gofmt
	PaymentMethod      PaymentMethod `gorm:"foreignkey:PaymentMethodID"`
	UserID             uint          `json:"UserID" gorm:"user_id"` //nolint:gofmt
	User               User          `gorm:"foreignkey:UserID"`
	Amount             string        `json:"Amount" gorm:"amount"`
	Status             string        `json:"Status" gorm:"status"`
	GatewayOrderID     string        `json:"GatewayOrderID" gorm:"gateway_order_id"`
	GatewayOrderStatus string        `json:"GatewayOrderStatus" gorm:"gateway_order_status"`
	Currency           string        `json:"Currency" gorm:"curreny"`
	PayPalOrderDetail  *paypal.Order `gorm:"-"`
}

func (l *Payment) Create() (*Payment, error) {
	DB.Create(l)
	return l, nil
}

func GetPaymentByGatewayOrderID(id string) (*Payment, error) {
	var p Payment
	if err := DB.Where(&Payment{GatewayOrderID: id}).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil //nolint:wsl
}

func (l *Payment) UpdatePaymentStatusByGatewayOrderID(status string) {
	l.Status = status
	l.GatewayOrderStatus = l.PayPalOrderDetail.Status
	DB.Save(&l)
}
