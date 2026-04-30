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
    export NEO4J_DATABASE=movies
    ```

    Or start Neo4j via Docker Compose:

    ```sh
    docker compose up -d
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
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Movie struct {
	Title    string `json:"title"`
	Tagline  string `json:"tagline"`
	Released int64  `json:"released"`
	Director string `json:"director"`
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	uri := envOrDefault("NEO4J_URI", "neo4j://localhost:7687")
	user := envOrDefault("NEO4J_USER", "neo4j")
	password := envOrDefault("NEO4J_PASSWORD", "password")
	database := envOrDefault("NEO4J_DATABASE", "movies")

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if closeErr := driver.Close(context.Background()); closeErr != nil {
			log.Printf("failed to close neo4j driver: %v", closeErr)
		}
	}()

	if err := driver.VerifyConnectivity(context.Background()); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/movie", func(c fiber.Ctx) error {
		movie := new(Movie)
		if err := c.Bind().Body(movie); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		ctx := c.Context()
		session := driver.NewSession(ctx, neo4j.SessionConfig{
			DatabaseName: database,
			AccessMode:   neo4j.AccessModeWrite,
		})
		defer func() { _ = session.Close(ctx) }()

		query := `CREATE (n:Movie {title: $title, tagline: $tagline, released: $released, director: $director})`
		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			_, runErr := tx.Run(ctx, query, map[string]any{
				"title": movie.Title, "tagline": movie.Tagline,
				"released": movie.Released, "director": movie.Director,
			})
			return nil, runErr
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(movie)
	})

	app.Get("/movie/:title", func(c fiber.Ctx) error {
		title := c.Params("title")
		ctx := c.Context()
		session := driver.NewSession(ctx, neo4j.SessionConfig{
			DatabaseName: database,
			AccessMode:   neo4j.AccessModeRead,
		})
		defer func() { _ = session.Close(ctx) }()

		query := `MATCH (n:Movie {title: $title})
		          RETURN n.title AS title, n.tagline AS tagline, n.released AS released, n.director AS director`
		result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			res, runErr := tx.Run(ctx, query, map[string]any{"title": title})
			if runErr != nil {
				return nil, runErr
			}
			if !res.Next(ctx) {
				if res.Err() != nil {
					return nil, res.Err()
				}
				return nil, nil
			}
			row := res.Record().AsMap()
			return Movie{
				Title:    row["title"].(string),
				Tagline:  row["tagline"].(string),
				Released: row["released"].(int64),
				Director: row["director"].(string),
			}, nil
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if result == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	log.Fatal(app.Listen(":3000"))
}
```

### curl Examples

**Create a movie:**

```sh
curl -X POST http://localhost:3000/movie \
  -H "Content-Type: application/json" \
  -d '{"title":"The Matrix","tagline":"Welcome to the Real World","released":1999,"director":"Lana Wachowski"}'
```

**Get a movie by title:**

```sh
curl http://localhost:3000/movie/The%20Matrix
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Neo4j Documentation](https://neo4j.com/docs/)
- [Neo4j Go Driver Documentation](https://pkg.go.dev/github.com/neo4j/neo4j-go-driver/v5/neo4j)
