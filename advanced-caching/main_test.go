package main

import (
	"bytes"
	"context"
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

// memStorage implements fiber.Storage in-memory for testing purposes.
type memStorage struct {
	data map[string][]byte
}

func newMemStorage() *memStorage {
	return &memStorage{data: make(map[string][]byte)}
}

func (m *memStorage) Get(key string) ([]byte, error) {
	val, ok := m.data[key]
	if !ok {
		return nil, nil
	}
	return val, nil
}

func (m *memStorage) GetWithContext(_ context.Context, key string) ([]byte, error) {
	return m.Get(key)
}

func (m *memStorage) Set(key string, val []byte, exp time.Duration) error {
	m.data[key] = val
	return nil
}

func (m *memStorage) SetWithContext(_ context.Context, key string, val []byte, exp time.Duration) error {
	return m.Set(key, val, exp)
}

func (m *memStorage) Delete(key string) error {
	delete(m.data, key)
	return nil
}

func (m *memStorage) DeleteWithContext(_ context.Context, key string) error {
	return m.Delete(key)
}

func (m *memStorage) Reset() error {
	m.data = make(map[string][]byte)
	return nil
}

func (m *memStorage) ResetWithContext(_ context.Context) error {
	return m.Reset()
}

func (m *memStorage) Close() error {
	return nil
}

func TestAdvancedCaching_HitMissAndHeaders(t *testing.T) {
	storage := newMemStorage()
	app, _ := SetupApp(AppConfig{Storage: storage})

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
	storage := newMemStorage()
	app, _ := SetupApp(AppConfig{Storage: storage})

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

func TestAdvancedCaching_BypassHeader(t *testing.T) {
	storage := newMemStorage()
	app, _ := SetupApp(AppConfig{Storage: storage})

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

func TestAdvancedCaching_TagBasedInvalidation(t *testing.T) {
	storage := newMemStorage()
	app, _ := SetupApp(AppConfig{Storage: storage})

	// 1. Warm cache for category "hardware"
	req1 := httptest.NewRequest(http.MethodGet, "/api/items?category=hardware", nil)
	resp1, err := app.Test(req1)
	require.NoError(t, err)
	assert.Equal(t, "miss", resp1.Header.Get("X-Cache-Status"))

	// Verify second request hits
	req2 := httptest.NewRequest(http.MethodGet, "/api/items?category=hardware", nil)
	resp2, err := app.Test(req2)
	require.NoError(t, err)
	assert.Equal(t, "hit", resp2.Header.Get("X-Cache-Status"))

	// 2. Invalidate tag "hardware"
	invBody, _ := json.Marshal(map[string]string{"tag": "hardware"})
	invReq := httptest.NewRequest(http.MethodPost, "/api/items/invalidate", bytes.NewReader(invBody))
	invReq.Header.Set("Content-Type", fiber.MIMEApplicationJSON)

	invResp, err := app.Test(invReq)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, invResp.StatusCode)

	var invResult map[string]interface{}
	err = json.NewDecoder(invResp.Body).Decode(&invResult)
	require.NoError(t, err)
	assert.Equal(t, "hardware", invResult["invalidated_tag"])
	assert.Equal(t, "purged", invResult["status"])

	// 3. Next request should be cache miss again due to CacheInvalidator
	req3 := httptest.NewRequest(http.MethodGet, "/api/items?category=hardware", nil)
	resp3, err := app.Test(req3)
	require.NoError(t, err)
	assert.Equal(t, "miss", resp3.Header.Get("X-Cache-Status"))
}
