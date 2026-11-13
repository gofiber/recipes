package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	fastws "github.com/fasthttp/websocket"
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

// waitForSubscriber blocks until the server has registered a connection for sessionID,
// so a mutation sent right after Dial is not broadcast to nobody.
func waitForSubscriber(t *testing.T, sm *StateManager, sessionID string) {
	t.Helper()
	require.Eventually(t, func() bool {
		sm.clientMu.RLock()
		defer sm.clientMu.RUnlock()
		return len(sm.clients[sessionID]) > 0
	}, 2*time.Second, 5*time.Millisecond)
}

func TestSharedState_WebSocketBroadcastDelivery(t *testing.T) {
	app, manager := SetupApp()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer ln.Close()

	go func() {
		_ = app.Listener(ln)
	}()

	addr := ln.Addr().String()
	wsURL := fmt.Sprintf("ws://%s/ws/room-live", addr)

	// Dial WebSocket client
	conn, resp, err := fastws.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn.Close()
	if resp != nil && resp.Body != nil {
		_ = resp.Body.Close()
	}
	waitForSubscriber(t, manager, "room-live")

	// Send REST mutation
	postBody, err := json.Marshal(UpdatePayload{
		Key:   "broadcast_key",
		Value: "delivered_val",
	})
	require.NoError(t, err)

	httpResp, err := http.Post(fmt.Sprintf("http://%s/api/sessions/room-live", addr), "application/json", bytes.NewReader(postBody))
	require.NoError(t, err)
	defer httpResp.Body.Close()
	assert.Equal(t, http.StatusOK, httpResp.StatusCode)

	// Assert WebSocket receives the broadcast
	_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	_, msg, err := conn.ReadMessage()
	require.NoError(t, err)

	var received SessionState
	err = json.Unmarshal(msg, &received)
	require.NoError(t, err)
	assert.Equal(t, "room-live", received.ID)
	assert.Equal(t, "delivered_val", received.Data["broadcast_key"])
}

func TestSharedState_ConcurrentWebSocketBroadcast(t *testing.T) {
	app, manager := SetupApp()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer ln.Close()

	go func() {
		_ = app.Listener(ln)
	}()

	addr := ln.Addr().String()
	wsURL := fmt.Sprintf("ws://%s/ws/room-concurrent", addr)

	conn, resp, err := fastws.DefaultDialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn.Close()
	if resp != nil && resp.Body != nil {
		_ = resp.Body.Close()
	}
	waitForSubscriber(t, manager, "room-concurrent")

	const numMessages = 50
	var wg sync.WaitGroup
	wg.Add(numMessages)

	for i := 0; i < numMessages; i++ {
		go func(idx int) {
			defer wg.Done()
			payload, _ := json.Marshal(UpdatePayload{
				Key:   fmt.Sprintf("key-%d", idx),
				Value: fmt.Sprintf("val-%d", idx),
			})
			r, postErr := http.Post(fmt.Sprintf("http://%s/api/sessions/room-concurrent", addr), "application/json", bytes.NewReader(payload))
			if postErr == nil {
				_ = r.Body.Close()
			}
		}(i)
	}

	receivedCount := 0
	var latest SessionState
	_ = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	for receivedCount < numMessages {
		_, msg, readErr := conn.ReadMessage()
		if readErr != nil {
			break
		}
		receivedCount++

		var snapshot SessionState
		require.NoError(t, json.Unmarshal(msg, &snapshot))
		if snapshot.Version > latest.Version {
			latest = snapshot
		}
	}

	wg.Wait()
	assert.Equal(t, numMessages, receivedCount)

	// Frames may arrive out of order; the highest version is the current state.
	final, ok := manager.Get("room-concurrent")
	require.True(t, ok)
	assert.Equal(t, uint64(numMessages), final.Version)
	assert.Equal(t, final.Version, latest.Version)
	assert.Equal(t, final.Data, latest.Data)
}
