---
title: Socketio
keywords: [websocket, chat, socketio, chatroom, contrib]
description: A chatroom application using Socket.IO.
---

# Socket.IO Chat Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/socketio) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/socketio)

This example demonstrates how to create a simple chatroom with Socket.IO. The chatroom supports multiple users and allows them to send messages to each other.

The server uses the [`socketio`](https://github.com/gofiber/contrib/tree/main/v3/socketio) package from Fiber contrib, which implements the Engine.IO v4 / Socket.IO v5 protocol. Clients therefore connect with the official [`socket.io-client`](https://socket.io/docs/v4/client-api/) library (or any other Socket.IO client). A plain WebSocket client does not work: the server expects the Socket.IO handshake first and answers a message sent before it with `Expected SIO CONNECT`. For plain WebSockets, see the [WebSocket](../websocket/README.md) recipe.

## Prerequisites

- Go 1.27 or higher
- Go modules
- Node.js, to run the example client below

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/socketio
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

## Connecting a Client

The user id is part of the route (`/ws/:id`), so it goes into the `path` option of the client. Set `transports` to `["websocket"]`: this recipe does not enable HTTP long-polling, and the client tries polling first by default, which fails with `xhr poll error`.

```js
import { io } from "socket.io-client";

const socket = io("http://localhost:3000", {
  path: "/ws/<user-id>",
  transports: ["websocket"],
});

socket.on("message", (msg) => console.log(msg));
```

Install the client library with `npm install socket.io-client`.

## Message Object Example

Chat messages are Socket.IO events named `message`. To send a message to another user, emit a message object whose `to` field is the user id of the recipient:
```json
{
  "from": "<user-id>",
  "to": "<recipient-user-id>",
  "data": "hello"
}
```

The server delivers the object unchanged to the `message` event of the recipient. A message for a user who is not connected is dropped, and the server only logs `message cannot be delivered invalid/gone connection`.

The server also sends every new connection `Hello user: <user-id> with UUID: <uuid>` and tells the other connected users `New user connected: <user-id> and UUID: <uuid>`. Both arrive as `message` events as well.

## Example

A complete example with two users. Save it as `chat.mjs` and run `node chat.mjs` while the server is running (stop it with Ctrl+C):

```js
import { io } from "socket.io-client";

const connect = (id) =>
  io("http://localhost:3000", { path: `/ws/${id}`, transports: ["websocket"] });

const alice = connect("alice");
const bob = connect("bob");

alice.on("message", (msg) => console.log("alice got:", msg));
bob.on("message", (msg) => console.log("bob got:", msg));

// The server greets every new connection, so the first message bob receives
// also means the server has registered him.
bob.once("message", () => {
  alice.emit("message", { from: "alice", to: "bob", data: "hello" });
});
```

It prints the following. The order of the first three lines can vary between runs:

```
alice got: Hello user: alice with UUID: <uuid>
bob got: Hello user: bob with UUID: <uuid>
alice got: New user connected: bob and UUID: <uuid>
bob got: { from: 'alice', to: 'bob', data: 'hello' }
```

## References

- [Socket.IO package for Fiber](https://github.com/gofiber/contrib/tree/main/v3/socketio): polling, acknowledgements, namespaces and handshake auth
- [Socket.IO client API](https://socket.io/docs/v4/client-api/)
