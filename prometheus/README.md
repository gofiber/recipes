---
title: Prometheus Monitoring
keywords: [prometheus, metrics, monitoring, grafana, observability, http]
description: Demonstrates how to instrument Fiber v3 applications and export Prometheus metrics.
---

# Prometheus Monitoring in Fiber v3

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/prometheus) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/prometheus)

This recipe demonstrates production-ready Prometheus metrics exposition in Fiber v3 using the official [`github.com/gofiber/contrib/v3/prometheus`](https://github.com/gofiber/contrib/tree/main/v3/prometheus) middleware:

1. **HTTP Metrics Instrumentation**: Automatically tracks request duration histograms, request counts by method, path, and status code, and in-flight requests.
2. **Dedicated `/metrics` Endpoint**: Serves metrics in the standard Prometheus text exposition format.
3. **Custom Application / Domain Metrics**: Registers a custom business metric (`app_orders_processed_total` Counter) with a shared Prometheus registry, next to the Go runtime and process metrics the middleware registers.
4. **Endpoint Exemption (`SkipURIs`)**: Excludes internal health probes (`/health`) from skewing request latency histograms and count metrics.
5. **Graceful Shutdown**: On `SIGINT` or `SIGTERM`, `ListenConfig.GracefulContext` drains active requests before the process exits (bounded by Fiber's default 10s `ShutdownTimeout`).

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

Navigate to the recipe directory and run:

```bash
cd prometheus
go run main.go
```

The server starts on port `:3000`.

### 1. Process an Order (Increments Business Metric)

```bash
curl -i -X POST http://localhost:3000/orders
```

Returns:
```json
{"message":"Order processed successfully","status":"created"}
```

### 2. Inspect Prometheus Metrics

```bash
curl -s http://localhost:3000/metrics
```

Output includes both Fiber HTTP server metrics and custom application metrics:

```text
# HELP app_orders_processed_total Total number of orders successfully processed by the application.
# TYPE app_orders_processed_total counter
app_orders_processed_total 1

# HELP http_server_requests_total Count all http requests by status code, method and path.
# TYPE http_server_requests_total counter
http_server_requests_total{method="POST",path="/orders",service="order-service",status_code="201"} 1

# HELP http_server_request_duration_seconds Duration of all HTTP requests by status code, method and path.
# TYPE http_server_request_duration_seconds histogram
http_server_request_duration_seconds_bucket{method="POST",path="/orders",service="order-service",status_code="201",le="0.005"} 1
...
```

### 3. Check Health Probe Exemption

Health checks at `/health` return `200 OK`, but do not pollute the Prometheus latency histograms:

```bash
curl -i http://localhost:3000/health
```

## Running Tests

Run the unit test suite:

```bash
go test -v ./...
```
