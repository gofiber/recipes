---
title: Proxy Resilience
keywords: [proxy, retry, circuit breaker, backoff, resilience]
description: Implementing retry and circuit breaker patterns in proxy middleware.
---

# Proxy Resilience (Retry and Circuit Breaker)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/proxy-resilience) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/proxy-resilience)

This recipe shows how to make a Fiber v3 proxy route resilient with:

1. **Exponential backoff retry** (`github.com/gofiber/fiber/v3/addon/retry`) for transient network errors and upstream 5xx responses.
2. **A circuit breaker** ([`github.com/gofiber/contrib/v3/circuitbreaker`](https://github.com/gofiber/contrib/tree/main/v3/circuitbreaker)) that fails fast with `503 Service Unavailable` while the upstream is down, so a failing dependency is not hammered further.

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

## Features

- **Circuit breaker**: `circuitbreaker.Middleware` opens after `FailureThreshold` failed requests, answers `503` while open, and after `Timeout` lets one probe request through (`429` for the rest) to decide whether to close again. A request counts as failed when it ends with a 5xx status.
- **Retry with backoff**: `addon/retry` retries with a configurable `InitialInterval`, `Multiplier` and `MaxRetryCount`. Only idempotent methods are retried, so a `POST` never runs twice upstream.
- **Clean error responses**: when the retries run out, the upstream response is discarded and the client gets a `502` with a JSON error.
- **Status endpoint**: `cb.HealthHandler()` at `/circuit/status`.

> **Note**: `proxy.Do` refuses upstreams that resolve to private, loopback or link-local addresses, which protects against SSRF. Internal services, such as Kubernetes service IPs, therefore need `proxy.WithSecurityPolicy(proxy.SecurityPolicy{AllowPrivateIPs: true})`. Enable it only when the upstream address is not controlled by the client, as in this recipe, where it is fixed in `main.go`.

## Quickstart

### 1. Run the application

```bash
go run .
```

The service listens on `:3000` and proxies to `https://httpbin.org`.

### 2. Check the circuit status

```bash
curl -i http://localhost:3000/circuit/status
```

Response:

```json
{"healthy":true,"state":"closed"}
```

While the circuit is open, the endpoint answers `503` with `"state":"open"`.

### 3. Proxy a request to the upstream

```bash
curl -i "http://localhost:3000/proxy/get?hello=fiber"
```

## Running Tests

```bash
go test -v ./...
```

The tests use `httptest` upstreams and the breaker's `Clock` option, so the half-open recovery runs without waiting for the timeout.
