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

## Example Usage

1. Open your browser and navigate to `http://localhost:3000`.
2. The client will automatically connect to the SSE endpoint and start receiving updates from the server.

## Code Overview

### `main.go`

The main Go file sets up the Fiber application and handles the SSE connections. It includes the necessary configuration to send events to the client.

### `index.html`

The HTML file provides a simple user interface to connect to the SSE endpoint and display the received messages.

## Additional Information

Server-Sent Events (SSE) is a standard allowing servers to push data to web clients over HTTP. Unlike WebSockets, which require a full-duplex connection, SSE uses a unidirectional connection from the server to the client. This makes SSE simpler to implement and more efficient for scenarios where only the server needs to send updates.

For more information on SSE, you can refer to the following resources:
- [Server-Sent Events on MDN](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
- [Server-Sent Events on Wikipedia](https://en.wikipedia.org/wiki/Server-sent_events)
