package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	fibersentry "github.com/gofiber/contrib/v3/sentry"
	"github.com/gofiber/fiber/v3"
)

// SentryRecover recovers from panics, captures the stack trace in Sentry,
// and returns a structured 500 JSON response with the Sentry EventID.
func SentryRecover() fiber.Handler {
	return func(c fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				hub := fibersentry.GetHubFromContext(c)
				var eventID *sentry.EventID
				if hub != nil {
					eventID = hub.RecoverWithContext(c.Context(), r)
				}

				resp := fiber.Map{
					"status": fiber.StatusInternalServerError,
					"error":  "Internal Server Error",
				}
				if eventID != nil {
					resp["event_id"] = string(*eventID)
				}

				err = c.Status(fiber.StatusInternalServerError).JSON(resp)
			}
		}()
		return c.Next()
	}
}

// SentryErrorHandler is a custom Fiber ErrorHandler that reports server errors (>= 500)
// to Sentry while returning clean JSON error responses to clients.
func SentryErrorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if fiberErr, ok := errors.AsType[*fiber.Error](err); ok {
		code = fiberErr.Code
	}

	resp := fiber.Map{
		"status": code,
		"error":  err.Error(),
	}

	// Only capture unhandled server errors (5xx) to Sentry to avoid noise from 4xx client errors
	if code >= fiber.StatusInternalServerError {
		hub := fibersentry.GetHubFromContext(c)
		if hub != nil {
			if eventID := hub.CaptureException(err); eventID != nil {
				resp["event_id"] = string(*eventID)
			}
		}
	}

	return c.Status(code).JSON(resp)
}

// SetupApp initializes the Fiber app with contrib/v3/sentry middleware, recovery, and custom error handling.
func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: SentryErrorHandler,
	})

	// 1. Official contrib/v3/sentry middleware: clones a per-request hub, attaches deep-copied HTTP request data
	app.Use(fibersentry.New())

	// 2. Attach user context to the per-request hub when authenticated headers are present
	app.Use(func(c fiber.Ctx) error {
		if userID := c.Get("X-User-ID"); userID != "" {
			if hub := fibersentry.GetHubFromContext(c); hub != nil {
				hub.Scope().SetUser(sentry.User{ID: userID})
			}
		}
		return c.Next()
	})

	// 3. Panic recovery handler placed inside fibersentry.New(): recovers panics, logs to Sentry, returns 500 JSON with event_id
	app.Use(SentryRecover())

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Fiber with Sentry Error Handling",
		})
	})

	app.Get("/users/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		if id == "invalid" {
			// Client errors (4xx) return clean JSON and do not trigger Sentry events
			return fiber.NewError(fiber.StatusBadRequest, "invalid user id provided")
		}
		return c.JSON(fiber.Map{
			"id":   id,
			"name": "Jane Doe",
		})
	})

	app.Get("/error", func(c fiber.Ctx) error {
		// Server errors (5xx) are automatically forwarded to Sentry
		return errors.New("database connection failed: timeout connecting to replica")
	})

	app.Get("/panic", func(c fiber.Ctx) error {
		// Panics are safely caught and reported to Sentry with full stack traces
		panic("runtime error: nil pointer dereference in critical pipeline")
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	return app
}

// SentryOptions returns production Sentry client configuration.
func SentryOptions(dsn string) sentry.ClientOptions {
	return sentry.ClientOptions{
		Dsn:              dsn,
		AttachStacktrace: true,
		TracesSampleRate: 1.0,
		Environment:      "production",
		Release:          "recipe-sentry@1.0.0",
	}
}

func main() {
	dsn := os.Getenv("SENTRY_DSN")
	err := sentry.Init(SentryOptions(dsn))
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	app := SetupApp()
	log.Fatal(app.Listen(":3000"))
}
