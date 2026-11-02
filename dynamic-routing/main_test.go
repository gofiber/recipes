package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func Test_DynamicRouteLifecycle(t *testing.T) {
	rm := newRouteManager()
	app := setupApp(rm)

	// 1. Initial lookup on dynamic endpoint (should be 404)
	req1 := httptest.NewRequest(http.MethodGet, "/d/features", nil)
	resp1, err := app.Test(req1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp1.Body.Close()
	if resp1.StatusCode != fiber.StatusNotFound {
		t.Errorf("expected 404, got %d", resp1.StatusCode)
	}

	// 2. Register new dynamic route via POST /admin/routes
	payload := DynamicRoute{
		Name:     "beta-features",
		Path:     "/features",
		Response: "Beta features enabled for v3",
	}
	payloadBytes, _ := json.Marshal(payload)

	req2 := httptest.NewRequest(http.MethodPost, "/admin/routes", bytes.NewReader(payloadBytes))
	req2.Header.Set("Content-Type", "application/json")
	resp2, err := app.Test(req2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != fiber.StatusCreated {
		t.Errorf("expected 201 Created, got %d", resp2.StatusCode)
	}

	// 3. Request dynamic route again (should now be 200 OK)
	req3 := httptest.NewRequest(http.MethodGet, "/d/features", nil)
	resp3, err := app.Test(req3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp3.Body.Close()
	if resp3.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200 OK, got %d", resp3.StatusCode)
	}

	body3, _ := io.ReadAll(resp3.Body)
	var resMap map[string]interface{}
	if err := json.Unmarshal(body3, &resMap); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}
	if resMap["response"] != "Beta features enabled for v3" {
		t.Errorf("unexpected response body: %v", string(body3))
	}

	// 4. List all routes via GET /admin/routes
	req4 := httptest.NewRequest(http.MethodGet, "/admin/routes", nil)
	resp4, err := app.Test(req4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp4.Body.Close()
	if resp4.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200 OK for list, got %d", resp4.StatusCode)
	}

	// 5. Delete dynamic route via DELETE /admin/routes/features
	req5 := httptest.NewRequest(http.MethodDelete, "/admin/routes/features", nil)
	resp5, err := app.Test(req5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp5.Body.Close()
	if resp5.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200 OK on deletion, got %d", resp5.StatusCode)
	}

	// 6. Verify dynamic route is removed (404 again)
	req6 := httptest.NewRequest(http.MethodGet, "/d/features", nil)
	resp6, err := app.Test(req6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp6.Body.Close()
	if resp6.StatusCode != fiber.StatusNotFound {
		t.Errorf("expected 404 after deletion, got %d", resp6.StatusCode)
	}
}

func Test_HealthzAndRoot(t *testing.T) {
	rm := newRouteManager()
	app := setupApp(rm)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	reqRoot := httptest.NewRequest(http.MethodGet, "/", nil)
	respRoot, err := app.Test(reqRoot)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer respRoot.Body.Close()
	if respRoot.StatusCode != fiber.StatusOK {
		t.Errorf("expected 200, got %d", respRoot.StatusCode)
	}
}
