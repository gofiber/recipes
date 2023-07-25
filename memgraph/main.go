package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Person struct {
	Name string `json:"name"`
}

func ConnectDriverToDB() (neo4j.Driver, error) {
	dbUri := "bolt://localhost:7687"
	var (
		driver neo4j.Driver
		err    error
	)
	if driver, err = neo4j.NewDriver(dbUri, neo4j.BasicAuth("", "", ""),
		func(conf *neo4j.Config) { conf.Encrypted = false }); err != nil {
		return nil, err
	}
	return driver, err
}

func executeQuery(driver neo4j.Driver, query string) (neo4j.Result, error) {
	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	result, err := session.Run(query, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func main() {

	driver, err := ConnectDriverToDB()
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer driver.Close()

	result, err := executeQuery(driver, "CREATE (n:Person {name:'Andy', title:'Developer'})")
	fmt.Println(err)
	fmt.Println(result)

	app := fiber.New()

	app.Get("/person/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		query := fmt.Sprintf(`MATCH (n:Person {name:'%s'}) RETURN n`, name)

		result, err := executeQuery(driver, query)
		if err != nil {
			return err
		}

		res := &Person{}
		for result.Next() {
			fmt.Println(result.Record())
			record := result.Record()
			res.Name = record.GetByIndex(0).(string)
		}

		return c.JSON(res)
	})
	log.Fatal(app.Listen(":3000"))
}
