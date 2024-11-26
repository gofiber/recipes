---
title: Optional Parameter Example
keywords: [optional, parameter]
---

# Optional Parameter Example

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
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/user/:id?", func(c *fiber.Ctx) error {
        id := c.Params("id", "defaultID")
        return c.SendString("User ID: " + id)
    })

    app.Listen(":3000")
}
```

In this example:
- The `:id?` parameter in the route is optional.
- If the `id` parameter is not provided, it defaults to `"defaultID"`.

## References

- [Fiber Documentation](https://docs.gofiber.io)
