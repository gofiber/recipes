---
title: Server-Sent Events
keywords: [sse, server-sent events, real-time]
description: Implementing Server-Sent Events in an application.
---

# Server-Sent Events with Fiber

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/sse) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/sse)

This example demonstrates how to implement Server-Sent Events (SSE) in a Fiber application.

## Description

Server-Sent Events (SSE) allow servers to push updates to the client over a single HTTP connection. This is useful for real-time applications where the server needs to continuously send data to the client, such as live feeds, notifications, or real-time charts.

## Prerequisites

- Go 1.16 or higher
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/sse
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

## Endpoints

- **GET /**: Index page
- **GET /sse**: SSE route
- **PUT /publish**: Send messages via SSE

## Example Usage

By default the example will run on port `3000`, this can be changed by modifying the `appPort` constant in `main.go`

1. Open your browser and navigate to `http://localhost:3000`.
2. The client will automatically connect to the SSE endpoint and start receiving updates from the server.
3. The `/sse` endpoint will publish the current time to the client every two seconds

### Custom Messages

To send a custom message, send a `PUT` request to the `/publish` endpoint in the following JSON format

```json
{
  "message": "Hello, World!"
}
```

Messages sent to the `/publish` endpoint will be added to a queue that is read from in FIFO order. You can test this
by using curl in an iterator

If you are using bash/zsh:
```sh
for i in {1..10}; do
  curl -X PUT -H 'Content-type: application/json' --data "{\"message\":\"SSE TEST $i\"}" http://localhost:3000/publish
done
```

If you are using fish:
```sh
for i in (seq 1 10)
  curl -X PUT -H 'Content-type: application/json' --data "{\"message\":\"SSE TEST $i\"}" http://localhost:3000/publish
end
```

Once published, your added messages will begin appearing in the output at `http://localhost:3000`. Once the queue is empty
and no user-published messages are left, `/sse` will return to it's standard behavior of displaying the current time.


## Code Overview

### `main.go`

The main Go file sets up the Fiber application and handles the SSE connections. It includes the necessary configuration to send events to the client.

## Additional Information

Server-Sent Events (SSE) is a standard allowing servers to push data to web clients over HTTP. Unlike WebSockets, which require a full-duplex connection, SSE uses a unidirectional connection from the server to the client. This makes SSE simpler to implement and more efficient for scenarios where only the server needs to send updates.

For more information on SSE, you can refer to the following resources:
- [Server-Sent Events on MDN](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
- [Server-Sent Events on Wikipedia](https://en.wikipedia.org/wiki/Server-sent_events)
