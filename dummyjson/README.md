---
title: Dummy JSON Proxy
keywords: [dummyjson, proxy, json, server]
---

# Simple Fiber Proxy Server

This is a basic Go application using the Fiber framework to create a web server. The server listens on port 3000 and has a single route (`GET /`) that fetches data from an external URL (`https://dummyjson.com/products/1`) and forwards it to the client.

### How to Run

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go run main.go`.
4. Visit `http://localhost:3000/` in a web browser or use a tool like `curl` to test it.

### What It Does

- Fetches data from an external service, in this case `DummyJson.com`
- Forwards the fetched data or an error message to the client.

### Error Handling

- Returns a 500 Internal Server Error if any issue occurs during the fetch.
- Returns the same status code as the external service if it's not a 200 OK.
