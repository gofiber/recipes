---
title: Svelte Netlify
keywords: [netlify, deploy, svelte]
description: Deploying a Svelte + Fiber application on Netlify.
---

# Svelte + Fiber on Netlify

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/svelte-netlify) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/svelte-netlify)

[![Netlify Status](https://api.netlify.com/api/v1/badges/143c3c42-60f7-427a-b3fd-8ca3947a2d40/deploy-status)](https://app.netlify.com/sites/gofiber-svelte/deploys)

A Go Fiber API deployed as a Netlify Function, with a Svelte frontend. The API provides IP geolocation via [ip-api.com](http://ip-api.com).

**Demo:** https://gofiber-svelte.netlify.app/

## Prerequisites

- **Go** 1.21+
- **Node.js** 18+ (for the Svelte frontend)
- **Netlify CLI** — `npm install -g netlify-cli`

## Local Development

1. Install frontend dependencies and build the Svelte app:

   ```bash
   npm install
   npm run build
   ```

2. Build the Go function:

   ```bash
   ./build.sh
   ```

3. Start the local dev server:

   ```bash
   netlify dev
   ```

   The app will be available at `http://localhost:8888`.

## Deploy to Netlify

### Via Netlify CLI

```bash
netlify deploy --prod
```

### Via Git

Connect your repository in the Netlify dashboard. The `netlify.toml` configures the build automatically.

## How It Works

- `./build.sh` compiles the Go binary to `cmd/gateway/gateway` and places it in the `functions/` directory.
- Netlify serves the binary as a [Netlify Function](https://functions.netlify.com/).
- Static files under `public/` are served directly (entry point: `index.html`).
- API calls to `/api/*` are redirected server-side to `/.netlify/functions/gateway/:splat` (status 200).

## Project Structure

```
.
├── adapter/        # AWS Lambda <-> Fiber adapter
├── cmd/gateway/    # Lambda entry point (main.go)
├── handler/        # Fiber route handlers
├── public/         # Compiled Svelte frontend
├── build.sh        # Build script for the Go function
└── netlify.toml    # Netlify build configuration
```

## Notes

- Netlify Functions are limited to 125,000 requests/month on the free tier (~2.89 req/min). Response caching with a 10-minute TTL is applied to the geolocation endpoint to stay within limits.
