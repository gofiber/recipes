package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	fiberprom "github.com/gofiber/contrib/v3/prometheus"
	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus"
)

// AppMetrics holds custom application-level business metrics.
type AppMetrics struct {
	OrdersProcessed prometheus.Counter
}

// NewAppMetrics initializes and registers domain metrics with the provided Prometheus registry.
func NewAppMetrics(reg prometheus.Registerer) *AppMetrics {
	m := &AppMetrics{
		OrdersProcessed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "app",
			Name:      "orders_processed_total",
			Help:      "Total number of orders successfully processed by the application.",
		}),
	}

	reg.MustRegister(m.OrdersProcessed)

	return m
}

// SetupApp configures the Fiber application with Prometheus HTTP middleware and routes.
func SetupApp(reg *prometheus.Registry) *fiber.App {
	// Register domain-specific metrics
	metrics := NewAppMetrics(reg)

	app := fiber.New(fiber.Config{
		AppName: "Fiber-Prometheus-Example v3",
	})

	// Also registers the Go runtime and process collectors on reg.
	app.Use(fiberprom.New(fiberprom.Config{
		ServiceName: "order-service",
		Namespace:   "http",
		Subsystem:   "server",
		MetricsPath: "/metrics",
		Registerer:  reg,
		Gatherer:    reg,
		// Tailored latency histogram buckets for HTTP API latencies (in seconds)
		RequestDurationBuckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5},
		// Skip internal health probes from skewing request metrics
		SkipURIs: []string{"/health"},
	}))

	// API Routes
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Fiber v3 Prometheus monitoring service",
			"status":  "ok",
		})
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "healthy",
		})
	})

	app.Post("/orders", func(c fiber.Ctx) error {
		// Simulate business order processing
		metrics.OrdersProcessed.Inc()

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  "created",
			"message": "Order processed successfully",
		})
	})

	return app
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	app := SetupApp(prometheus.NewRegistry())
	// Listen returns once in-flight requests are drained, or after Fiber's 10s ShutdownTimeout.
	err := app.Listen(":3000", fiber.ListenConfig{GracefulContext: ctx})
	stop()
	if err != nil {
		log.Fatal(err)
	}
}
