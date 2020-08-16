package models

import "time"

type UserTransactionLog struct {
	ID         uint64 `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	UserID     uint       `json:"UserID" gorm:"user_id"` //nolint:gofmt
	User       User       `gorm:"foreignkey:UserID"`
	Item       string     `json:"Item" gorm:"item"`
	ClientCost float32    `json:"Cost" gorm:"cost"`
}
