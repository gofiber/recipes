---
title: MongoDB
keywords: [mongodb, database]
description: Connecting to a MongoDB database.
---

# MongoDB Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/mongodb) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/mongodb)

This project demonstrates how to connect to a MongoDB database in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- MongoDB
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/mongodb
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Set up your MongoDB database and update the connection string in the code.

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

## Example

Here is an example of how to connect to a MongoDB database in a Fiber application:

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // MongoDB connection
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    // Fiber instance
    app := fiber.New()

    // Routes
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, MongoDB!")
    })

    // Start server
    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [MongoDB Documentation](https://docs.mongodb.com)
- [MongoDB Go Driver Documentation](https://pkg.go.dev/go.mongodb.org/mongo-driver)
