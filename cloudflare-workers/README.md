# Cloudflare Container Workers with Go Fiber

This example demonstrates how to use [Go Fiber v3](https://github.com/gofiber/fiber) with [Cloudflare Container Workers](https://developers.cloudflare.com/containers/).

## Features

- **Go Fiber v3** framework
- **Distroless container** for minimal attack surface (`gcr.io/distroless/static-debian12`)
- **JSON API** response
- **Environment variables** support
- **Logger and Recover** middleware

## Prerequisites

- Bun
- Go 1.24+
- Wrangler CLI
- Cloudflare account with Container Workers access

## Getting Started

1. Install dependencies:
```bash
bun install
```

2. Run locally:
```bash
bun run dev
```

3. Deploy to Cloudflare:
```bash
bun run deploy
```

## Project Structure

```
.
├── src/index.ts           # Worker entry point
├── container_src/
│   ├── main.go           # Go Fiber application
│   ├── go.mod            # Go module file
│   └── go.sum            # Go dependencies
├── Dockerfile            # Container configuration
└── wrangler.jsonc        # Cloudflare Workers configuration
```

## How it Works

1. The Worker (TypeScript) receives HTTP requests
2. Requests are forwarded to the Go Fiber container
3. The container responds with JSON data including environment variables

## Container Configuration

The container is configured with:
- 2 minute sleep timeout for inactivity
- Environment variable `MESSAGE` passed from the container class
- Port 8080 (default)

## Learn More

- [Fiber Documentation](https://docs.gofiber.io/)
- [Cloudflare Container Workers](https://developers.cloudflare.com/containers/)
- [Cloudflare Workers](https://developers.cloudflare.com/workers/)
