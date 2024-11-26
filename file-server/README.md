---
title: File Server
keywords: [file server, static files]
---

# File Server Example

This project demonstrates how to set up a simple file server in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/file-server
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

2. Access the file server at `http://localhost:3000`.

## Example

Here is an example `main.go` file for the Fiber application serving static files:

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Serve static files from the "public" directory
    app.Static("/", "./public")

    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Golang Documentation](https://golang.org/doc/)
