---
title: Angular SSR
keywords: [angular, ssr, server-side-rendering, universal, proxy, spa, fiber-v3]
description: Serving Angular with Server-Side Rendering (SSR) using Fiber v3 and reverse proxy.
---

# Angular SSR with Fiber v3

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/angular-ssr)

This recipe demonstrates how to integrate **Angular Server-Side Rendering (SSR / Angular Universal)** with **Go Fiber v3**.

## Architecture

In an Angular SSR deployment:
1. **Static Browser Assets**: Built frontend chunks (`*.js`, `*.css`, fonts, images) are served directly by Fiber via high-performance zero-copy static middleware (`middleware/static`).
2. **Dynamic / Page Routes**: HTML requests are routed via reverse proxy (`middleware/proxy`) to the Angular SSR Node.js server (`http://127.0.0.1:4000`) with a bounded timeout, ensuring full SEO and fast initial page load.
3. **Backend API Endpoints**: Fiber handles native Go API routes under `/api/*` with sub-millisecond response times.
4. **Resilience / Fallback**: If the SSR server is offline or fails, Fiber automatically falls back to serving the client-side SPA `index.html`.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.26 or higher
- [Node.js](https://nodejs.org/) (v20+ LTS) and [Angular CLI](https://angular.dev/) (optional, if building your own Angular project from scratch)

## Project Structure

```
.
├── dist/
│   └── browser/        # Built Angular browser assets
│       ├── index.html
│       ├── main.js
│       └── styles.css
├── go.mod
├── go.sum
├── main.go             # Fiber server & SSR proxy configuration
├── main_test.go        # Unit and integration test suite
└── README.md
```

## Running the Application

Always navigate into the `angular-ssr` recipe directory first:

```bash
cd recipes/angular-ssr
```

### 1. Run Unit Tests

```bash
go test -v ./...
```

### 2. Building or Providing Browser Assets

A ready-to-run sample build is included under `./dist/browser`. If you have an existing Angular SSR project:
```bash
# In your Angular application workspace:
ng build --configuration production
# Place the compiled artifacts into dist/browser:
cp -r dist/<project-name>/browser/* ./dist/browser/
```

### 3. Running Angular SSR Backend (Optional)

In production or local testing with Node.js SSR:
```bash
# Start your Angular SSR Node server (defaults to port 4000):
node dist/<project-name>/server/server.mjs
```

### 4. Start the Fiber Application

```bash
# Optional: specify custom SSR backend (or leave empty for standalone SPA mode) and port
export SSR_ENDPOINT="http://127.0.0.1:4000"
export PORT="3000"

go run main.go
```

The application will be accessible at `http://localhost:3000`.
- Health check: `http://localhost:3000/api/health`
- Static assets: `http://localhost:3000/main.js`, `http://localhost:3000/styles.css`
- Dynamic routes: Forwarded to Angular SSR backend (or SPA `index.html` fallback).
