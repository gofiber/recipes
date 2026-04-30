---
title: GraphQL
keywords: [graphql]
description: Setting up a GraphQL server.
---

# GraphQL Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/graphql) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/graphql)

This project demonstrates how to set up a GraphQL server in a Go application using the Fiber framework and the [graphql-go](https://github.com/graphql-go/graphql) library.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- [graphql-go](https://github.com/graphql-go/graphql) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/graphql
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. The server listens on port `9090` and exposes a single endpoint at `/` that accepts both GET and POST requests.

## Usage

### GET request

Pass the GraphQL query as a URL-encoded `query` parameter:

```sh
curl 'http://localhost:9090/?query=query%7Bhello%7D'
```

### POST request

Send the query as a JSON body with `Content-Type: application/json`:

```sh
curl 'http://localhost:9090/' \
  --header 'content-type: application/json' \
  --data-raw '{"query":"query{hello}"}'
```

Both return a JSON response like:

```json
{"data":{"hello":"world"}}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [graphql-go Documentation](https://github.com/graphql-go/graphql)
- [GraphQL Documentation](https://graphql.org/)
