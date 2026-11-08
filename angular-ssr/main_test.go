package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/gofiber/fiber/v3/middleware/proxy"
)

func TestMain(m *testing.M) {
	// SSR_ENDPOINT is set by the operator, not the client
	proxy.WithSecurityPolicy(proxy.SecurityPolicy{AllowPrivateIPs: true})
	os.Exit(m.Run())
}

func TestAPIHealth(t *testing.T) {
	app := App("")

	req := httptest.NewRequest("GET", "/api/health", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		t.Fatalf("failed to parse json: %v", err)
	}

	if data["status"] != "ok" {
		t.Errorf("status = %q, want %q", data["status"], "ok")
	}
}

func TestStaticAssetsDelivery(t *testing.T) {
	app := App("")

	// Request static JS asset
	req := httptest.NewRequest("GET", "/main.js", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	if !strings.Contains(string(body), "angular chunk") {
		t.Errorf("expected JS content, got %q", string(body))
	}
}

func TestSSRProxyForwarding(t *testing.T) {
	var mu sync.Mutex
	var receivedPath string

	// Mock Angular SSR Node backend recording exact path and query
	mockSSR := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		receivedPath = r.URL.RequestURI()
		mu.Unlock()

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<!DOCTYPE html><html><body><h1>Server Rendered Angular</h1></body></html>"))
	}))
	defer mockSSR.Close()

	app := App(mockSSR.URL)

	req := httptest.NewRequest("GET", "/dashboard?tab=overview", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	expected := "<!DOCTYPE html><html><body><h1>Server Rendered Angular</h1></body></html>"
	if string(body) != expected {
		t.Errorf("body = %q, want %q", string(body), expected)
	}

	mu.Lock()
	gotPath := receivedPath
	mu.Unlock()

	if gotPath != "/dashboard?tab=overview" {
		t.Errorf("mock received URI = %q, want %q", gotPath, "/dashboard?tab=overview")
	}
}

func TestSSRProxyFailureFallbackToSPA(t *testing.T) {
	// Pointing to a dead port so proxy immediately fails and triggers SPA fallback
	app := App("http://127.0.0.1:59999")

	req := httptest.NewRequest("GET", "/about", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	if !strings.Contains(string(body), "Client SPA") {
		t.Errorf("expected SPA fallback content, got %q", string(body))
	}
}

func TestSSRProxy5xxFallbackResetsHeaders(t *testing.T) {
	mockSSR := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Set-Cookie", "session=evil; Path=/")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"internal ssr error"}`))
	}))
	defer mockSSR.Close()

	app := App(mockSSR.URL)

	req := httptest.NewRequest("GET", "/failing-ssr", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	if enc := resp.Header.Get("Content-Encoding"); enc != "" {
		t.Errorf("Content-Encoding = %q, want empty", enc)
	}
	if cookie := resp.Header.Get("Set-Cookie"); cookie != "" {
		t.Errorf("Set-Cookie = %q, want empty", cookie)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	if !strings.Contains(string(body), "Client SPA") {
		t.Errorf("expected SPA fallback content, got %q", string(body))
	}
}

func TestResolveConfig(t *testing.T) {
	// Default case
	t.Setenv("SSR_ENDPOINT", "")
	t.Setenv("PORT", "")
	endpoint, port := resolveConfig()
	if endpoint != "" || port != "3000" {
		t.Errorf("default resolveConfig() = (%q, %q), want ('', '3000')", endpoint, port)
	}
}
