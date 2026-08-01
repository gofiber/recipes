---
title: Dummy JSON Proxy
keywords: [dummyjson, proxy, json, server]
description: Proxying dummy JSON data.
---

# Simple Fiber Proxy Server

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/dummyjson) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/dummyjson)

This is a basic Go application using the Fiber framework to create a proxy server. The server listens on port 3000 and has a single route (`GET /proxy`) that accepts an optional `?url=<target>` query parameter, fetches data from that URL, and forwards it to the client. Without the parameter it falls back to a default upstream.

> **Run this locally only.** The `url` parameter is forwarded to the HTTP client without validation, so anyone who can reach the route can make the server fetch arbitrary addresses on its behalf, including internal ones. That is a [server-side request forgery](https://owasp.org/API-Security/editions/2023/en/0xa7-server-side-request-forgery/) proxy. The example is kept minimal on purpose; before exposing anything like it, restrict the allowed schemes and hosts and reject private and loopback ranges.

## Prerequisites

Ensure you have the following installed:

- Go 1.25 or newer, required by Fiber v3
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/dummyjson
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

Start the server:

```sh
go run main.go
```

The server listens on port 3000.

## Usage

Call the route without parameters to use the default upstream:

```sh
curl "http://localhost:3000/proxy"
```

Pass a `url` query parameter to override it:

```sh
curl "http://localhost:3000/proxy?url=https://dummyjson.com/products/2"
```

Either way the server fetches the data from the external service and forwards the response to the client.

### Error Handling

- Returns 500 Internal Server Error if anything goes wrong during the fetch.
- Returns the same status code as the external service if it is not 200 OK.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [DummyJSON](https://dummyjson.com)
