---
title: Dummy JSON Proxy
keywords: [dummyjson, proxy, json, server]
description: Proxying dummy JSON data.
---

# Simple Fiber Proxy Server

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/dummyjson) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/dummyjson)

This is a basic Go application using the Fiber framework to create a proxy server. The server listens on port 3000 and has a single route (`GET /proxy`) that accepts a `?url=<target>` query parameter, fetches data from the given external URL, and forwards it to the client.

### How to Run

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go run main.go`.
4. Use a tool like `curl` to test it: `curl "http://localhost:3000/proxy?url=https://dummyjson.com/products/1"`

### What It Does

- Fetches data from an external service, in this case `DummyJson.com`
- Forwards the fetched data or an error message to the client.

### Error Handling

- Returns a 500 Internal Server Error if any issue occurs during the fetch.
- Returns the same status code as the external service if it's not a 200 OK.
