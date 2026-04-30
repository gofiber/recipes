---
title: HTTPS with PKCS12 TLS
keywords: [https, tls, pkcs12]
description: Setting up an HTTPS server with PKCS12 TLS certificates.
---

# HTTPS with PKCS12 TLS Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/https-pkcs12-tls) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/https-pkcs12-tls)

This project demonstrates how to set up an HTTPS server with PKCS12 TLS in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- PKCS12 certificate file (`cert.p12`)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/https-pkcs12-tls
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Place your PKCS12 certificate file (`server.p12`) in the `security/` directory.

4. Optionally set the PKCS12 password via environment variable (defaults to `changeit`):
    ```sh
    export PKCS12_PASSWORD=yourpassword
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the application at `https://localhost:443`.

## Example

Here is an example of how to set up an HTTPS server with PKCS12 TLS in a Fiber application:

```go
package main

import (
    "crypto"
    "crypto/tls"
    "log"
    "os"

    "github.com/gofiber/fiber/v3"
    "golang.org/x/crypto/pkcs12"
)

func main() {
    path := "./security/server.p12"
    password := os.Getenv("PKCS12_PASSWORD")
    if password == "" {
        password = "changeit"
    }

    // Read and decode PKCS12 file
    pkcs12Data, err := os.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    key, cert, err := pkcs12.Decode(pkcs12Data, password)
    if err != nil {
        log.Fatal(err)
    }

    tlsCert := tls.Certificate{
        Certificate: [][]byte{cert.Raw},
        PrivateKey:  key.(crypto.PrivateKey),
        Leaf:        cert,
    }

    config := &tls.Config{Certificates: []tls.Certificate{tlsCert}}

    app := fiber.New()
    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("This page is being served over TLS using a PKCS12 store type!")
    })

    ln, err := tls.Listen("tcp", ":443", config)
    if err != nil {
        log.Fatal(err)
    }

    log.Fatal(app.Listener(ln))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [TLS in Go](https://golang.org/pkg/crypto/tls/)
- [PKCS12 in Go](https://pkg.go.dev/golang.org/x/crypto/pkcs12)
