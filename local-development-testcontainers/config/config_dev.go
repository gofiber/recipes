//go:build dev

package config

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/contrib/v3/testcontainers"
	"github.com/gofiber/fiber/v3"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// ConfigureApp configures the fiber app, including the database connection string.
// The connection string is retrieved from the environment variable DB, or using
// tries to connect to a local postgres instance if the environment variable is not set.
func ConfigureApp(cfg fiber.Config) (*AppConfig, error) {
	// Define a context provider for the services startup.
	// The timeout is applied when the context is actually used during startup.
	startupCtx, startupCancel := context.WithCancel(context.Background())
	var startupTimeoutCancel context.CancelFunc
	cfg.ServicesStartupContextProvider = func() context.Context {
		// Cancel any previous timeout context
		if startupTimeoutCancel != nil {
			startupTimeoutCancel()
		}
		// Create a new timeout context
		ctx, cancel := context.WithTimeout(startupCtx, 10*time.Second)
		startupTimeoutCancel = cancel
		return ctx
	}

	// Define a context provider for the services shutdown.
	// The timeout is applied when the context is actually used during shutdown.
	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())
	var shutdownTimeoutCancel context.CancelFunc
	cfg.ServicesShutdownContextProvider = func() context.Context {
		// Cancel any previous timeout context
		if shutdownTimeoutCancel != nil {
			shutdownTimeoutCancel()
		}
		// Create a new timeout context
		ctx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		shutdownTimeoutCancel = cancel
		return ctx
	}

	// Add the Postgres service to the app, including custom configuration.
	srv, err := setupPostgres(&cfg)
	if err != nil {
		if startupTimeoutCancel != nil {
			startupTimeoutCancel()
		}
		if shutdownTimeoutCancel != nil {
			shutdownTimeoutCancel()
		}
		startupCancel()
		shutdownCancel()
		return nil, fmt.Errorf("add postgres service: %w", err)
	}

	app := fiber.New(cfg)

	// Retrieve the Postgres service from the app, using the service key.
	postgresSrv := fiber.MustGetService[*testcontainers.ContainerService[*postgres.PostgresContainer]](app.State(), srv.Key())

	connString, err := postgresSrv.Container().ConnectionString(context.Background())
	if err != nil {
		if startupTimeoutCancel != nil {
			startupTimeoutCancel()
		}
		if shutdownTimeoutCancel != nil {
			shutdownTimeoutCancel()
		}
		startupCancel()
		shutdownCancel()
		return nil, fmt.Errorf("get postgres connection string: %w", err)
	}

	// Override the default database connection string with the one from the Testcontainers service.
	DB = connString

	return &AppConfig{
		App: app,
		StartupCancel: func() {
			if startupTimeoutCancel != nil {
				startupTimeoutCancel()
			}
			startupCancel()
		},
		ShutdownCancel: func() {
			if shutdownTimeoutCancel != nil {
				shutdownTimeoutCancel()
			}
			shutdownCancel()
		},
	}, nil
}

// setupPostgres adds a Postgres service to the app, including custom configuration to allow
// reusing the same container while developing locally.
func setupPostgres(cfg *fiber.Config) (*testcontainers.ContainerService[*postgres.PostgresContainer], error) {
	// Add the Postgres service to the app, including custom configuration.
	srv, err := testcontainers.AddService(cfg, testcontainers.NewModuleConfig(
		"postgres-db",
		"postgres:16",
		postgres.Run,
		postgres.BasicWaitStrategies(),
		postgres.WithDatabase("todos"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		tc.WithReuseByName("postgres-db-todos"),
	))
	if err != nil {
		return nil, fmt.Errorf("add postgres service: %w", err)
	}

	return srv, nil
}
