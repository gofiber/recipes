---
title: Template
keywords: [template, tailwindcss, parcel]
description: Setting up a Go application with template rendering.
---

# Template Project

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/template) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/template)

This project demonstrates how to set up a Go application with template rendering, Tailwind CSS, and Parcel for asset bundling.

## Prerequisites

Ensure you have the following installed:

- Golang
- Node.js
- npm

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/template
    ```

2. Install dependencies:
    ```sh
    npm install
    ```

## Usage

### Building Assets

1. Build the assets:
    ```sh
    npm run build
    ```

2. Watch assets for changes:
    ```sh
    npm run dev
    ```

### Running the Application

1. Start the Fiber application:
    ```sh
    go run main.go
    ```

## Example

Here is an example of how to set up a basic route with template rendering in Go:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func main() {
    // Initialize the template engine
    engine := html.New("./views", ".html")

    // Create a new Fiber instance with the template engine
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Define a route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
            "Title": "Hello, World!",
        })
    })

    // Start the server
    app.Listen(":3000")
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Tailwind CSS Documentation](https://tailwindcss.com/docs)
- [Parcel Documentation](https://parceljs.org/docs)
