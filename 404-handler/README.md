---
title: 404 Handler
keywords: [404, not found, handler, errorhandler, custom]
---

# Custom 404 Not Found Handler Example

This example demonstrates how to implement a custom 404 Not Found handler using the [Fiber](https://gofiber.io) web framework in Go. The purpose of this example is to show how to handle requests to undefined routes gracefully by returning a 404 status code.

## Description

In web applications, it's common to encounter requests to routes that do not exist. Handling these requests properly is important to provide a good user experience and to inform the user that the requested resource is not available. This example sets up a simple Fiber application with a custom 404 handler to manage such cases.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Running the Example

To run the example, use the following command:
```bash
go run main.go
```

The server will start and listen on `localhost:3000`.

## Example Routes

- **GET /hello**: Returns a simple greeting message.
- **Undefined Routes**: Any request to a route not defined will trigger the custom 404 handler.

## Custom 404 Handler

The custom 404 handler is defined to catch all undefined routes and return a 404 status code with a "Not Found" message.

## Code Overview

### `main.go`

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Fiber instance
    app := fiber.New()

    // Routes
    app.Get("/hello", hello)

    // 404 Handler
    app.Use(func(c *fiber.Ctx) error {
        return c.SendStatus(404) // => 404 "Not Found"
    })

    // Start server
    log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
    return c.SendString("I made a ☕ for you!")
}
```

## Conclusion

This example provides a basic setup for handling 404 Not Found errors in a Fiber application. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [GitHub Repository](https://github.com/gofiber/fiber)
