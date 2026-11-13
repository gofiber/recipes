---
title: Angular SSR
keywords: [angular, ssr, server-side-rendering, universal, proxy, spa, fiber-v3]
description: Serving Angular with Server-Side Rendering (SSR) using Fiber v3 and reverse proxy.
---

# Angular SSR with Fiber v3

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/angular-ssr)

This recipe demonstrates how to put **Go Fiber v3** in front of an **Angular Server-Side Rendering (SSR)** application. The Angular app in `web/` was created with `ng new web --ssr` and trimmed to a single page that tells you where it was rendered.

## Architecture

1. **Static Browser Assets**: Built frontend files (`*.js`, `*.css`, `favicon.ico`, ...) are served directly by Fiber from `web/dist/web/browser` via `middleware/static`. Only paths with a file extension are treated as assets, so every page (including `/`) reaches SSR.
2. **Page Routes**: All other requests are forwarded via `middleware/proxy` to the Angular SSR Node.js server (`http://127.0.0.1:4000`) with a bounded timeout. Its responses, including its own 404 pages, are passed through.
3. **Backend API Endpoints**: Fiber handles native Go API routes under `/api/*`.
4. **Resilience / Fallback**: If the SSR server is offline or answers with a 5xx, Fiber serves Angular's client-side rendering shell `index.csr.html` instead, and the page renders in the browser.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.26 or higher
- [Node.js](https://nodejs.org/) 24 and npm

## Project Structure

```
.
├── web/                         # Angular workspace (ng new web --ssr)
│   ├── angular.json
│   ├── package.json
│   └── src/
│       ├── app/                 # App component and SSR route config
│       ├── main.ts              # Browser entry point
│       ├── main.server.ts       # SSR entry point
│       └── server.ts            # Node.js (Express) SSR server
├── go.mod
├── go.sum
├── main.go                      # Fiber server & SSR proxy configuration
├── main_test.go
├── Makefile
└── README.md
```

## Running the Application

All commands run in the `angular-ssr` directory.

### 1. Build

```bash
make build
```

This installs the npm dependencies, builds the Angular app into `web/dist/web` (`browser/` for the client bundle, `server/` for the SSR server) and builds the Go binary.

### 2. Start the Angular SSR server

```bash
node web/dist/web/server/server.mjs
```

It listens on port `4000` (set `PORT` to change it). Angular only accepts requests whose `Host` header is listed in `security.allowedHosts` in `web/angular.json`. The Fiber proxy sends the host of `SSR_ENDPOINT`, so `127.0.0.1` is allowed there; add your host if you point `SSR_ENDPOINT` elsewhere.

### 3. Start Fiber

```bash
# Optional: custom SSR backend (an empty value serves only the client-side shell) and port
export SSR_ENDPOINT="http://127.0.0.1:4000"
export PORT="3000"

go run .
```

Open `http://localhost:3000`. The HTML comes back server-rendered:

```bash
curl -s http://localhost:3000/ | grep "Rendered on"
#   <p>Rendered on the <strong>server</strong>.</p>
```

In the browser the page switches to "browser" once Angular has hydrated it. Stop the Node.js server and reload: Fiber falls back to `index.csr.html` and the page is rendered in the browser only.

- Health check: `http://localhost:3000/api/health`
- Static assets: `http://localhost:3000/favicon.ico` and the hashed `main-*.js` / `styles-*.css` files

## Running Tests

The Go tests use a temporary directory as browser bundle and a mock SSR server, so they do not need Node.js:

```bash
go test -v ./...
```
