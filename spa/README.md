---
title: Single Page Application (SPA)
keywords: [spa, react, tailwindcss, parcel]
description: Setting up a Single Page Application (SPA) using React for the frontend and Go for the backend.
---

# Single Page Application (SPA)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/spa) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/spa)

This project demonstrates how to set up a Single Page Application (SPA) using React for the frontend and Go with the Fiber framework for the backend.

## Prerequisites

Ensure you have the following installed:

- Golang
- Node.js
- npm

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/spa
    ```

2. Install frontend dependencies:
    ```sh
    cd frontend
    npm install
    ```

3. Install backend dependencies:
    ```sh
    cd ../backend
    go get
    ```

## Usage

### Building Frontend Assets

1. Build the frontend assets:
    ```sh
    cd frontend
    npm run build
    ```

2. Watch frontend assets for changes:
    ```sh
    npm run dev
    ```

### Running the Application

1. Start the Fiber backend application:
    ```sh
    cd backend
    go run main.go
    ```

## Example

Here is an example of how to serve a SPA with Fiber v3 using the `static` middleware:

```go
package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
    app := fiber.New()

    // Serve Single Page Application on "/web".
    // The NotFoundHandler falls back to index.html so client-side routing works.
    app.Get("/web*", static.New("dist", static.Config{
        NotFoundHandler: func(c fiber.Ctx) error {
            return c.SendFile("./dist/index.html")
        },
    }))

    // Start server
    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [React Documentation](https://reactjs.org/docs/getting-started.html)
- [Tailwind CSS Documentation](https://tailwindcss.com/docs)
- [Parcel Documentation](https://parceljs.org/docs)
