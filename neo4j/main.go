package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Movie represent the movie schema
type Movie struct {
	Title    string `json:"title" db:"title"`
	Tagline  string `json:"tagline" db:"tagline"`
	Released int64  `json:"released" db:"released"`
	Director string `json:"director" db:"director"`
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

		ctx := context.Background()
		session := driver.NewSession(ctx, neo4j.SessionConfig{
			DatabaseName: "movies",
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
		title := c.Params("title")
		ctx := context.Background()
		session := driver.NewSession(ctx, neo4j.SessionConfig{
			DatabaseName: "movies",
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
