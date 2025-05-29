package config

import (
	"context"

	"github.com/gofiber/fiber/v3"
)

// AppConfig holds the application configuration and cleanup functions
type AppConfig struct {
	// App is the Fiber app instance.
	App *fiber.App
	// StartupCancel is the context cancel function for the services startup.
	StartupCancel context.CancelFunc
	// ShutdownCancel is the context cancel function for the services shutdown.
	ShutdownCancel context.CancelFunc
}
