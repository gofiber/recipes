---
title: WebSocket Shared State
keywords: [websocket, shared-state, concurrency, real-time, zero-allocation, broadcast]
description: Demonstrates thread-safe shared state management and real-time broadcasting across multiple WebSocket connections and HTTP handlers in Fiber v3.
---

# WebSocket Shared State in Fiber v3

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/websocket-shared-state)

This recipe demonstrates thread-safe shared state management and real-time state synchronization across multiple WebSocket connections and concurrent REST endpoints in Fiber v3.

## Architectural Highlights

1. **Thread-Safe State Manager**: Encapsulates session data with `sync.RWMutex` providing atomic mutations and read snapshots.
2. **Zero-Allocation Safety**: Explicitly clones request string slices (`strings.Clone`) to prevent memory corruption when storing context data in long-lived in-memory maps (addressing community report [gofiber/fiber#3195](https://github.com/gofiber/fiber/issues/3195)).
3. **Multi-Client State Synchronization**: Connect multiple WebSocket clients to the same session ID. State mutations (via REST or WebSocket) are immediately broadcast to all active subscribers.
4. **Clean Connection Lifecycle**: Handles client unregistration upon connection drops without memory leaks or race conditions.

## Project Structure

```
.
├── go.mod
├── go.sum
├── main.go
├── main_test.go
└── README.md
```

## Running the Recipe

Start the server:

```bash
go run main.go
```

The server listens on `http://localhost:3000`.

### 1. Connect WebSocket Client (Subscribed to Session)

Using `wscat` or your browser's WebSocket API:

```bash
# Connect client 1 to session 'room-101'
wscat -c ws://localhost:3000/ws/room-101
```

If the session already has state, the server immediately sends a state snapshot upon connection.

### 2. Read State via REST API

```bash
curl -i http://localhost:3000/api/sessions/room-101
```

Returns:
```json
{"id":"room-101","data":{"status":"active"},"updated_at":"2026-11-04T15:00:00Z"}
```

### 3. Mutate State via REST API (Broadcasts to WebSocket)

```bash
curl -i -X POST http://localhost:3000/api/sessions/room-101 \
  -H "Content-Type: application/json" \
  -d '{"key": "order_step", "value": "shipped"}'
```

Every connected WebSocket client subscribed to `room-101` immediately receives:
```json
{
  "id": "room-101",
  "data": {
    "order_step": "shipped"
  },
  "updated_at": "2026-11-04T15:00:01Z"
}
```

### 4. Mutate State via WebSocket Message

Send a JSON message through the `wscat` terminal:
```json
{"key": "user_status", "value": "typing"}
```

The updated state is stored in the `StateManager` and broadcast to all other peer connections on the session.

## Running Tests

Run the test suite with race detection enabled:

```bash
go test -v -race ./...
```
