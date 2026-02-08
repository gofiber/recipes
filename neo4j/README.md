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

3. Set up your Neo4j database and export connection settings:

    ```sh
    export NEO4J_URI=neo4j://localhost:7687
    export NEO4J_USER=neo4j
    export NEO4J_PASSWORD=password
    ```

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
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	uri := os.Getenv("NEO4J_URI")
	user := os.Getenv("NEO4J_USER")
	pass := os.Getenv("NEO4J_PASSWORD")

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, pass, ""))
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close(context.Background())

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		ctx := context.Background()
		session := driver.NewSession(ctx, neo4j.SessionConfig{
			DatabaseName: "movies",
			AccessMode:   neo4j.AccessModeRead,
		})
		defer session.Close(ctx)

		result, err := session.Run(ctx, "RETURN 'Hello, World!' AS message", nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if result.Next(ctx) {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": result.Record().AsMap()["message"]})
		}
		if result.Err() != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Err().Error()})
		}

		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Neo4j Documentation](https://neo4j.com/docs/)
- [Neo4j Go Driver Documentation](https://pkg.go.dev/github.com/neo4j/neo4j-go-driver/v5/neo4j)
