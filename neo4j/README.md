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
    export NEO4J_DATABASE=neo4j
    ```

    `NEO4J_DATABASE` defaults to `neo4j`, the database that every installation has. The Community Edition used by the Docker Compose file supports only that one user database. With the Enterprise Edition or Aura you can create another one (`CREATE DATABASE movies`) and set `NEO4J_DATABASE=movies`.

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

This is the whole `main.go` of the recipe. It connects to Neo4j, stores movies with `POST /movie` and reads them back with `GET /movie/:title`:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Movie represent the movie schema
type Movie struct {
	Title    string `json:"title"`
	Tagline  string `json:"tagline"`
	Released int64  `json:"released"`
	Director string `json:"director"`
}

func movieFromRecord(record *neo4j.Record) (Movie, error) {
	row := record.AsMap()

	title, ok := row["title"].(string)
	if !ok {
		return Movie{}, fmt.Errorf("invalid title type: %T", row["title"])
	}

	tagline, ok := row["tagline"].(string)
	if !ok {
		return Movie{}, fmt.Errorf("invalid tagline type: %T", row["tagline"])
	}

	director, ok := row["director"].(string)
	if !ok {
		return Movie{}, fmt.Errorf("invalid director type: %T", row["director"])
	}

	released, ok := row["released"].(int64)
	if !ok {
		return Movie{}, fmt.Errorf("invalid released type: %T", row["released"])
	}

	return Movie{
		Title:    title,
		Tagline:  tagline,
		Released: released,
		Director: director,
	}, nil
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func main() {
	uri := envOrDefault("NEO4J_URI", "neo4j://localhost:7687")
	user := envOrDefault("NEO4J_USER", "neo4j")
	password := envOrDefault("NEO4J_PASSWORD", "password")
	// "neo4j" is the default database of every installation. The Community Edition (the image in
	// docker-compose.yml) can only have this one user database, so defaulting to another name
	// would make every request fail with "database does not exist".
	database := envOrDefault("NEO4J_DATABASE", "neo4j")

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

	// Create a Fiber app
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
		defer func() {
			_ = session.Close(ctx)
		}()

		query := `CREATE (n:Movie {title: $title, tagline: $tagline, released: $released, director: $director})`
		params := map[string]any{
			"title":    movie.Title,
			"tagline":  movie.Tagline,
			"released": movie.Released,
			"director": movie.Director,
		}

		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			_, runErr := tx.Run(ctx, query, params)
			return nil, runErr
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(movie)
	})

	app.Get("/movie/:title", func(c fiber.Ctx) error {
		// Fiber leaves route parameters percent-encoded, so "The%20Matrix" has to be decoded
		// before it can match the stored title "The Matrix".
		title, err := url.PathUnescape(c.Params("title"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid movie title"})
		}
		ctx := c.Context()
		session := driver.NewSession(ctx, neo4j.SessionConfig{
			DatabaseName: database,
			AccessMode:   neo4j.AccessModeRead,
		})
		defer func() {
			_ = session.Close(ctx)
		}()

		query := `MATCH (n:Movie {title: $title})
		          RETURN n.title AS title, n.tagline AS tagline, n.released AS released, n.director AS director`
		resultValue, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			result, runErr := tx.Run(ctx, query, map[string]any{"title": title})
			if runErr != nil {
				return nil, runErr
			}

			if !result.Next(ctx) {
				if result.Err() != nil {
					return nil, result.Err()
				}
				return nil, nil
			}

			return movieFromRecord(result.Record())
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if resultValue == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.Status(fiber.StatusOK).JSON(resultValue)
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
