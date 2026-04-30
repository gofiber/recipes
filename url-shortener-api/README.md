---
title: URL Shortener
keywords: [url shortener, redis, api]
description: URL shortening service with a simple API.
---

# URL Shortener API

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/url-shortener-api) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/url-shortener-api)

This project provides a URL shortening service built with [Fiber](https://gofiber.io) and Redis.

## Tech Stack

- Go + [Fiber v3](https://gofiber.io)
- Redis (via [go-redis/v8](https://github.com/go-redis/redis))

## Environment Variables

Copy `api/.env.example` to `api/.env` and adjust the values:

| Variable    | Default          | Description                                    |
|-------------|------------------|------------------------------------------------|
| `DB_ADDR`   | `localhost:6379` | Redis address (`host:port`)                    |
| `DB_PASS`   | _(empty)_        | Redis password (leave empty if none)           |
| `APP_PORT`  | `:3000`          | Port the API listens on (include leading colon)|
| `DOMAIN`    | `localhost:3000` | Public domain used to build the short URL      |
| `API_QUOTA` | `10`             | Max API calls per IP per 30-minute window      |

## Quick Start

### Redis only (local dev)

Start Redis with Docker Compose:

```sh
docker compose up db -d
```

Copy and edit the env file, then run the API:

```sh
cp api/.env.example api/.env
cd api && go run .
```

### Full stack (API + Redis)

```sh
docker compose up -d
```

## API Documentation

**Endpoint:** `POST http://localhost:3000/api/v1`

### Request body

| Field   | Type   | Required | Description                              |
|---------|--------|----------|------------------------------------------|
| `url`   | string | yes      | The original URL to shorten              |
| `short` | string | no       | Custom alias (auto-generated if omitted) |
| `expiry`| int    | no       | Expiry in hours (default: 24)            |

### Response body

| Field              | Type   | Description                              |
|--------------------|--------|------------------------------------------|
| `url`              | string | Original URL                             |
| `short`            | string | Full short URL including domain          |
| `expiry`           | int    | Expiry in hours                          |
| `rate_limit`       | int    | Remaining API calls in current window    |
| `rate_limit_reset` | int    | Minutes until rate limit window resets   |

> Rate limit: 10 calls per IP every 30 minutes (configurable via `API_QUOTA` in `.env`).

### curl Examples

**Shorten a URL (auto-generated alias):**

```sh
curl -X POST http://localhost:3000/api/v1 \
  -H "Content-Type: application/json" \
  -d '{"url": "https://gofiber.io", "expiry": 24}'
```

**Shorten a URL with a custom alias:**

```sh
curl -X POST http://localhost:3000/api/v1 \
  -H "Content-Type: application/json" \
  -d '{"url": "https://gofiber.io", "short": "fiber", "expiry": 48}'
```

**Resolve a short URL (browser or curl):**

```sh
curl -L http://localhost:3000/fiber
```

## Setup
