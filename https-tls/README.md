---
title: HTTPS with TLS
keywords: [https, tls, ssl, self-signed]
description: Setting up an HTTPS server with self-signed TLS certificates.
---

# HTTPS with TLS Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/https-tls) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/https-tls)

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

3. Generate a self-signed certificate and key:
    ```sh
    mkdir -p certs
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
      -keyout certs/ssl.key -out certs/ssl.cert \
      -subj "/CN=localhost"
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the application at `https://localhost:443`.

## Example

Here is an example of how to set up an HTTPS server with TLS in a Fiber application:

```go
package main

import (
    "crypto/tls"
    "log"

    "github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString(c.Protocol()) // => https
    })

    // Create tls certificate
    cer, err := tls.LoadX509KeyPair("certs/ssl.cert", "certs/ssl.key")
    if err != nil {
        log.Fatal(err)
    }

    config := &tls.Config{Certificates: []tls.Certificate{cer}}

    // Create custom listener
    ln, err := tls.Listen("tcp", ":443", config)
    if err != nil {
        log.Fatal(err)
    }

    // Start server with https/ssl enabled on http://localhost:443
    log.Fatal(app.Listener(ln))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [TLS in Go](https://golang.org/pkg/crypto/tls/)
