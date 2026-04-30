---
title: Optional Parameter
keywords: [optional, parameter]
description: Handling optional parameters.
---

# Optional Parameter Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/optional-parameter) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/optional-parameter)

This project demonstrates how to handle optional parameters in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/optional-parameter
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

Here is an example of how to handle optional parameters in a Fiber application:

```go
package main

import (
    "log"
    "strconv"

    "github.com/gofiber/fiber/v3"
)

func main() {
    // user list
    users := [...]string{"Alice", "Bob", "Charlie", "David"}

    // Fiber instance
    app := fiber.New()

    // Route to profile
    app.Get("/:id?", func(c fiber.Ctx) error {
        id, err := strconv.Atoi(c.Params("id")) // transform id to array index
        if err != nil || id < 0 || id >= len(users) {
            return c.SendStatus(fiber.StatusNotFound) // invalid parameter returns 404
        }
        return c.SendString("Hello, " + users[id] + "!") // custom hello message to user with the id
    })

    // Start server
    log.Fatal(app.Listen(":3000"))
}
```

In this example:
- The `:id?` parameter in the route is optional.
- If no valid `id` is provided, a `404 Not Found` is returned.
- Valid `id` values (0-3) map to users Alice, Bob, Charlie, and David.

## References

- [Fiber Documentation](https://docs.gofiber.io)
