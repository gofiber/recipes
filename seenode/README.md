---
title: Seenode
keywords: [seenode, deploy, cloud]
description: Deploying to Seenode cloud platform.
---

# Seenode Deployment Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/seenode) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/seenode)

This project demonstrates how to deploy a Go application using the Fiber framework on Seenode.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- [Seenode account](https://cloud.seenode.com)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/seenode
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Create a Seenode account and connect your repository:
    - Go to [Seenode Dashboard](https://cloud.seenode.com)
    - Create a new Web Service
    - Connect your Git repository

4. Configure deployment:
    - **Build Command**: `go build -o app main.go`
    - **Start Command**: `./app`

5. Deploy the application:
    ```sh
    git add .
    git commit -m "Deploy to Seenode"
    git push
    ```

## Running the Application

1. Open the application in your browser using the provided Seenode URL.

## Example

Here is an example `main.go` file for the Fiber application:

```go
package main

import (
    "fmt"
    "os"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Welcome to seenode 👋")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }
    
    app.Listen(fmt.Sprintf(":%s", port))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Seenode Documentation](https://seenode.com/docs/frameworks/go/fiber/)
- [Seenode Dashboard](https://cloud.seenode.com)
