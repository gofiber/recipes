package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestApp(t *testing.T) *sentry.MockTransport {
	transport := &sentry.MockTransport{}
	opts := SentryOptions("")
	opts.Transport = transport
	err := sentry.Init(opts)
	require.NoError(t, err)
	return transport
}

func TestSentry_SuccessfulRequest(t *testing.T) {
	transport := setupTestApp(t)
	app := SetupApp()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// No Sentry events for 200 OK
	assert.Empty(t, transport.Events())
}

func TestSentry_ClientErrorIgnored(t *testing.T) {
	transport := setupTestApp(t)
	app := SetupApp()

	// 1. Test 400 Bad Request
	req400 := httptest.NewRequest(http.MethodGet, "/users/invalid", nil)
	resp400, err := app.Test(req400)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp400.StatusCode)

	// 2. Test 404 Not Found
	req404 := httptest.NewRequest(http.MethodGet, "/route-that-does-not-exist", nil)
	resp404, err := app.Test(req404)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp404.StatusCode)

	// Sentry must not capture 4xx client errors
	assert.Empty(t, transport.Events())
}

func TestSentry_ServerErrorCaptured(t *testing.T) {
	transport := setupTestApp(t)
	app := SetupApp()

	req := httptest.NewRequest(http.MethodGet, "/error", nil)
	req.Header.Set("X-User-ID", "user-123")

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var result map[string]any
	err = json.Unmarshal(bodyBytes, &result)
	require.NoError(t, err)

	assert.Equal(t, float64(http.StatusInternalServerError), result["status"])
	assert.Equal(t, "Internal Server Error", result["error"])
	assert.NotContains(t, string(bodyBytes), "database connection failed")

	eventID, ok := result["event_id"].(string)
	require.True(t, ok)
	assert.NotEmpty(t, eventID)

	// Verify Sentry event captured with request URL and user context
	events := transport.Events()
	require.Len(t, events, 1)

	event := events[0]
	assert.Equal(t, string(event.EventID), eventID)
	assert.Equal(t, "user-123", event.User.ID)
	require.NotNil(t, event.Request)
	assert.Equal(t, "GET", event.Request.Method)
	require.NotEmpty(t, event.Exception)
	assert.Contains(t, event.Exception[0].Value, "database connection failed")
}

func TestSentry_PanicRecovery(t *testing.T) {
	transport := setupTestApp(t)
	app := SetupApp()

	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	req.Header.Set("X-User-ID", "user-admin")

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var result map[string]any
	err = json.Unmarshal(bodyBytes, &result)
	require.NoError(t, err)

	assert.Equal(t, float64(http.StatusInternalServerError), result["status"])
	assert.Equal(t, "Internal Server Error", result["error"])

	eventID, ok := result["event_id"].(string)
	require.True(t, ok)
	assert.NotEmpty(t, eventID)

	// Verify panic captured by Sentry exactly once with stack trace
	events := transport.Events()
	require.Len(t, events, 1)

	event := events[0]
	assert.Equal(t, string(event.EventID), eventID)
	assert.Equal(t, "user-admin", event.User.ID)
	assert.Equal(t, sentry.LevelFatal, event.Level)
	assert.Contains(t, event.Message, "nil pointer dereference")
	require.NotNil(t, event.Request)
	assert.Equal(t, "GET", event.Request.Method)

	require.NotEmpty(t, event.Threads)
	require.NotNil(t, event.Threads[0].Stacktrace)
	hasHandlerFrame := false
	for _, frame := range event.Threads[0].Stacktrace.Frames {
		if strings.HasPrefix(frame.Function, "SetupApp.") {
			hasHandlerFrame = true
			break
		}
	}
	assert.True(t, hasHandlerFrame, "stacktrace should include SetupApp handler frame")
}
