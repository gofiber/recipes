package model

import (
	"aws-ses-sender/config"
	"time"

	"gorm.io/gorm"
)

const (
	EmailMessageStatusCreated    = iota // Creation complete
	EmailMessageStatusProcessing        // Processing
	EmailMessageStatusSent              // Sent
	EmailMessageStatusFailed            // Failed
	EmailMessageStatusStopped           // Stopped
)

type Request struct {
	gorm.Model
	TopicID     string     `json:"topic_id" gorm:"index"`
	MessageID   string     `json:"message_id" gorm:"index;type:varchar(255);default:''"`
	To          string     `json:"to" gorm:"not null;type:varchar(255)"`
	Subject     string     `json:"subject" gorm:"not null;type:varchar(255)"`
	Content     string     `json:"content" gorm:"not null;type:text"`
	ScheduledAt *time.Time `json:"scheduled_at" gorm:"null;index;type:timestamp with time zone"`
	Status      int        `json:"status" gorm:"default:0;index;not null;type:smallint"`
	Error       string     `json:"error" gorm:"default:'';type:varchar(255)"`
}

func (m *Request) TableName() string {
	return "email_requests"
}

type Result struct {
	gorm.Model
	RequestId uint    `json:"request_id" gorm:"index;not null"`
	Request   Request `json:"request" gorm:"foreignKey:RequestId;references:ID"`
	Status    string  `json:"status" gorm:"not null;index;type:varchar(50)"`
	Raw       string  `json:"raw" gorm:"null;type:json"`
}

func (m *Result) TableName() string {
	return "email_results"
}

func init() {
	db := config.GetDB()
	_ = db.AutoMigrate(&Request{})
	_ = db.AutoMigrate(&Result{})
}
