package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type File struct {
	*gorm.Model
	FileName   string `json:"FileName" gorm:"file_name"`  //nolint:gofmt
	Title      string `json:"Title" gorm:"title"`         //nolint:gofmt
	MimeType   string `json:"MimeType" gorm:"mime_type"`  //nolint:gofmt
	Size       string `json:"Size" gorm:"size"`           //nolint:gofmt
	Extension  string `json:"Extension" gorm:"extension"` //nolint:gofmt
	RowCount   int64  `json:"RowCount" gorm:"row_count"`  //nolint:gofmt
	ModifiedAt time.Time
}

type UserFile struct {
	FileID   uint `json:"FileID" gorm:"file_id"` //nolint:gofmt
	UserID   uint `json:"UserID" gorm:"user_id"` //nolint:gofmt
	IsActive bool `json:"IsActive" gorm:"is_active"`
}

func (UserFile) TableName() string {
	return "user_files"
}
