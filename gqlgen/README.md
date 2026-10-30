---
title: GraphQL Subscriptions with gqlgen
keywords: [graphql, gqlgen, subscriptions, websocket, graphql-transport-ws]
description: Schema-first GraphQL with gqlgen, with queries and mutations over HTTP and subscriptions over WebSocket.
---

# GraphQL Subscriptions with gqlgen

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/gqlgen) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/gqlgen)

This example shows a schema-first GraphQL server built with [gqlgen](https://gqlgen.com) on Fiber v3. It serves a query, a mutation and a subscription that fires on every mutation, all on one `/query` path.

The domain is a single message feed. `postMessage` stores a message, `messages` returns the most recent ones and `messageAdded` streams every new message to subscribed clients.

## How It Works on Fiber

Queries and mutations go through the Fiber [adaptor](https://docs.gofiber.io/next/middleware/adaptor) middleware, which runs the gqlgen `net/http` handler on Fiber.

The adaptor cannot carry a websocket upgrade, so subscriptions use the Fiber [websocket](https://github.com/gofiber/contrib/tree/main/v3/websocket) middleware on the same `/query` path, and a small bridge in `graphqlws.go` hands the upgraded connection to gqlgen's own websocket transport.

The protocol is `graphql-transport-ws`, which the graphql-ws library and GraphiQL speak. Upgrades that do not offer it, including legacy `graphql-ws`, get a 400 response.

## Project Structure

- `main.go`: gqlgen server setup, routes and graceful shutdown.
- `graphqlws.go`: bridge from the Fiber websocket connection to the gqlgen websocket transport.
- `complexity.go`: the operation complexity limit.
- `chat/hub.go`: in-memory message history and pub-sub fan-out.
- `graph/schema.graphqls`: the GraphQL schema.
- `graph/resolver.go`: the root resolver and the per-connection subscription limit.
- `graph/schema.resolvers.go`: resolver implementations.

## Requirements

- [Go](https://golang.org/dl/) 1.26 or higher
- [Git](https://git-scm.com/downloads)
- Optional: [websocat](https://github.com/vi/websocat) to test subscriptions from the command line

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/gqlgen
    ```

2. Run the application:

    ```bash
    go run .
    ```

The server listens on `http://localhost:3000`. GraphiQL opens at `http://localhost:3000/`.

## Example Usage

1. Query the recent messages:

    ```bash
    curl -s http://localhost:3000/query -H 'Content-Type: application/json' -d '{"query":"{ messages { id author text createdAt } }"}'
    ```

    ```json
    {"data":{"messages":[]}}
    ```

2. In a second terminal, open a subscription. The command sends `connection_init` and waits a second for `connection_ack`. Then it subscribes, and `cat` keeps the connection open:

    ```bash
    (printf '%s\n' '{"type":"connection_init"}'; sleep 1; printf '%s\n' '{"id":"1","type":"subscribe","payload":{"query":"subscription { messageAdded { id author text createdAt } }"}}'; cat) | websocat --protocol graphql-transport-ws ws://localhost:3000/query
    ```

    ```json
    {"type":"connection_ack"}
    ```

3. Post a message:

    ```bash
    curl -s http://localhost:3000/query -H 'Content-Type: application/json' -d '{"query":"mutation($in: NewMessage!) { postMessage(input: $in) { id author text } }","variables":{"in":{"author":"ana","text":"hello"}}}'
    ```

    ```json
    {"data":{"postMessage":{"id":"1","author":"ana","text":"hello"}}}
    ```

    The websocat terminal prints the event:

    ```json
    {"payload":{"data":{"messageAdded":{"id":"1","author":"ana","text":"hello","createdAt":"2026-10-29T11:31:16.823051Z"}}},"id":"1","type":"next"}
    ```

4. Type `{"id":"1","type":"complete"}` in the websocat terminal to stop the subscription. The server confirms with `{"id":"1","type":"complete"}`, and later posts no longer arrive. The server also sends `{"type":"ping"}` every 10 seconds.

5. Invalid input returns a GraphQL error:

    ```bash
    curl -s http://localhost:3000/query -H 'Content-Type: application/json' -d '{"query":"mutation { postMessage(input: {author: \"ana\", text: \"\"}) { id } }"}'
    ```

    ```json
    {"errors":[{"message":"text must be 1 to 280 characters","path":["postMessage"],"locations":[{"line":1,"column":12}]}],"data":null}
    ```

## Regenerating Code

1. Edit `graph/schema.graphqls`.
2. Run `go generate ./...` or `make generate`. The generator version comes from the `tool` line in `go.mod`.
3. Implement any new resolver stubs in `graph/schema.resolvers.go`.
4. Commit the generated files. `make build`, which CI runs, regenerates the code and fails if that changes any file in `graph/`.

## Design Notes

- The history keeps the last 100 messages. `messages(last: n)` accepts 1 to 100.
- Every subscriber has a buffer of 16 messages. A subscriber that falls further behind is evicted and receives `complete`, so a slow client never blocks `postMessage`.
- The server accepts at most 10000 websocket connections, 10000 open subscriptions in total and 10 subscriptions per connection. Extra connections are closed with code 1013, try again later.
- Every operation is limited to 1000 parser tokens and a complexity of 500, since gqlgen sets no cost limit by default. `complexity.go` computes the cost like `extension.FixedComplexityLimit`, but walks each named fragment only once.
- `MissingPongOk: true` keeps clients that do not answer pings, such as websocat, connected. Set it to `false` in production so dead peers are dropped.
- Websocket operations do not see the HTTP request headers. Authentication for subscriptions belongs in `transport.Websocket.InitFunc`, which reads the `connection_init` payload.
- An empty `websocket.Config.Origins` allows every origin. Set it for browser clients in production. Once it lists specific origins, clients that send no `Origin` header, such as websocat or the Go dialer, are refused unless `AllowEmptyOrigin` is `true`.

## Graceful Shutdown

Ctrl+C or `SIGTERM` closes open subscriptions with code 1000 and gives HTTP requests and websocket connections one shared 5-second deadline. A websocket client that has not sent `connection_init` is closed with code 1002 by the 3-second init timeout, before that deadline.

## References

- [gqlgen documentation](https://gqlgen.com)
- [graphql-transport-ws protocol](https://github.com/enisdenjo/graphql-ws/blob/master/PROTOCOL.md)
- [Fiber websocket middleware](https://github.com/gofiber/contrib/tree/main/v3/websocket)
- [Fiber adaptor middleware](https://docs.gofiber.io/next/middleware/adaptor)
