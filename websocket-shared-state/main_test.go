package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSharedState_RESTWorkflow(t *testing.T) {
	app, _ := SetupApp()

	// 1. Initial GET on non-existent session returns 404
	reqGet := httptest.NewRequest(http.MethodGet, "/api/sessions/sess-1", nil)
	respGet, err := app.Test(reqGet)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, respGet.StatusCode)

	// 2. POST to create/update state
	bodyBytes, err := json.Marshal(UpdatePayload{
		Key:   "status",
		Value: "in-progress",
	})
	require.NoError(t, err)

	reqPost := httptest.NewRequest(http.MethodPost, "/api/sessions/sess-1", bytes.NewReader(bodyBytes))
	reqPost.Header.Set("Content-Type", "application/json")
	respPost, err := app.Test(reqPost)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, respPost.StatusCode)

	var postResult SessionState
	err = json.NewDecoder(respPost.Body).Decode(&postResult)
	require.NoError(t, err)
	assert.Equal(t, "sess-1", postResult.ID)
	assert.Equal(t, "in-progress", postResult.Data["status"])

	// 3. GET retrieves the stored state
	reqGet2 := httptest.NewRequest(http.MethodGet, "/api/sessions/sess-1", nil)
	respGet2, err := app.Test(reqGet2)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, respGet2.StatusCode)

	var getResult SessionState
	err = json.NewDecoder(respGet2.Body).Decode(&getResult)
	require.NoError(t, err)
	assert.Equal(t, "in-progress", getResult.Data["status"])
}

func TestSharedState_ZeroAllocationBufferSafety(t *testing.T) {
	app, _ := SetupApp()

	// Verify that rapid sequential updates from different pseudo-requests
	// retain independent, non-corrupted strings in the StateManager
	for i := 0; i < 50; i++ {
		sessID := fmt.Sprintf("tenant-%d", i)
		val := fmt.Sprintf("dynamic-val-%d", i)

		payload, _ := json.Marshal(UpdatePayload{
			Key:   "metric",
			Value: val,
		})

		req := httptest.NewRequest(http.MethodPost, "/api/sessions/"+sessID, bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	}

	// Verify each session retained its exact string without buffer overwrites
	for i := 0; i < 50; i++ {
		sessID := fmt.Sprintf("tenant-%d", i)
		expectedVal := fmt.Sprintf("dynamic-val-%d", i)

		req := httptest.NewRequest(http.MethodGet, "/api/sessions/"+sessID, nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var state SessionState
		err = json.NewDecoder(resp.Body).Decode(&state)
		require.NoError(t, err)
		assert.Equal(t, expectedVal, state.Data["metric"])
	}
}

func TestSharedState_ConcurrentAccess(t *testing.T) {
	sm := NewStateManager()
	var wg sync.WaitGroup

	const numWorkers = 20
	const iterations = 50

	// Concurrent Writers
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				sessionID := fmt.Sprintf("session-%d", workerID%4)
				key := fmt.Sprintf("key-%d", i%10)
				val := fmt.Sprintf("worker-%d-iter-%d", workerID, i)
				sm.Set(sessionID, key, val)
			}
		}(w)
	}

	// Concurrent Readers
	for r := 0; r < numWorkers; r++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				sessionID := fmt.Sprintf("session-%d", readerID%4)
				if state, exists := sm.Get(sessionID); exists {
					// Verify deep-copy: mutating the returned map must not affect internal state
					state.Data["tamper"] = "illegal"
				}
			}
		}(r)
	}

	wg.Wait()

	// Verify that no "tamper" key leaked into the internal state
	for s := 0; s < 4; s++ {
		sessionID := fmt.Sprintf("session-%d", s)
		state, exists := sm.Get(sessionID)
		if exists {
			_, hasTamper := state.Data["tamper"]
			assert.False(t, hasTamper, "internal state must be protected from external modification")
		}
	}
}

func TestSharedState_WebSocketUpgradeGuard(t *testing.T) {
	app, _ := SetupApp()

	// Plain HTTP GET to WebSocket endpoint must be rejected with 426 Upgrade Required
	req := httptest.NewRequest(http.MethodGet, "/ws/sess-1", nil)
	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUpgradeRequired, resp.StatusCode)
}

func TestSharedState_ValidationErrors(t *testing.T) {
	app, _ := SetupApp()

	// Missing body
	reqEmpty := httptest.NewRequest(http.MethodPost, "/api/sessions/sess-1", nil)
	reqEmpty.Header.Set("Content-Type", "application/json")
	respEmpty, err := app.Test(reqEmpty)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, respEmpty.StatusCode)

	// Missing key
	reqNoKey := httptest.NewRequest(http.MethodPost, "/api/sessions/sess-1", bytes.NewReader([]byte(`{"value":"test"}`)))
	reqNoKey.Header.Set("Content-Type", "application/json")
	respNoKey, err := app.Test(reqNoKey)
	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, respNoKey.StatusCode)

	body, _ := io.ReadAll(respNoKey.Body)
	assert.Contains(t, string(body), "key is required")
}
