---
title: Stream Request Body
keywords: [stream, request body]
description: Streaming request bodies.
---

# Stream Request Body

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/stream-request-body) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/stream-request-body)

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
	"io"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	// Enable request body streaming.
	app.Server().StreamRequestBody = true

	// You can test the route by using cURL:
	// curl -X POST --data-binary @/path/to/large/file localhost:3000
	app.Post("/", func(c fiber.Ctx) error {
		reader := c.RequestCtx().RequestBodyStream()
		if reader == nil {
			return nil
		}
		// Read 1MiB at a time
		buffer := make([]byte, 0, 1024*1024)
		for {
			length, err := io.ReadFull(reader, buffer[:cap(buffer)])
			// Cap the buffer based on the actual length read
			buffer = buffer[:length]
			if length > 0 {
				// Process the chunk - e.g., write to file, parse data, etc.
				log.Printf("Read %d bytes", length)
			}
			if err != nil {
				// EOF or ErrUnexpectedEOF means all data has been read.
				// ErrUnexpectedEOF means the last chunk was smaller than the
				// buffer, which is normal for the final (or only) chunk.
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					break
				}
				return err
			}
		}
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Go io Package Documentation](https://pkg.go.dev/io)
