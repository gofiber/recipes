package models

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name      string `json:"name"`
	Breed     string `json:"breed"`
	Age       int    `json:"age"`
	IsGoodBoy bool   `json:"isGoodBoy" gorm:"default:true"`
}
