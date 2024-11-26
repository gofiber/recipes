---
title: Socketio
keywords: [websocket, chat, socketio, chatroom, contrib]
---

# WebSocket Chat Example

This example demonstrates how to create a simple chatroom using WebSockets. The chatroom supports multiple users and allows them to send messages to each other.

## Prerequisites

- Go 1.16 or higher
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/socketio-chat
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

## Connecting to the WebSocket

To connect to the WebSocket, use the following URL:
```
ws://localhost:3000/ws/<user-id>
```

## Message Object Example

Here is an example of a message object that can be sent between users:
```json
{
  "from": "<user-id>",
  "to": "<recipient-user-id>",
  "data": "hello"
}
```
