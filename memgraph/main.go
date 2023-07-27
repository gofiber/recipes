package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Developer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Technology struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getDatasetCreationQueries() []string {

	//Create developer nodes
	developer_nodes := []string{
		"CREATE (n:Developer {id: 1, name:'Andy'});",
		"CREATE (n:Developer {id: 2, name:'John'});",
		"CREATE (n:Developer {id: 3, name:'Michael'});",
	}

	//Create technology nodes
	technology_nodes := []string{
		"CREATE (n:Technology {id: 1, name:'Fiber'})",
		"CREATE (n:Technology {id: 2, name:'Memgraph'})",
		"CREATE (n:Technology {id: 3, name:'Go'})",
		"CREATE (n:Technology {id: 4, name:'Neo4j'})",
		"CREATE (n:Technology {id: 5, name:'Docker'})",
		"CREATE (n:Technology {id: 6, name:'Kubernetes'})",
		"CREATE (n:Technology {id: 7, name:'Python'})",
	}
	//Create indexes on developer and technology nodes
	indexes := []string{
		"CREATE INDEX ON :Developer(id);",
		"CREATE INDEX ON :Technology(id);",
	}

	//Create relationships between developers and technologies
	edges := []string{
		"MATCH (a:Developer {id: 1}),(b:Technology {id: 1}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 1}),(b:Technology {id: 2}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 1}),(b:Technology {id: 3}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 2}),(b:Technology {id: 1}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 2}),(b:Technology {id: 2}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 2}),(b:Technology {id: 3}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 2}),(b:Technology {id: 4}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 2}),(b:Technology {id: 5}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 2}),(b:Technology {id: 6}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 2}),(b:Technology {id: 7}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 3}),(b:Technology {id: 1}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 3}),(b:Technology {id: 2}) CREATE (a)-[r:LOVES]->(b);",
		"MATCH (a:Developer {id: 3}),(b:Technology {id: 3}) CREATE (a)-[r:LOVES]->(b);",
	}

	var allQueries []string
	allQueries = append(allQueries, developer_nodes...)
	allQueries = append(allQueries, technology_nodes...)
	allQueries = append(allQueries, indexes...)
	allQueries = append(allQueries, edges...)

	return allQueries
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

	datasetCreationQueries := getDatasetCreationQueries()

	for _, query := range datasetCreationQueries {
		_, err := executeQuery(driver, query)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	app := fiber.New()

	//Get developer and technologies he loves
	app.Get("/developer/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		query := fmt.Sprintf(`MATCH (dev:Developer {name:'%s'})-[loves:LOVES]->(tech:Technology) RETURN dev, loves, tech `, name)

		result, err := executeQuery(driver, query)
		if err != nil {
			return err
		}

		res := &Developer{}

		for result.Next() {
			record := result.Record().Values()
			fmt.Println(record)
			fmt.Println(record[0].(neo4j.Node).Props()["name"])

		}
		return c.JSON(res)
	})

	// //Get whole graph
	// app.Get("/graph", func(c *fiber.Ctx) error {
	// 	query := `MATCH (dev)-[loves]->(tech) RETURN dev, loves, tech`

	// 	result, err := executeQuery(driver, query)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	for result.Next() {
	// 		record := result.Record().Values()
	// 		fmt.Println(record[0].(neo4j.Node).Props()["name"])

	// 	}
	// 	return c.JSON("Graph")
	// })
	log.Fatal(app.Listen(":3000"))

}
