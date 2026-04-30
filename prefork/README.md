---
title: Prefork
keywords: [prefork]
description: Running an application in prefork mode.
---

# Prefork Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/prefork) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/prefork)

This project demonstrates how to use the `Prefork` feature in a Go application using the Fiber framework. Preforking can improve performance by utilizing multiple CPU cores.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/prefork
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

Here is an example of how to set up the `Prefork` feature in a Fiber application:

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gofiber/fiber/v3"
)

func main() {
    // Print current process
    if fiber.IsChild() {
        fmt.Printf("[%d] Child\n", os.Getppid())
    } else {
        fmt.Printf("[%d] Master\n", os.Getppid())
    }

    // Fiber instance
    app := fiber.New()

    // Routes
    app.Get("/", hello)

    // Start server with prefork enabled
    log.Fatal(app.Listen(":3000", fiber.ListenConfig{EnablePrefork: true}))
}

// Handler
func hello(c fiber.Ctx) error {
    return c.SendString("Hello, World!")
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Fiber Prefork Documentation](https://docs.gofiber.io/api/fiber#prefork)
