package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/addon/retry"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

func TestMain(m *testing.M) {
	// Enable private/loopback IPs for httptest local test servers
	old := proxy.WithSecurityPolicy(proxy.SecurityPolicy{AllowPrivateIPs: true})
	defer proxy.WithSecurityPolicy(old)
	os.Exit(m.Run())
}

func Test_Proxy_Success(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok from upstream"))
	}))
	defer upstream.Close()

	cb := NewCircuitBreaker(CircuitBreakerConfig{MaxFailures: 3, Timeout: 1 * time.Second})
	backoff := retry.NewExponentialBackoff(retry.Config{
		InitialInterval: 10 * time.Millisecond,
		MaxRetryCount:   2,
	})

	app := setupApp(upstream.URL, cb, backoff)

	req := httptest.NewRequest(http.MethodGet, "/proxy/test", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected app.Test error: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "ok from upstream" {
		t.Fatalf("expected 'ok from upstream', got %q", string(body))
	}

	if cb.State() != StateClosed {
		t.Fatalf("expected circuit CLOSED, got %s", cb.State())
	}
}

func Test_Proxy_Retry_OnTransientError(t *testing.T) {
	var attempts int32

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		current := atomic.AddInt32(&attempts, 1)
		if current < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("transient error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("recovered"))
	}))
	defer upstream.Close()

	cb := NewCircuitBreaker(CircuitBreakerConfig{MaxFailures: 3, Timeout: 1 * time.Second})
	backoff := retry.NewExponentialBackoff(retry.Config{
		InitialInterval: 5 * time.Millisecond,
		MaxBackoffTime:  50 * time.Millisecond,
		Multiplier:      1.5,
		MaxRetryCount:   3,
	})

	app := setupApp(upstream.URL, cb, backoff)

	req := httptest.NewRequest(http.MethodGet, "/proxy/flaky", nil)
	resp, err := app.Test(req, fiber.TestConfig{Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("unexpected app.Test error: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 after retries, got %d", resp.StatusCode)
	}

	if atomic.LoadInt32(&attempts) != 3 {
		t.Fatalf("expected 3 attempts, got %d", attempts)
	}

	if cb.State() != StateClosed {
		t.Fatalf("expected circuit to remain CLOSED, got %s", cb.State())
	}
}

func Test_CircuitBreaker_TripsOpen(t *testing.T) {
	var upstreamHits int32

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&upstreamHits, 1)
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte("upstream down"))
	}))
	defer upstream.Close()

	cb := NewCircuitBreaker(CircuitBreakerConfig{MaxFailures: 2, Timeout: 500 * time.Millisecond})
	backoff := retry.NewExponentialBackoff(retry.Config{
		InitialInterval: 5 * time.Millisecond,
		MaxRetryCount:   1, // 1 try per call
	})

	app := setupApp(upstream.URL, cb, backoff)

	// Call 1: fails
	req1 := httptest.NewRequest(http.MethodGet, "/proxy/fail", nil)
	resp1, _ := app.Test(req1)
	_ = resp1.Body.Close()
	if resp1.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected 502 Bad Gateway, got %d", resp1.StatusCode)
	}

	// Call 2: fails, trips circuit open
	req2 := httptest.NewRequest(http.MethodGet, "/proxy/fail", nil)
	resp2, _ := app.Test(req2)
	_ = resp2.Body.Close()
	if resp2.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected 502 Bad Gateway, got %d", resp2.StatusCode)
	}

	if cb.State() != StateOpen {
		t.Fatalf("expected circuit OPEN, got %s", cb.State())
	}

	hitsBeforeFastFail := atomic.LoadInt32(&upstreamHits)

	// Call 3: should fail fast with 503 without contacting upstream
	req3 := httptest.NewRequest(http.MethodGet, "/proxy/fail", nil)
	resp3, _ := app.Test(req3)
	defer func() { _ = resp3.Body.Close() }()

	if resp3.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 Service Unavailable, got %d", resp3.StatusCode)
	}

	if atomic.LoadInt32(&upstreamHits) != hitsBeforeFastFail {
		t.Fatalf("circuit open did not fail fast; upstream hits increased from %d to %d",
			hitsBeforeFastFail, atomic.LoadInt32(&upstreamHits))
	}
}

func Test_CircuitBreaker_HalfOpen_Recovery(t *testing.T) {
	var failUpstream int32 = 1

	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if atomic.LoadInt32(&failUpstream) == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("healthy"))
	}))
	defer upstream.Close()

	timeout := 100 * time.Millisecond
	cb := NewCircuitBreaker(CircuitBreakerConfig{MaxFailures: 1, Timeout: timeout})
	backoff := retry.NewExponentialBackoff(retry.Config{
		InitialInterval: 5 * time.Millisecond,
		MaxRetryCount:   1,
	})

	app := setupApp(upstream.URL, cb, backoff)

	// Cause failure to trip circuit
	req := httptest.NewRequest(http.MethodGet, "/proxy/call", nil)
	resp, _ := app.Test(req)
	_ = resp.Body.Close()

	if cb.State() != StateOpen {
		t.Fatalf("expected circuit OPEN, got %s", cb.State())
	}

	// Wait past timeout for Half-Open window
	time.Sleep(timeout + 20*time.Millisecond)

	// Mark upstream healthy
	atomic.StoreInt32(&failUpstream, 0)

	// Next request should transition through Half-Open and close circuit on success
	req2 := httptest.NewRequest(http.MethodGet, "/proxy/call", nil)
	resp2, _ := app.Test(req2)
	defer func() { _ = resp2.Body.Close() }()

	if resp2.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK on trial request, got %d", resp2.StatusCode)
	}

	if cb.State() != StateClosed {
		t.Fatalf("expected circuit to recover to CLOSED, got %s", cb.State())
	}
}

func Test_Circuit_Status_Endpoint(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{MaxFailures: 3, Timeout: 1 * time.Second})
	backoff := retry.NewExponentialBackoff(retry.Config{MaxRetryCount: 1})
	app := setupApp("http://example.com", cb, backoff)

	req := httptest.NewRequest(http.MethodGet, "/circuit/status", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected app.Test error: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var body map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&body)

	if body["state"] != "CLOSED" {
		t.Fatalf("expected status CLOSED, got %s", body["state"])
	}
}
