package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAdminKey = "test-admin-key"

func TestAdvancedCaching_HitMissAndHeaders(t *testing.T) {
	app, _ := SetupApp(testAdminKey)

	req1 := httptest.NewRequest(http.MethodGet, "/api/items?category=tech", nil)
	req1.Header.Set("X-Tenant-ID", "tenant-1")
	req1.Header.Set("X-User-Role", "developer")

	resp1, err := app.Test(req1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)
	assert.Equal(t, "miss", resp1.Header.Get("X-Cache-Status"))

	body1, err := io.ReadAll(resp1.Body)
	require.NoError(t, err)

	// Second request with same attributes should hit the cache
	req2 := httptest.NewRequest(http.MethodGet, "/api/items?category=tech", nil)
	req2.Header.Set("X-Tenant-ID", "tenant-1")
	req2.Header.Set("X-User-Role", "developer")

	resp2, err := app.Test(req2)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp2.StatusCode)
	assert.Equal(t, "hit", resp2.Header.Get("X-Cache-Status"))

	body2, err := io.ReadAll(resp2.Body)
	require.NoError(t, err)
	assert.Equal(t, string(body1), string(body2))
}

func TestAdvancedCaching_TenantAndRoleIsolation(t *testing.T) {
	app, _ := SetupApp(testAdminKey)

	// Tenant 1 request
	req1 := httptest.NewRequest(http.MethodGet, "/api/items?category=finance", nil)
	req1.Header.Set("X-Tenant-ID", "tenant-1")
	resp1, err := app.Test(req1)
	require.NoError(t, err)
	assert.Equal(t, "miss", resp1.Header.Get("X-Cache-Status"))

	// Tenant 2 request with identical URL should NOT hit Tenant 1's cache
	req2 := httptest.NewRequest(http.MethodGet, "/api/items?category=finance", nil)
	req2.Header.Set("X-Tenant-ID", "tenant-2")
	resp2, err := app.Test(req2)
	require.NoError(t, err)
	assert.Equal(t, "miss", resp2.Header.Get("X-Cache-Status"))

	// Collision resistance: (tenant c, role a:b) vs (tenant b:c, role a)
	reqCol1 := httptest.NewRequest(http.MethodGet, "/api/items?category=sec", nil)
	reqCol1.Header.Set("X-Tenant-ID", "c")
	reqCol1.Header.Set("X-User-Role", "a:b")
	respCol1, err := app.Test(reqCol1)
	require.NoError(t, err)
	assert.Equal(t, "miss", respCol1.Header.Get("X-Cache-Status"))

	reqCol2 := httptest.NewRequest(http.MethodGet, "/api/items?category=sec", nil)
	reqCol2.Header.Set("X-Tenant-ID", "b:c")
	reqCol2.Header.Set("X-User-Role", "a")
	respCol2, err := app.Test(reqCol2)
	require.NoError(t, err)
	assert.Equal(t, "miss", respCol2.Header.Get("X-Cache-Status"))

	// Query canonicalization: ?a=1&b=2 and ?b=2&a=1 must hit the same cache entry
	reqQ1 := httptest.NewRequest(http.MethodGet, "/api/items?a=1&b=2", nil)
	respQ1, err := app.Test(reqQ1)
	require.NoError(t, err)
	assert.Equal(t, "miss", respQ1.Header.Get("X-Cache-Status"))

	reqQ2 := httptest.NewRequest(http.MethodGet, "/api/items?b=2&a=1", nil)
	respQ2, err := app.Test(reqQ2)
	require.NoError(t, err)
	assert.Equal(t, "hit", respQ2.Header.Get("X-Cache-Status"))
}

