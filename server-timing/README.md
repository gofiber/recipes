---
title: Server Timing
keywords: [server timing]
description: Adding Server Timing headers to an application.
---

# Server Timing

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/server-timing) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/server-timing)

This project demonstrates how to implement Server-Timing headers in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/server-timing
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

## Example

Here is an example of how to set up Server-Timing headers in a Fiber application:

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    app.Use(func(c fiber.Ctx) error {
        start := time.Now()
        err := c.Next()
        // dur value must be in milliseconds per W3C spec
        c.Append("Server-Timing", fmt.Sprintf("app;dur=%.2f", float64(time.Since(start).Microseconds())/1000.0))
        return err
    })

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))
}
```

### Testing with curl

```sh
curl -i http://localhost:3000/
```

Example response header:

```
Server-Timing: app;dur=2001.23
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Server-Timing Header Documentation](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server-Timing)
