---
title: Advanced Caching
keywords: [cache, caching, invalidation, multi-tenant, performance]
description: Demonstrates advanced caching strategies in Fiber v3 including deterministic key generation, tenant isolation, client bypass, and tag-based invalidation.
---

# Advanced Caching in Fiber v3

This recipe illustrates production-ready caching patterns using Fiber v3's built-in `middleware/cache`:

1. **Deterministic & Isolated Cache Keys**: Uses native `KeyHeaders` (`Accept`, `Accept-Encoding`, `Accept-Language`, `X-Tenant-ID`, `X-User-Role`) to partition cache keys with built-in header escaping and canonical query sorting (`?a=1&b=2` matches `?b=2&a=1`), preventing delimiter collisions and cross-tenant pollution.
2. **Native Cache-Control Support**: Leverages Fiber's built-in `middleware/cache` support for `Cache-Control: no-cache, no-store` and HTTP method constraints (caching only GET/HEAD requests by default) without custom wrappers.
3. **Tag-Based Invalidation**: Each tag (the `category` query, `general` when absent) has a generation that is part of the cache key via `KeyHeaders`. Invalidating a tag bumps its generation, so its next request misses and refills the cache while other tags keep their entries. Old entries are never served again and expire after `Expiration`. `CacheInvalidator` is not used for this, because it only sees the request, not the stored entry, and cannot tell stale entries from fresh ones.
4. **Cache Status Introspection**: Return `X-Cache-Status` (`hit` or `miss`) headers for monitoring and observability.

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

```bash
export ADMIN_API_KEY=change-me
go run main.go
```

The application starts on port `:3000`. It refuses to start without `ADMIN_API_KEY`, which protects the invalidation endpoint.

### 1. Requesting Cached Endpoints

First request returns a cache miss (`X-Cache-Status: miss`):

```bash
curl -i -H "X-Tenant-ID: acme" -H "X-User-Role: member" "http://localhost:3000/api/items?category=books"
```

Subsequent identical request returns a cache hit (`X-Cache-Status: hit`):

```bash
curl -i -H "X-Tenant-ID: acme" -H "X-User-Role: member" "http://localhost:3000/api/items?category=books"
```

### 2. Tenant & Role Isolation

Requests with differing `X-Tenant-ID` or `X-User-Role` headers are partitioned into independent cache spaces:

```bash
curl -i -H "X-Tenant-ID: beta-corp" -H "X-User-Role: member" "http://localhost:3000/api/items?category=books"
# Returns X-Cache-Status: miss
```

### 3. Invalidation by Tag

Invalidate all entries tagged with `books` (requires the admin key as a bearer token):

```bash
curl -i -X POST "http://localhost:3000/api/items/invalidate" \
  -H "Authorization: Bearer $ADMIN_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"tag": "books"}'
```

The next request to `books` returns a cache miss and refills the cache; the ones after it hit again.

The generations live in process memory, so with several instances each one keeps its own; put them into shared storage (e.g. Redis) to invalidate across instances.

## Running Tests

```bash
go test -v ./...
```