func TestAdvancedCaching_RejectsRepeatedCategory(t *testing.T) {
	app, _ := SetupApp(testAdminKey)

	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/items?category=books&category=tech", nil))
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAdvancedCaching_BypassHeader(t *testing.T) {
	app, _ := SetupApp(testAdminKey)

	// 1. Initial request misses and populates cache
	req1 := httptest.NewRequest(http.MethodGet, "/api/items", nil)
	resp1, err := app.Test(req1)
	require.NoError(t, err)
	assert.Equal(t, "miss", resp1.Header.Get("X-Cache-Status"))

	body1, err := io.ReadAll(resp1.Body)
	require.NoError(t, err)

	time.Sleep(2 * time.Millisecond)

	// 2. Request with Cache-Control: no-cache bypasses stored entry, fetches fresh, and stores it
	reqBypass := httptest.NewRequest(http.MethodGet, "/api/items", nil)
	reqBypass.Header.Set("Cache-Control", "no-cache")
	respBypass, err := app.Test(reqBypass)
	require.NoError(t, err)
	assert.Equal(t, "miss", respBypass.Header.Get("X-Cache-Status"))

	bodyRefreshed, err := io.ReadAll(respBypass.Body)
	require.NoError(t, err)
	assert.NotEqual(t, string(body1), string(bodyRefreshed))

	// 3. Subsequent normal request hits the refreshed cache entry
	reqHit := httptest.NewRequest(http.MethodGet, "/api/items", nil)
	respHit, err := app.Test(reqHit)
	require.NoError(t, err)
	assert.Equal(t, "hit", respHit.Header.Get("X-Cache-Status"))

	bodyHit, err := io.ReadAll(respHit.Body)
	require.NoError(t, err)
	assert.Equal(t, string(bodyRefreshed), string(bodyHit))
}

func cacheStatus(t *testing.T, app *fiber.App, url string) string {
	t.Helper()
	resp, err := app.Test(httptest.NewRequest(http.MethodGet, url, nil))
	require.NoError(t, err)
	return resp.Header.Get("X-Cache-Status")
}

func invalidate(t *testing.T, app *fiber.App, tag, key string) *http.Response {
	t.Helper()
	invBody, err := json.Marshal(map[string]string{"tag": tag})
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/api/items/invalidate", bytes.NewReader(invBody))
	req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)
	if key != "" {
		req.Header.Set("Authorization", "Bearer "+key)
	}
	resp, err := app.Test(req)
	require.NoError(t, err)
	return resp
}

func TestAdvancedCaching_TagBasedInvalidation(t *testing.T) {
	app, _ := SetupApp(testAdminKey)

	// 1. Warm cache for categories "hardware" and "books"
	assert.Equal(t, "miss", cacheStatus(t, app, "/api/items?category=hardware"))
	assert.Equal(t, "hit", cacheStatus(t, app, "/api/items?category=hardware"))
	assert.Equal(t, "miss", cacheStatus(t, app, "/api/items?category=books"))

	// 2. Invalidate tag "hardware"
	invResp := invalidate(t, app, "hardware", testAdminKey)
	assert.Equal(t, http.StatusOK, invResp.StatusCode)

	var invResult map[string]any
	require.NoError(t, json.NewDecoder(invResp.Body).Decode(&invResult))
	assert.Equal(t, "hardware", invResult["invalidated_tag"])
	assert.Equal(t, "invalidated", invResult["status"])

	// 3. The next request refills the cache, and later ones hit it again
	assert.Equal(t, "miss", cacheStatus(t, app, "/api/items?category=hardware"))
	assert.Equal(t, "hit", cacheStatus(t, app, "/api/items?category=hardware"))

	// Other tags keep their entries
	assert.Equal(t, "hit", cacheStatus(t, app, "/api/items?category=books"))
}

func TestAdvancedCaching_InvalidationRequiresAPIKey(t *testing.T) {
	app, _ := SetupApp(testAdminKey)

	assert.Equal(t, "miss", cacheStatus(t, app, "/api/items?category=hardware"))

	assert.Equal(t, http.StatusUnauthorized, invalidate(t, app, "hardware", "").StatusCode)
	assert.Equal(t, http.StatusUnauthorized, invalidate(t, app, "hardware", "wrong-key").StatusCode)

	assert.Equal(t, "hit", cacheStatus(t, app, "/api/items?category=hardware"))
}
