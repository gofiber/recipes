package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gofiber/contrib/v3/circuitbreaker"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/addon/retry"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

func TestMain(m *testing.M) {
	// httptest upstreams listen on loopback, which proxy.Do refuses by default.
	proxy.WithSecurityPolicy(proxy.SecurityPolicy{AllowPrivateIPs: true})
	os.Exit(m.Run())
}

func fastBackoff(maxRetries int) *retry.ExponentialBackoff {
	return retry.NewExponentialBackoff(retry.Config{
		InitialInterval: 5 * time.Millisecond,
		MaxBackoffTime:  20 * time.Millisecond,
		Multiplier:      2,
		MaxRetryCount:   maxRetries,
	})
}

func send(t *testing.T, app *fiber.App, method, target string) *http.Response {
	t.Helper()
	resp, err := app.Test(httptest.NewRequest(method, target, nil), fiber.TestConfig{Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	t.Cleanup(func() { _ = resp.Body.Close() })
	return resp
}

func Test_Proxy_Success(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok from upstream"))
	}))
	defer upstream.Close()

	cb := circuitbreaker.New(circuitbreaker.Config{FailureThreshold: 3})
	resp := send(t, setupApp(upstream.URL, cb, fastBackoff(2)), http.MethodGet, "/proxy/test")

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || string(body) != "ok from upstream" {
		t.Fatalf("got %d %q, want 200 %q", resp.StatusCode, body, "ok from upstream")
	}
	if cb.GetState() != circuitbreaker.StateClosed {
		t.Fatalf("expected circuit closed, got %s", cb.GetState())
	}
}

func Test_Proxy_ForwardsQueryString(t *testing.T) {
	var query atomic.Value
	upstream := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		query.Store(r.URL.RawQuery)
	}))
	defer upstream.Close()

	app := setupApp(upstream.URL, circuitbreaker.New(circuitbreaker.Config{}), fastBackoff(1))
	send(t, app, http.MethodGet, "/proxy/get?a=1&b=2")

	if got := query.Load(); got != "a=1&b=2" {
		t.Fatalf("upstream got query %q, want %q", got, "a=1&b=2")
	}
}

func Test_Proxy_Retry_OnTransientError(t *testing.T) {
	var attempts atomic.Int32
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if attempts.Add(1) < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write([]byte("recovered"))
	}))
	defer upstream.Close()

	cb := circuitbreaker.New(circuitbreaker.Config{FailureThreshold: 3})
	resp := send(t, setupApp(upstream.URL, cb, fastBackoff(3)), http.MethodGet, "/proxy/flaky")

	if resp.StatusCode != http.StatusOK || attempts.Load() != 3 {
		t.Fatalf("got %d after %d attempts, want 200 after 3", resp.StatusCode, attempts.Load())
	}
	if cb.GetState() != circuitbreaker.StateClosed {
		t.Fatalf("expected circuit closed, got %s", cb.GetState())
	}
}

func Test_Proxy_RetryKeepsRequestBody(t *testing.T) {
	const payload = `{"name":"fiber"}`
	var attempts atomic.Int32
	bodies := make(chan string, 4)
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		bodies <- string(b)
		if attempts.Add(1) == 1 {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer upstream.Close()

	app := setupApp(upstream.URL, circuitbreaker.New(circuitbreaker.Config{}), fastBackoff(3))
	req := httptest.NewRequest(http.MethodPut, "/proxy/items/1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, fiber.TestConfig{Timeout: 2 * time.Second})
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK || attempts.Load() != 2 {
		t.Fatalf("got %d after %d attempts, want 200 after 2", resp.StatusCode, attempts.Load())
	}
	close(bodies)
	for b := range bodies {
		if b != payload {
			t.Fatalf("upstream got body %q, want %q on every attempt", b, payload)
		}
	}
}

func Test_Proxy_DoesNotRetryPost(t *testing.T) {
	var attempts atomic.Int32
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer upstream.Close()

	app := setupApp(upstream.URL, circuitbreaker.New(circuitbreaker.Config{FailureThreshold: 10}), fastBackoff(3))
	resp := send(t, app, http.MethodPost, "/proxy/orders")

	if resp.StatusCode != http.StatusBadGateway || attempts.Load() != 1 {
		t.Fatalf("got %d after %d attempts, want 502 after 1", resp.StatusCode, attempts.Load())
	}
}

func Test_Proxy_ErrorDropsUpstreamHeaders(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Set-Cookie", "upstream=1")
		w.WriteHeader(http.StatusInternalServerError)
		gz := gzip.NewWriter(w)
		_, _ = gz.Write([]byte("boom"))
		_ = gz.Close()
	}))
	defer upstream.Close()

	app := setupApp(upstream.URL, circuitbreaker.New(circuitbreaker.Config{}), fastBackoff(1))
	resp := send(t, app, http.MethodGet, "/proxy/fail")

	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("502 body is not plain JSON: %v", err)
	}
	if resp.Header.Get("Content-Encoding") != "" || resp.Header.Get("Set-Cookie") != "" {
		t.Fatalf("upstream headers leaked: %v", resp.Header)
	}
}

func Test_CircuitBreaker_TripsOpen(t *testing.T) {
	var hits atomic.Int32
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		hits.Add(1)
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer upstream.Close()

	cb := circuitbreaker.New(circuitbreaker.Config{FailureThreshold: 2})
	app := setupApp(upstream.URL, cb, fastBackoff(1))

	for range 2 {
		if resp := send(t, app, http.MethodGet, "/proxy/fail"); resp.StatusCode != http.StatusBadGateway {
			t.Fatalf("expected 502, got %d", resp.StatusCode)
		}
	}
	if cb.GetState() != circuitbreaker.StateOpen {
		t.Fatalf("expected circuit open, got %s", cb.GetState())
	}

	before := hits.Load()
	if resp := send(t, app, http.MethodGet, "/proxy/fail"); resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 503, got %d", resp.StatusCode)
	}
	if hits.Load() != before {
		t.Fatalf("open circuit reached the upstream: hits %d -> %d", before, hits.Load())
	}
}

func Test_CircuitBreaker_HalfOpen_Recovery(t *testing.T) {
	var failing atomic.Bool
	failing.Store(true)
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if failing.Load() {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer upstream.Close()

	start := time.Now()
	var elapsed atomic.Int64
	cb := circuitbreaker.New(circuitbreaker.Config{
		FailureThreshold: 1,
		Timeout:          time.Minute,
		Clock:            func() time.Time { return start.Add(time.Duration(elapsed.Load())) },
	})
	app := setupApp(upstream.URL, cb, fastBackoff(1))

	send(t, app, http.MethodGet, "/proxy/call")
	if cb.GetState() != circuitbreaker.StateOpen {
		t.Fatalf("expected circuit open, got %s", cb.GetState())
	}

	elapsed.Store(int64(2 * time.Minute))
	failing.Store(false)

	if resp := send(t, app, http.MethodGet, "/proxy/call"); resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 on the probe, got %d", resp.StatusCode)
	}
	if cb.GetState() != circuitbreaker.StateClosed {
		t.Fatalf("expected circuit closed, got %s", cb.GetState())
	}
}

func Test_Circuit_Status_Endpoint(t *testing.T) {
	app := setupApp("http://example.com", circuitbreaker.New(circuitbreaker.Config{}), fastBackoff(1))
	resp := send(t, app, http.MethodGet, "/circuit/status")

	var body map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body["state"] != "closed" || body["healthy"] != true {
		t.Fatalf("got %v, want closed and healthy", body)
	}
}
