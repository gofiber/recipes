# Proxy Resilience (Retry and Circuit Breaker)

[![Go Reference](https://pkg.go.dev/badge/github.com/gofiber/recipes/proxy-resilience.svg)](https://pkg.go.dev/github.com/gofiber/recipes/proxy-resilience)

An enterprise-ready recipe demonstrating how to build resilient proxy architectures in **Fiber v3** using:
1. **Exponential Backoff Retry** (`github.com/gofiber/fiber/v3/addon/retry`) to automatically retry transient upstream network blips and 5xx errors.
2. **Circuit Breaker** pattern (`Closed` ➔ `Open` ➔ `Half-Open`) to fail fast with HTTP `503 Service Unavailable` when an upstream dependency degrades or crashes, protecting downstream resources from cascading failure.

---

## Architecture Overview

```
                      +-------------------+
                      |   Client Request  |
                      +---------+---------+
                                |
                                v
                      +---------+---------+
                      |   Fiber Proxy     |
                      +---------+---------+
                                |
                   [Circuit Breaker Check]
                                |
            +-------------------+-------------------+
            | (State: Closed)                       | (State: Open)
            v                                       v
   [Execute with Backoff Retry]           [Fail Fast: HTTP 503]
            |
    +-------+-------+
    | (Attempt 1-N) |
    v               v
 [Upstream OK]   [Upstream 5xx / Timeout]
 (HTTP 200)               |
                 (Failures >= Threshold)
                          v
                 [Trip Circuit OPEN]
```

---

## Features

- **Stateful Circuit Breaker**: Tracks consecutive failures (`MaxFailures`), transitions to `Open` state upon threshold breach, and enters `Half-Open` state after cooldown duration (`Timeout`).
- **Configurable Backoff**: Leverages Fiber's built-in `addon/retry` with customizable `InitialInterval`, `Multiplier`, and `MaxRetryCount`.
- **SSRF & Private IP Awareness**: Demonstrates safe upstream proxying using Fiber v3's `proxy.Do`.
- **Inspection Endpoint**: Real-time circuit breaker status check at `/circuit/status`.

---

## Quickstart

### 1. Run the application

```bash
go run .
```

The service will start listening on `:3000`.

### 2. Check Circuit Status

```bash
curl -i http://localhost:3000/circuit/status
```

Response:
```json
{
  "state": "CLOSED"
}
```

### 3. Proxy a Request to Upstream

```bash
curl -i http://localhost:3000/proxy/get
```

---

## Running Tests

Run the complete regression test suite:

```bash
go test -v ./...
```
