package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v3"
)

// AppConfig holds optional configuration for the Fiber app.
type AppConfig struct {
	Client *sentry.Client
}

// SentryMiddleware attaches an isolated Sentry Hub with request metadata to each request context.
func SentryMiddleware(client *sentry.Client) fiber.Handler {
	return func(c fiber.Ctx) error {
		hub := sentry.NewHub(client, sentry.NewScope())

		scope := hub.Scope()
		scope.SetTag("http.method", c.Method())
		scope.SetTag("http.route", c.Route().Path)
		scope.SetTag("client_ip", c.IP())

		if userID := c.Get("X-User-ID"); userID != "" {
			scope.SetUser(sentry.User{ID: userID})
		}

		c.SetContext(sentry.SetHubOnContext(c.Context(), hub))
		return c.Next()
	}
}

// SentryRecover recovers from panics, captures the stack trace in Sentry,
// and returns a structured 500 JSON response with the Sentry EventID.
func SentryRecover() fiber.Handler {
	return func(c fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				hub := sentry.GetHubFromContext(c.Context())
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
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
	}

	resp := fiber.Map{
		"status": code,
		"error":  err.Error(),
	}

	// Only capture unhandled server errors (5xx) to Sentry to avoid noise from 4xx client errors
	if code >= fiber.StatusInternalServerError {
		hub := sentry.GetHubFromContext(c.Context())
		if hub != nil {
			if eventID := hub.CaptureException(err); eventID != nil {
				resp["event_id"] = string(*eventID)
			}
		}
	}

	return c.Status(code).JSON(resp)
}

// SetupApp initializes the Fiber app with Sentry middleware and error handling.
func SetupApp(cfg ...AppConfig) *fiber.App {
	var client *sentry.Client
	if len(cfg) > 0 && cfg[0].Client != nil {
		client = cfg[0].Client
	} else {
		client = sentry.CurrentHub().Client()
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: SentryErrorHandler,
	})

	app.Use(SentryMiddleware(client))
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

func main() {
	dsn := os.Getenv("SENTRY_DSN")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
		Environment:      "production",
		Release:          "recipe-sentry@1.0.0",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	app := SetupApp()
	log.Fatal(app.Listen(":3000"))
}
