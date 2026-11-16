package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	fiberprom "github.com/gofiber/contrib/v3/prometheus"
	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// AppMetrics holds custom application-level business metrics.
type AppMetrics struct {
	OrdersProcessed prometheus.Counter
	ActiveWorkers   prometheus.Gauge
}

// NewAppMetrics initializes and registers domain metrics with the provided Prometheus registry.
func NewAppMetrics(reg prometheus.Registerer) *AppMetrics {
	m := &AppMetrics{
		OrdersProcessed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "app",
			Name:      "orders_processed_total",
			Help:      "Total number of orders successfully processed by the application.",
		}),
		ActiveWorkers: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "active_workers",
			Help:      "Current count of active background processing workers.",
		}),
	}

	reg.MustRegister(m.OrdersProcessed)
	reg.MustRegister(m.ActiveWorkers)

	return m
}

// SetupApp configures the Fiber application with Prometheus HTTP middleware and routes.
func SetupApp(reg *prometheus.Registry) (*fiber.App, *AppMetrics) {
	if reg == nil {
		reg = prometheus.NewRegistry()
	}

	// Register standard Go runtime and process metrics
	reg.MustRegister(collectors.NewGoCollector())
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	// Register domain-specific metrics
	metrics := NewAppMetrics(reg)

	app := fiber.New(fiber.Config{
		AppName: "Fiber-Prometheus-Example v3",
	})

	// Configure official Fiber Prometheus middleware
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

	return app, metrics
}

func main() {
	reg := prometheus.NewRegistry()
	app, _ := SetupApp(reg)

	// Graceful shutdown channel
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting HTTP server on :3000...")
		if err := app.Listen(":3000"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server listen failed: %v", err)
		}
	}()

	<-shutdown
	log.Println("Gracefully shutting down Fiber server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	log.Println("Server gracefully stopped.")
}
