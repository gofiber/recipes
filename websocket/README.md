# WebSocket Example

This example demonstrates a simple WebSocket application using Go Fiber.

## Description

This project provides a basic setup for a WebSocket server using Go Fiber. It includes the necessary configuration and code to run a real-time WebSocket server.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Project Structure

- `main.go`: The main application entry point.
- `go.mod`: The Go module file.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/websocket
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

The application should now be running on `http://localhost:3000`.

## WebSocket Endpoint

- **GET /ws**: WebSocket endpoint for the application.

## Example Usage

1. Connect to the WebSocket server at `ws://localhost:3000/ws`.
2. Send a message to the server.
3. The server will echo the message back to the client.

## Code Overview

### `main.go`

The main Go file sets up the Fiber application, handles WebSocket connections, and manages the WebSocket communication.

```go
package main

import (
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/contrib/websocket"
)

func main() {
    app := fiber.New()

    // Optional middleware
    app.Use("/ws", func(c *fiber.Ctx) error {
        if c.Get("host") == "localhost:3000" {
            c.Locals("Host", "Localhost:3000")
            return c.Next()
        }
        return c.Status(403).SendString("Request origin not allowed")
    })

    // Upgraded websocket request
    app.Get("/ws", websocket.New(func(c *websocket.Conn) {
        fmt.Println(c.Locals("Host")) // "Localhost:3000"
        for {
            mt, msg, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                break
            }
            log.Printf("recv: %s", msg)
            err = c.WriteMessage(mt, msg)
            if err != nil {
                log.Println("write:", err)
                break
            }
        }
    }))

    // ws://localhost:3000/ws
    log.Fatal(app.Listen(":3000"))
}
```

## Conclusion

This example provides a basic setup for a WebSocket server using Go Fiber. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [WebSocket Documentation](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket)
