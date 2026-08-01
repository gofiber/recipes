---
title: Dummy JSON Proxy
keywords: [dummyjson, proxy, json, server]
description: Proxying dummy JSON data.
---

# Simple Fiber Proxy Server

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/dummyjson) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/dummyjson)

This is a basic Go application using the Fiber framework to create a proxy server. The server listens on port 3000 and has a single route (`GET /proxy`) that accepts a `?url=<target>` query parameter, fetches data from the given external URL, and forwards it to the client.

## Prerequisites

Ensure you have the following installed:

- Golang
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

Pass the target URL as a `url` query parameter:

```sh
curl "http://localhost:3000/proxy?url=https://dummyjson.com/products/1"
```

The server fetches the data from the external service, in this case `dummyjson.com`, and forwards the response to the client.

### Error Handling

- Returns 500 Internal Server Error if anything goes wrong during the fetch.
- Returns the same status code as the external service if it is not 200 OK.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [DummyJSON](https://dummyjson.com)
