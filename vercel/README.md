---
title: Vercel
keywords: [vercel, deploy, serverless]
description: Deploy a Go application to Vercel.
---

# Vercel Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/vercel) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/vercel)

This example demonstrates how to deploy a Go Fiber application to Vercel.

## Description

This project provides a starting point for deploying a Go Fiber application to Vercel. It includes the necessary configuration files and code to run a serverless application on Vercel.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)
- [Vercel CLI](https://vercel.com/download)

## Project Structure

- `api/index.go`: The main entry point for the serverless function.
- `vercel.json`: Configuration file for Vercel.
- `go.mod`: The Go module file.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/vercel
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

## Configuration

Ensure the `vercel.json` file is present in the root directory to handle routing properly. This file rewrites all requests to the `api/index.go` handler.

```json
{
  "rewrites": [
    { "source": "(.*)", "destination": "api/index.go" }
  ]
}
```

## Deploy

1. Install the Vercel CLI:
    ```bash
    npm install -g vercel
    ```

2. Log in to Vercel:
    ```bash
    vercel login
    ```

3. Deploy the application:
    ```bash
    vercel
    ```

Follow the prompts to complete the deployment. Your application will be deployed to Vercel and a URL will be provided.

## Example Usage

1. Open your browser and navigate to the provided Vercel URL.
2. You should see the JSON response with the URI and path.

## Code Overview

### `api/index.go`

The main Go file sets up the Fiber application, handles HTTP requests, and manages the routing.

```go
package handler

import (
 "github.com/gofiber/fiber/v2/middleware/adaptor"
 "github.com/gofiber/fiber/v2"
 "net/http"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
 // This is needed to set the proper request path in `*fiber.Ctx`
 r.RequestURI = r.URL.String()

 handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
 app := fiber.New()

 app.Get("/v1", func(ctx *fiber.Ctx) error {
  return ctx.JSON(fiber.Map{
   "version": "v1",
  })
 })

 app.Get("/v2", func(ctx *fiber.Ctx) error {
  return ctx.JSON(fiber.Map{
   "version": "v2",
  })
 })

 app.Get("/", func(ctx *fiber.Ctx) error {
  return ctx.JSON(fiber.Map{
   "uri":  ctx.Request().URI().String(),
   "path": ctx.Path(),
  })
 })

 return adaptor.FiberApp(app)
}
```

## Conclusion

This example provides a basic setup for deploying a Go Fiber application to Vercel. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Vercel Documentation](https://vercel.com/docs)
- [Fiber Documentation](https://docs.gofiber.io)
