---
title: Early Hints (HTTP 103)
keywords: [early hints, http 103, preload, performance, assets]
description: Demonstrating HTTP 103 Early Hints in Fiber v3 to preload critical assets before delivering the final response.
---

# Early Hints (HTTP 103)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/early-hints) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/early-hints)

This recipe demonstrates how to send `HTTP 103 Early Hints` in Fiber v3 using `c.SendEarlyHints()`.

Early Hints ([RFC 8297](https://datatracker.ietf.org/doc/html/rfc8297)) allows the server to send preliminary response headers containing `Link` preload hints before the full response body is generated. Browsers supporting Early Hints can start fetching critical stylesheets, scripts, and fonts in parallel while the server executes database queries or business logic.

## Prerequisites

Ensure you have the following installed:

- Go 1.26.0 or newer
- [Fiber v3](https://github.com/gofiber/fiber) package (`github.com/gofiber/fiber/v3`)

> **Note**: HTTP 103 Early Hints operates most effectively over **HTTP/2** and **HTTP/3** connections. When requested over HTTP/1.1 or legacy environments, Fiber continues and safely serves the final response.

## Setup

1. Navigate to the recipe directory:
   ```sh
   cd recipes/early-hints
   ```

2. Download and verify dependencies:
   ```sh
   go mod tidy
   ```

## Running the Application

1. In your terminal, start the Fiber server:
   ```sh
   go run main.go
   ```

2. In a separate terminal, test the endpoints using `curl`:
   ```sh
   # Request root page
   curl -i http://127.0.0.1:3000/

   # Check health endpoint
   curl -i http://127.0.0.1:3000/healthz
   ```

## Example Code

```go
package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		// Assets for the client to preload
		hints := []string{
			"</css/main.css>; rel=preload; as=style",
			"</js/bundle.js>; rel=preload; as=script",
		}

		// Dispatch HTTP 103 Early Hints
		if err := c.SendEarlyHints(hints); err != nil {
			log.Printf("SendEarlyHints warning: %v", err)
		}

		// Simulate server work
		time.Sleep(50 * time.Millisecond)

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString("<!DOCTYPE html><html>...</html>")
	})

	log.Fatal(app.Listen(":3000"))
}
```
