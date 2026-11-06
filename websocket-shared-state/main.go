package main

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/contrib/v3/websocket/event"
	"github.com/gofiber/fiber/v3"
)

// SessionState holds the shared mutable state for a given session/entity.
type SessionState struct {
	ID        string            `json:"id"`
	Data      map[string]string `json:"data"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// StateManager orchestrates thread-safe access to shared session state
// and coordinates real-time broadcasts across connected WebSocket clients.
type StateManager struct {
	stateMu  sync.RWMutex
	sessions map[string]*SessionState

	clientMu sync.RWMutex
	clients  map[string]map[string]bool // sessionID -> set of client UUIDs
}

// NewStateManager creates an initialized thread-safe StateManager.
func NewStateManager() *StateManager {
	return &StateManager{
		sessions: make(map[string]*SessionState),
		clients:  make(map[string]map[string]bool),
	}
}

// Get returns a thread-safe snapshot copy of the session state.
func (sm *StateManager) Get(sessionID string) (*SessionState, bool) {
	sm.stateMu.RLock()
	defer sm.stateMu.RUnlock()

	state, exists := sm.sessions[sessionID]
	if !exists {
		return nil, false
	}

	// Return a deep copy to prevent race conditions with concurrent readers/writers
	dataCopy := make(map[string]string, len(state.Data))
	maps.Copy(dataCopy, state.Data)

	return &SessionState{
		ID:        state.ID,
		Data:      dataCopy,
		UpdatedAt: state.UpdatedAt,
	}, true
}

// Set updates a key-value pair in the session state.
// Crucially, it clones string inputs to prevent memory corruption caused by
// Fiber/fasthttp's zero-allocation request buffer reuse (see gofiber/fiber#3195).
func (sm *StateManager) Set(sessionID, key, value string) *SessionState {
	// Guard against zero-allocation buffer reuse: clone string slices into permanent heap memory
	cleanID := strings.Clone(sessionID)
	cleanKey := strings.Clone(key)
	cleanVal := strings.Clone(value)

	sm.stateMu.Lock()
	state, exists := sm.sessions[cleanID]
	if !exists {
		state = &SessionState{
			ID:   cleanID,
			Data: make(map[string]string),
		}
		sm.sessions[cleanID] = state
	}
	state.Data[cleanKey] = cleanVal
	state.UpdatedAt = time.Now().UTC()

	// Create snapshot for return and broadcast
	dataCopy := make(map[string]string, len(state.Data))
	maps.Copy(dataCopy, state.Data)
	snapshot := &SessionState{
		ID:        cleanID,
		Data:      dataCopy,
		UpdatedAt: state.UpdatedAt,
	}
	sm.stateMu.Unlock()

	// Broadcast snapshot to all active WebSocket clients subscribed to this session
	sm.Broadcast(cleanID, snapshot)

	return snapshot
}

// Register attaches an active WebSocket connection UUID to a session.
func (sm *StateManager) Register(sessionID, uuid string) {
	cleanID := strings.Clone(sessionID)
	cleanUUID := strings.Clone(uuid)
	sm.clientMu.Lock()
	defer sm.clientMu.Unlock()

	if _, exists := sm.clients[cleanID]; !exists {
		sm.clients[cleanID] = make(map[string]bool)
	}
	sm.clients[cleanID][cleanUUID] = true
}

// Unregister removes a disconnected WebSocket connection UUID from a session.
func (sm *StateManager) Unregister(sessionID, uuid string) {
	sm.clientMu.Lock()
	defer sm.clientMu.Unlock()

	if uuids, exists := sm.clients[sessionID]; exists {
		delete(uuids, uuid)
		if len(uuids) == 0 {
			delete(sm.clients, sessionID)
		}
	}
}

// Broadcast serializes state and pushes it to every active WebSocket connection for sessionID via event.EmitToList.
func (sm *StateManager) Broadcast(sessionID string, state *SessionState) {
	msg, err := json.Marshal(state)
	if err != nil {
		return
	}

	sm.clientMu.RLock()
	uuids := make([]string, 0, len(sm.clients[sessionID]))
	for u := range sm.clients[sessionID] {
		uuids = append(uuids, u)
	}
	sm.clientMu.RUnlock()

	if len(uuids) > 0 {
		event.EmitToList(uuids, msg, websocket.TextMessage)
	}
}

// UpdatePayload represents an incoming state modification request.
type UpdatePayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SetupApp constructs the Fiber application configured with shared state routes.
func SetupApp(customManager ...*StateManager) (*fiber.App, *StateManager) {
	var manager *StateManager
	if len(customManager) > 0 && customManager[0] != nil {
		manager = customManager[0]
	} else {
		manager = NewStateManager()
	}

	app := fiber.New()

	// 1. WebSocket Upgrade Guard
	app.Use("/ws", func(c fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// 2. WebSocket Route: Subscribes to session updates & allows real-time state mutation using contrib/v3/websocket/event
	app.Get("/ws/:id", event.New(func(kws *event.Websocket) {
		sessionID := kws.Params("id")
		uuid := kws.GetUUID()
		manager.Register(sessionID, uuid)

		kws.On(event.EventDisconnect, func(*event.EventPayload) {
			manager.Unregister(sessionID, uuid)
		})

		kws.On(event.EventMessage, func(ep *event.EventPayload) {
			var payload UpdatePayload
			if err := json.Unmarshal(ep.Data, &payload); err == nil && payload.Key != "" {
				manager.Set(sessionID, payload.Key, payload.Value)
			}
		})

		// Send initial state snapshot on connect if session already exists
		if state, exists := manager.Get(sessionID); exists {
			if msg, err := json.Marshal(state); err == nil {
				kws.Emit(msg, websocket.TextMessage)
			}
		}
	}))

	// 3. REST API: Read shared session state
	app.Get("/api/sessions/:id", func(c fiber.Ctx) error {
		sessionID := c.Params("id")
		state, exists := manager.Get(sessionID)
		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "session not found",
			})
		}
		return c.JSON(state)
	})

	// 4. REST API: Mutate shared session state & broadcast to WebSocket subscribers
	app.Post("/api/sessions/:id", func(c fiber.Ctx) error {
		sessionID := c.Params("id")
		var payload UpdatePayload
		if err := c.Bind().JSON(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}
		if payload.Key == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "key is required",
			})
		}

		// Mutate state and broadcast
		updated := manager.Set(sessionID, payload.Key, payload.Value)
		return c.Status(fiber.StatusOK).JSON(updated)
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	return app, manager
}

func main() {
	app, _ := SetupApp()
	fmt.Println("WebSocket Shared State Server listening on :3000")
	log.Fatal(app.Listen(":3000"))
}
