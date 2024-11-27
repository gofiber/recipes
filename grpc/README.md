---
title: gRPC
keywords: [grpc, server, client]
description: Using Fiber as a client to a gRPC server.
---

# Example for fiber as a client to gRPC server.

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/grpc) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/grpc)

A sample program to showcase fiber as a client to a gRPC server.

## Prerequisites

- Go 1.16 or higher
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/grpc
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Running the Application

1. Run the gRPC server:
    ```sh
    go run server/main.go
    ```

2. Run the Fiber client:
    ```sh
    go run client/main.go
    ```

3. The server will start on `http://localhost:3000`.

## Endpoints

| Method | URL           | Return value |
| ------ | ------------- | ------------ |
| GET    | /add/:a/:b    | a + b        |
| GET    | /mult/:a/:b   | a \* b       |

### Output

```bash
-> curl http://localhost:3000/add/33445/443234
{"result":"476679"}
-> curl http://localhost:3000/mult/33445/443234
{"result":"14823961130"}
```

## Additional Information

gRPC (gRPC Remote Procedure Calls) is a high-performance, open-source universal RPC framework initially developed by Google. It uses HTTP/2 for transport, Protocol Buffers as the interface description language, and provides features such as authentication, load balancing, and more.

For more information, visit the [official gRPC documentation](https://grpc.io/docs/).
