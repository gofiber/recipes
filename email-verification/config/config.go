package config

import (
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
	return &Config{
		SMTPHost:       "smtp.gmail.com",
		SMTPPort:       587,
		SMTPUser:       "tahasfhga@gmail.com",
		SMTPPassword:   "bakkcmkakpfxwuef",
		CodeExpiration: time.Minute * 1,
	}
}
