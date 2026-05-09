package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	SMTPHost       string
	SMTPPort       int
	SMTPUser       string
	SMTPPassword   string
	CodeExpiration time.Duration
}

func GetConfig() *Config {
	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		smtpHost = "smtp.gmail.com"
	}

	smtpPortStr := os.Getenv("SMTP_PORT")
	if smtpPortStr == "" {
		smtpPortStr = "587"
	}
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT value: %s", smtpPortStr)
	}

	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	if smtpUser == "" || smtpPass == "" {
		log.Fatal("SMTP_USER and SMTP_PASS environment variables must be set")
	}

	return &Config{
		SMTPHost:       smtpHost,
		SMTPPort:       smtpPort,
		SMTPUser:       smtpUser,
		SMTPPassword:   smtpPass,
		CodeExpiration: time.Minute * 1,
	}
}
