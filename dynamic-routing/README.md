---
title: Dynamic Route Management
keywords: [dynamic routing, route manager, app.stack, runtime routes, fiber v3]
description: Managing dynamic route registration via an in-memory registry and a catch-all dispatcher in Fiber v3.
---

# Dynamic Route Management

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/dynamic-routing) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/dynamic-routing)

This recipe demonstrates how to manage dynamic routes at runtime in Fiber v3 without restarting the application.

> **Note on Router Thread-Safety**: Fiber v3 can also add and remove real routes at runtime with `app.RebuildTree()` and `app.RemoveRoute()`, but those calls are not thread-safe while the server handles requests. This recipe keeps its routes in a `sync.RWMutex`-protected registry instead and serves them through one `/d/*` handler, so the router itself never changes.

It provides:
- A thread-safe runtime `RouteManager` using `sync.RWMutex`.
- Admin management endpoints to register, list, and unregister dynamic routes, protected by the `keyauth` middleware.
- Route stack inspection using Fiber's native `app.Stack()`.
- A dynamic dispatcher route (`/d/*`) to evaluate and execute active routes.

> **Note**: The admin routes require the API key from `ADMIN_API_KEY` as a bearer token, because anyone who can call them can publish content under `/d/*`. The server refuses to start without it.

## Prerequisites

Ensure you have the following installed:

- Go 1.26.0 or newer
- [Fiber v3](https://github.com/gofiber/fiber) package (`github.com/gofiber/fiber/v3`)

## Setup

1. Navigate to the recipe directory:
   ```sh
   cd recipes/dynamic-routing
   ```

2. Download and verify dependencies:
   ```sh
   go mod tidy
   ```

## Running the Application

1. In your terminal, start the server:
   ```sh
   export ADMIN_API_KEY=change-me
   go run main.go
   ```

2. In a separate terminal, test the endpoints using `curl`:

   **List active routes (static stack and dynamic registry):**
   ```sh
   curl -s http://127.0.0.1:3000/admin/routes \
     -H "Authorization: Bearer $ADMIN_API_KEY" | jq .
   ```

   **Access pre-seeded dynamic route:**
   ```sh
   curl -s http://127.0.0.1:3000/d/promo | jq .
   ```

   **Register a new dynamic route at runtime:**
   ```sh
   curl -s -X POST http://127.0.0.1:3000/admin/routes \
     -H "Authorization: Bearer $ADMIN_API_KEY" \
     -H "Content-Type: application/json" \
     -d '{"name": "flash-sale", "path": "/flash", "response": "Limited time offer 50% off!"}' | jq .
   ```

   The path must start with `/`. Registering a path that already exists returns `409 Conflict`.

   **Execute the newly registered dynamic route:**
   ```sh
   curl -s http://127.0.0.1:3000/d/flash | jq .
   ```

   **Remove the dynamic route:**
   ```sh
   curl -s -X DELETE http://127.0.0.1:3000/admin/routes/flash \
     -H "Authorization: Bearer $ADMIN_API_KEY" | jq .
   ```

## Example Code

```go
package main

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v3"
)

type DynamicRoute struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Response string `json:"response"`
}

type RouteManager struct {
	mu     sync.RWMutex
	routes map[string]DynamicRoute
}

// See main.go for full implementation
```
