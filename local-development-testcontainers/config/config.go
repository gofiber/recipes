//go:build !dev

package config

import (
	"github.com/gofiber/fiber/v3"
)

// ConfigureApp configures the fiber app, including the database connection string.
// The connection string is retrieved from the environment variable DB, or using
// falls back to a default connection string targeting localhost if DB is not set.
func ConfigureApp(cfg fiber.Config) (*AppConfig, error) {
	app := fiber.New(cfg)

	db := getEnv("DB", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	DB = db

	return &AppConfig{
		App:            app,
		StartupCancel:  func() {}, // No-op for production
		ShutdownCancel: func() {}, // No-op for production
	}, nil
}
