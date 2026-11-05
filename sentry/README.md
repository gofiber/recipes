---
title: Custom Error Handling with Sentry
keywords: [sentry, error-handling, recovery, monitoring, observability]
description: Demonstrates custom error handling and panic recovery in Fiber v3 integrated with Sentry.
---

# Custom Error Handling with Sentry in Fiber v3

This recipe demonstrates production-ready error handling and monitoring patterns in Fiber v3 integrated with [Sentry](https://sentry.io):

1. **Per-Request Sentry Hub**: Leverages official `github.com/gofiber/contrib/v3/sentry` middleware to clone a dedicated Sentry Hub for each request with deep-copied request data and authenticated user identity (`X-User-ID`).
2. **Selective Error Reporting**: Custom `fiber.ErrorHandler` captures internal server errors (`status >= 500`) to Sentry while avoiding alert noise from expected client errors (`4xx`, such as validation failures or 404s).
3. **Panic Recovery**: Catches unexpected runtime panics, forwards the fatal stack trace to Sentry, and returns a graceful 500 JSON response to the client.
4. **Sentry Event ID Correlation**: Injects the Sentry `event_id` into error responses so clients can reference specific incident IDs in support inquiries or logs.

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

Set your `SENTRY_DSN` environment variable and run:

```bash
export SENTRY_DSN="https://your-dsn@sentry.io/project-id"
go run main.go
```

The application starts on port `:3000`.

### 1. Normal Request (200 OK)

```bash
curl -i http://localhost:3000/
```

Returns:
```json
{"message":"Fiber with Sentry Error Handling","status":"ok"}
```

### 2. Client-Side Error (400 Bad Request)

Expected client errors return clean JSON without triggering Sentry issues:

```bash
curl -i http://localhost:3000/users/invalid
```

Returns HTTP 400:
```json
{"error":"invalid user id provided","status":400}
```

### 3. Server-Side Error (500 Internal Server Error)

Unhandled internal exceptions are captured in Sentry and return an `event_id`:

```bash
curl -i -H "X-User-ID: user-42" http://localhost:3000/error
```

Returns HTTP 500:
```json
{
  "error": "database connection failed: timeout connecting to replica",
  "event_id": "4b6840742f104d49a239be9b6b797b5e",
  "status": 500
}
```

### 4. Runtime Panic Recovery

Unexpected panics are caught by `SentryRecover()`, reported with stack traces to Sentry, and resolved safely:

```bash
curl -i -H "X-User-ID: user-admin" http://localhost:3000/panic
```

Returns HTTP 500:
```json
{
  "error": "Internal Server Error",
  "event_id": "9a1f28b4929841f2bb0da93510e19034",
  "status": 500
}
```

## Running Tests

Run the unit test suite (utilizing Sentry's in-memory mock transport):

```bash
go test -v ./...
```
