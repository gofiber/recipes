---
title: WebSocket Chat
keywords: [websocket, chat, chatroom, contrib]
description: Real-time chat application using WebSockets.
---

# WebSocket Chat Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/websocket-chat) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/websocket-chat)

This example demonstrates a simple chat application using Go Fiber and WebSockets.

## Description

This project provides a basic setup for a WebSocket-based chat application using Go Fiber. It includes the necessary configuration and code to run a real-time chat server.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Project Structure

- `main.go`: The main application entry point.
- `home.html`: The HTML file for the chat client.
- `go.mod`: The Go module file.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/websocket-chat
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

The application should now be running on `http://localhost:8080`.

## WebSocket Endpoints

- **GET /ws**: WebSocket endpoint for the chat application.

## Example Usage

1. Open your browser and navigate to `http://localhost:8080`.
2. Enter a message in the input field and click "Send".
3. The message should appear in the chat log.

## Code Overview

### `main.go`

The main Go file sets up the Fiber application, handles WebSocket connections, and manages the chat hub.

### `home.html`

The HTML file provides a simple user interface for the chat application, including a message log and input field.

## Conclusion

This example provides a basic setup for a WebSocket-based chat application using Go Fiber. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [WebSocket Documentation](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket)
