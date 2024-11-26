---
title: HTTPS with TLS
keywords: [https, tls, ssl, self-signed]
---

# HTTPS with TLS Example

This project demonstrates how to set up an HTTPS server with TLS in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- TLS certificates (self-signed or from a trusted CA)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/https-tls
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Place your TLS certificate (`cert.pem`) and private key (`key.pem`) in the project directory.

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the application at `https://localhost:3000`.

## Example

Here is an example of how to set up an HTTPS server with TLS in a Fiber application:

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, HTTPS with TLS!")
    })

    // Start server with TLS
    log.Fatal(app.ListenTLS(":3000", "cert.pem", "key.pem"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [TLS in Go](https://golang.org/pkg/crypto/tls/)
