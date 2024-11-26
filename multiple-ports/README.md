---
title: Multiple Ports
keywords: [multiple ports, server, port]
---

# Multiple Ports Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/multiple-ports) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/multiple-ports)

This project demonstrates how to run a Go application using the Fiber framework on multiple ports.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/multiple-ports
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

Here is an example of how to run a Fiber application on multiple ports:

```go
package main

import (
    "log"
    "sync"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    ports := []string{":3000", ":3001"}

    var wg sync.WaitGroup
    for _, port := range ports {
        wg.Add(1)
        go func(p string) {
            defer wg.Done()
            if err := app.Listen(p); err != nil {
                log.Printf("Error starting server on port %s: %v", p, err)
            }
        }(port)
    }

    wg.Wait()
}
```

In this example:
- The application listens on multiple ports (`:3000` and `:3001`).
- A `sync.WaitGroup` is used to wait for all goroutines to finish.

## References

- [Fiber Documentation](https://docs.gofiber.io)
