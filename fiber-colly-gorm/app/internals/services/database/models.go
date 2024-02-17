package database

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	ID     int    `json:"id"`
	Text   string `json:"quote"`
	Author string `json:"author"`
}
