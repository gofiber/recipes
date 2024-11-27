---
title: PostgreSQL
keywords: [postgresql]
description: Connecting to a PostgreSQL database.
---

# PostgreSQL Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/postgresql) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/postgresql)

This project demonstrates how to connect to a PostgreSQL database in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- PostgreSQL

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/postgresql
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Set up your PostgreSQL database and update the connection string in the code.

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the application at `http://localhost:3000`.

## Example

Here is an example of how to connect to a PostgreSQL database in a Fiber application:

```go
package main

import (
    "database/sql"
    "log"

    "github.com/gofiber/fiber/v2"
    _ "github.com/lib/pq"
)

func main() {
    // Database connection
    connStr := "user=username dbname=mydb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Fiber instance
    app := fiber.New()

    // Routes
    app.Get("/", func(c *fiber.Ctx) error {
        var greeting string
        err := db.QueryRow("SELECT 'Hello, World!'").Scan(&greeting)
        if err != nil {
            return err
        }
        return c.SendString(greeting)
    })

    // Start server
    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [pq Driver Documentation](https://pkg.go.dev/github.com/lib/pq)
