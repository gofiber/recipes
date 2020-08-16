package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	*gorm.Model
	PaymentID       uint    `json:"PaymentID" gorm:"payment_id"` //nolint:gofmt
	Payment         Payment `gorm:"foreignkey:PaymentID"`
	TotalAmount     float32 `json:"TotalAmount" gorm:"total_amount"`
	Status          string  `json:"Status" gorm:"status"`
	TransactionType string  `json:"TransactionType" gorm:"transaction_type"` // Debit | Credit
}
