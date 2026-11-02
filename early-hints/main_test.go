package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func Test_RootRoute(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	expectedSubstring := "Fiber Early Hints (HTTP 103)"
	if !strings.Contains(string(body), expectedSubstring) {
		t.Errorf("expected body to contain %q, got %q", expectedSubstring, string(body))
	}
}

func Test_StaticAssets(t *testing.T) {
	app := setupApp()

	cssReq := httptest.NewRequest(http.MethodGet, "/css/main.css", nil)
	cssResp, err := app.Test(cssReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer cssResp.Body.Close()

	if cssResp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status 200 for css, got %d", cssResp.StatusCode)
	}

	jsReq := httptest.NewRequest(http.MethodGet, "/js/bundle.js", nil)
	jsResp, err := app.Test(jsReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer jsResp.Body.Close()

	if jsResp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status 200 for js, got %d", jsResp.StatusCode)
	}
}

func Test_HealthzRoute(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected status 200 for /healthz, got %d", resp.StatusCode)
	}
}
