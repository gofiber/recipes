---
title: Air Live Reloading
keywords: [air, live reloading, development, air tool, hot reload, watch, changes]
description: Live reloading for Go applications.
---

# Live Reloading with Air Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/air) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/air)

This example demonstrates how to set up live reloading for a Go application using the [Air](https://github.com/air-verse/air) tool. The purpose of this example is to show how to automatically reload your application during development whenever you make changes to the source code.

## Description

Live reloading is a useful feature during development as it saves time by automatically restarting the application whenever changes are detected. This example sets up a simple Fiber application and configures Air to watch for changes and reload the application.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)
- [Air](https://github.com/air-verse/air)

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/air
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

3. Install Air:
    ```bash
    go install github.com/air-verse/air@latest
    ```

## Configuration

Air is configured using the `air/.air.conf` file. This file specifies the build command, binary name, and directories to watch for changes. The configuration files for different operating systems are provided:

- `air/.air.windows.conf` for Windows
- `air/.air.linux.conf` for Linux

## Running the Example

To run the example with live reloading, use the following command:
```bash
air -c .air.linux.conf
```
or for Windows:
```bash
air -c .air.windows.conf
```

The server will start and listen on `localhost:3000`. Any changes to the source code will automatically trigger a rebuild and restart of the application.

## Example Routes

- **GET /**: Returns a simple greeting message.

## Code Overview

### `main.go`

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Create new Fiber instance
    app := fiber.New()

    // Create new GET route on path "/"
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Start server on http://localhost:3000
    log.Fatal(app.Listen(":3000"))
}
```

## Conclusion

This example provides a basic setup for live reloading a Go application using Air. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Air Documentation](https://github.com/air-verse/air)
- [Fiber Documentation](https://docs.gofiber.io)
- [GitHub Repository](https://github.com/gofiber/fiber)
