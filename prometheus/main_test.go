package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrometheus_EndpointExposition(t *testing.T) {
	reg := prometheus.NewRegistry()
	app, _ := SetupApp(reg)

	// 1. Initial request to application root
	reqRoot := httptest.NewRequest(http.MethodGet, "/", nil)
	respRoot, err := app.Test(reqRoot)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, respRoot.StatusCode)

	// 2. Fetch Prometheus metrics endpoint
	reqMetrics := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	respMetrics, err := app.Test(reqMetrics)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, respMetrics.StatusCode)
	assert.Contains(t, respMetrics.Header.Get("Content-Type"), "text/plain")

	body, err := io.ReadAll(respMetrics.Body)
	require.NoError(t, err)
	metricsOutput := string(body)

	// Verify standard HTTP metrics exposed by gofiber/contrib/v3/prometheus
	assert.Contains(t, metricsOutput, "http_server_requests_total")
	assert.Contains(t, metricsOutput, "http_server_request_duration_seconds_bucket")
	assert.Contains(t, metricsOutput, `service="order-service"`)

	// Verify Go runtime collectors are present
	assert.Contains(t, metricsOutput, "go_goroutines")
	assert.Contains(t, metricsOutput, "go_memstats_alloc_bytes")
}

func TestPrometheus_CustomBusinessMetrics(t *testing.T) {
	reg := prometheus.NewRegistry()
	app, _ := SetupApp(reg)

	// Process two orders
	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(http.MethodPost, "/orders", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	}

	// Verify custom domain metric increment in Prometheus scrape output
	reqMetrics := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	respMetrics, err := app.Test(reqMetrics)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, respMetrics.StatusCode)

	body, err := io.ReadAll(respMetrics.Body)
	require.NoError(t, err)
	metricsOutput := string(body)

	assert.Contains(t, metricsOutput, "app_orders_processed_total 2")
}

func TestPrometheus_SkipURIs_HealthExemption(t *testing.T) {
	reg := prometheus.NewRegistry()
	app, _ := SetupApp(reg)

	// Probe /health endpoint multiple times
	for i := 0; i < 3; i++ {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	}

	// Check metrics output to confirm /health was skipped from request instrumentation
	reqMetrics := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	respMetrics, err := app.Test(reqMetrics)
	require.NoError(t, err)

	body, err := io.ReadAll(respMetrics.Body)
	require.NoError(t, err)
	metricsOutput := string(body)

	// Path label for /health must not be present in requests_total
	assert.False(t, strings.Contains(metricsOutput, `path="/health"`))
}
