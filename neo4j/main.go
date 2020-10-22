package main

import (
	"fmt"
	"log"
	"strconv"

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

// ConnectToDB makes a connection with the database
func ConnectToDB() (neo4j.Session, neo4j.Driver, error) {
	// define driver, session and result vars
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	// initialize driver to connect to localhost with ID and password
	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("mdfaizan7", "mdfaizan7", ""),
		func(conf *neo4j.Config) { conf.Encrypted = false }); err != nil {
		return nil, nil, err
	}
	// Open a new session with write access
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return nil, nil, err
	}
	return session, driver, nil
}

func main() {
	// connect to database
	session, _, err := ConnectToDB()
	if err != nil {
		fmt.Print(err)
		session.Close()
	}
	defer session.Close()

	// Create a Fiber app
	app := fiber.New()

	app.Post("/movie", func(c *fiber.Ctx) error {
		movie := new(Movie)
		if err := c.BodyParser(movie); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		query := fmt.Sprintf(`CREATE (n:Movie {title:'%s', tagline:'%s', released:'%d', director:'%s' })`,
			movie.Title, movie.Tagline, movie.Released, movie.Director)

		_, err := session.Run(query, nil)
		if err != nil {
			return err
		}

		return c.SendString("Movie successfully created")
	})

	app.Get("/movie/:title", func(c *fiber.Ctx) error {
		title := c.Params("title")
		query := fmt.Sprintf(`MATCH (n:Movie {title:'%s'}) RETURN n.title, n.tagline, n.released, n.director`, title)

		result, err := session.Run(query, nil)
		if err != nil {
			return err
		}

		res := &Movie{}
		for result.Next() {
			record := result.Record()
			res.Title = record.GetByIndex(0).(string)
			res.Tagline = record.GetByIndex(1).(string)
			res.Released, _ = strconv.ParseInt(record.GetByIndex(2).(string), 10, 64)
			res.Director = record.GetByIndex(3).(string)
		}

		return c.JSON(res)
	})

	log.Fatal(app.Listen(":3000"))
}
