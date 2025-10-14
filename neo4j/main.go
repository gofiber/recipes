package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Movie represent the movie schema
type Movie struct {
	Title    string `json:"title" db:"title"`
	Tagline  string `json:"tagline" db:"tagline"`
	Released int64  `json:"released" db:"released"`
	Director string `json:"director" db:"director"`
}

var driver neo4j.Driver

func ConnectToDB() error {
	var err error
	driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""),
		func(conf *neo4j.Config) { conf.Encrypted = false })
	return err
}

func main() {
	if err := ConnectToDB(); err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	app := fiber.New()

	app.Post("/movie", func(c *fiber.Ctx) error {
		var movie Movie
		if err := c.BodyParser(&movie); err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		session, err := driver.Session(neo4j.AccessModeWrite)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		defer session.Close()

		query := `CREATE (n:Movie {title: $title, tagline: $tagline, released: $released, director: $director})`
		params := map[string]any{
			"title":    movie.Title,
			"tagline":  movie.Tagline,
			"released": movie.Released,
			"director": movie.Director,
		}

		_, err = session.Run(query, params)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(http.StatusCreated).JSON(movie)
	})

	app.Get("/movie/:title", func(c *fiber.Ctx) error {
		title := c.Params("title")

		session, err := driver.Session(neo4j.AccessModeRead)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}
		defer session.Close()

		query := `MATCH (n:Movie {title: $title}) RETURN n.title, n.tagline, n.released, n.director`
		result, err := session.Run(query, map[string]interface{}{"title": title})
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		if result.Next() {
			values := result.Record().Values()
			movie := Movie{
				Title:    values[0].(string),
				Tagline:  values[1].(string),
				Released: values[2].(int64),
				Director: values[3].(string),
			}
			return c.JSON(movie)
		}

		return c.SendStatus(http.StatusNotFound)
	})

	log.Fatal(app.Listen(":3000"))
}
