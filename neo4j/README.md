---
title: Neo4j
keywords: [neo4j, database]
description: Connecting to a Neo4j database.
---

# Neo4j Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/neo4j) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/neo4j)

This project demonstrates how to connect to a Neo4j database in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- Neo4j
- [Neo4j Go Driver](https://github.com/neo4j/neo4j-go-driver)

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/gofiber/recipes.git
   cd recipes/neo4j
   ```

2. Install dependencies:

   ```sh
   go get
   ```

3. Set up your Neo4j database and update the connection string in the code.

## Running the Application

1. Start the application:

    ```sh
    go run main.go
    ```

## Example

Here is an example of how to connect to a Neo4j database in a Fiber application:

```go
package main

import (
    "log"
    "net/http"  
    "github.com/gofiber/fiber/v2"
    "github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
    driver, err := neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
    if err != nil {
        log.Fatal(err)
    }
    defer driver.Close()    
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        session := driver.NewSession(neo4j.SessionConfig{})
        defer session.Close()   
        result, err := session.Run("RETURN 'Hello, World!'", nil)
        if err != nil {
            return c.Status(http.StatusInternalServerError).SendString(err.Error())
        }   
        if result.Next() {
            return c.SendString(result.Record().Values[0].(string))
        }

        return c.SendStatus(http.StatusInternalServerError)
    })

    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Neo4j Documentation](https://neo4j.com/docs/)
- [Neo4j Go Driver Documentation](https://pkg.go.dev/github.com/neo4j/neo4j-go-driver)
