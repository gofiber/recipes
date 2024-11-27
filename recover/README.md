---
title: Recover Middleware
keywords: [recover, middleware]
description: Recover middleware for error handling.
---

# Recover Middleware Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/recover) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/recover)

This project demonstrates how to implement a recovery mechanism in a Go application using the Fiber framework's `Recover` middleware.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/recover
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

Here is an example of how to set up the `Recover` middleware in a Fiber application:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
    app := fiber.New()

    // Use the Recover middleware
    app.Use(recover.New())

    app.Get("/", func(c *fiber.Ctx) error {
        // This will cause a panic
        panic("something went wrong")
    })

    app.Listen(":3000")
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Fiber Recover Middleware Documentation](https://docs.gofiber.io/api/middleware/recover)
