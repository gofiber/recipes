---
title: RSS Feed
keywords: [rss, feed]
description: Generating an RSS feed.
---

# RSS Feed

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/rss-feed) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/rss-feed)

This project demonstrates how to create an RSS feed in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/rss-feed
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

This recipe uses [Mustache templates](https://github.com/cbroglie/mustache) (via the `gofiber/template/mustache` engine) to render an RSS XML response. The template lives in `./xmls/example.xml`.

**`xmls/example.xml`:**

```xml
<?xml version="1.0" encoding="UTF-8"?>
<note>
    <language>{{{Lang}}}</language>
    <title>{{{Title}}}</title>
    <greeting>{{{Greetings}}}</greeting>
</note>
```

**`main.go`:**

```go
package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/template/mustache/v3"
)

func main() {
    engineXML := mustache.New("./xmls", ".xml")
    if err := engineXML.Load(); err != nil {
        log.Fatal(err)
    }

    app := fiber.New()

    app.Get("/rss", func(c fiber.Ctx) error {
        // Set Content-Type to application/rss+xml
        c.Type("rss")

        // Render Mustache template with data
        return engineXML.Render(c, "example", fiber.Map{
            "Lang":      "en",
            "Title":     "hello-rss",
            "Greetings": "Hello World",
        })
    })

    log.Fatal(app.Listen(":3000"))
}
```

### Testing with curl

```sh
curl http://localhost:3000/rss
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Fiber Template - Mustache](https://github.com/gofiber/template/tree/master/mustache)
- [cbroglie/mustache](https://pkg.go.dev/github.com/cbroglie/mustache)
