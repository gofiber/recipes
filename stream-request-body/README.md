---
title: Stream Request Body
keywords: [stream, request body]
---

# Stream Request Body

This project demonstrates how to handle streaming request bodies in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/stream-request-body
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

Here is an example of how to handle a streaming request body in Go using Fiber:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "io"
    "log"
)

func main() {
    app := fiber.New()

    app.Post("/upload", func(c *fiber.Ctx) error {
        // Open a file to write the streamed data
        file, err := os.Create("uploaded_file")
        if err != nil {
            return err
        }
        defer file.Close()

        // Stream the request body to the file
        _, err = io.Copy(file, c.BodyStream())
        if err != nil {
            return err
        }

        return c.SendString("File uploaded successfully")
    })

    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Go io Package Documentation](https://pkg.go.dev/io)
