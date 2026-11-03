package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
	"time"

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

func Test_RootRouteSendsEarlyHints(t *testing.T) {
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	app := setupApp()
	go func() { _ = app.Listener(ln, fiber.ListenConfig{DisableStartupMessage: true}) }()
	defer func() { _ = app.Shutdown() }()

	conn, err := net.Dial("tcp4", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Write([]byte("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n")); err != nil {
		t.Fatal(err)
	}
	br := bufio.NewReader(conn)
	resp, err := http.ReadResponse(br, nil) // interim 103
	if err != nil {
		t.Fatal(err)
	}
	wantLinks := []string{
		"</css/main.css>; rel=preload; as=style",
		"</js/bundle.js>; rel=preload; as=script",
	}
	if resp.StatusCode != http.StatusEarlyHints || !slices.Equal(resp.Header.Values("Link"), wantLinks) {
		t.Fatalf("got %d %v, want 103 with %v", resp.StatusCode, resp.Header.Values("Link"), wantLinks)
	}
	if resp, err = http.ReadResponse(br, nil); err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("final response: %v %v", resp, err)
	}
}
