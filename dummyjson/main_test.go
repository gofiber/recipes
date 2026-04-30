package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func newApp() *fiber.App {
	app := fiber.New()

	app.Get("/proxy", func(c fiber.Ctx) error {
		target := c.Query("url", defaultUpstreamURL)

		resp, err := client.Get(target) //nolint:noctx
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return c.Status(resp.StatusCode).JSON(&fiber.Map{
				"success": false,
				"error":   "upstream returned status " + http.StatusText(resp.StatusCode),
			})
		}

		c.Set(fiber.HeaderContentType, resp.Header.Get(fiber.HeaderContentType))

		if _, err := io.Copy(c.Response().BodyWriter(), resp.Body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return nil
	})

	return app
}

func TestProxyDefaultURL(t *testing.T) {
	// Use a local test server as the upstream so tests are hermetic.
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":1}`))
	}))
	defer upstream.Close()

	app := newApp()
	req := httptest.NewRequest(http.MethodGet, "/proxy?url="+upstream.URL, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != `{"id":1}` {
		t.Fatalf("unexpected body: %s", body)
	}
}

func TestProxyUpstreamError(t *testing.T) {
	// Upstream returns non-200.
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer upstream.Close()

	app := newApp()
	req := httptest.NewRequest(http.MethodGet, "/proxy?url="+upstream.URL, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func TestProxyInvalidURL(t *testing.T) {
	app := newApp()
	req := httptest.NewRequest(http.MethodGet, "/proxy?url=http://127.0.0.1:1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
