package models

import "gorm.io/gorm"

// Book model
type Book struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}
