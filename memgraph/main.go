package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neo4j/neo4j-go-driver/neo4j" // Memgraph is compatible with Neo4j GO driver, and you can use it to connect to Memgraph
)

// Developer represents a developer node in the graph
type Developer struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

// Loves represents a relationship between developer and technology
type Loves struct {
	Label string `json:"label"`
}

// Technology represents a technology node in the graph
type Technology struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

// Queries to create the mock dataset
func getDatasetCreationQueries() []string {
	// Create developer nodes
	developer_nodes := []string{
		"CREATE (n:Developer {id: 1, name:'Andy'});",
		"CREATE (n:Developer {id: 2, name:'John'});",
		"CREATE (n:Developer {id: 3, name:'Michael'});",
	}

	// Create technology nodes
	technology_nodes := []string{
		"CREATE (n:Technology {id: 1, name:'Fiber'})",
		"CREATE (n:Technology {id: 2, name:'Memgraph'})",
		"CREATE (n:Technology {id: 3, name:'Go'})",
		"CREATE (n:Technology {id: 4, name:'Neo4j'})",
		"CREATE (n:Technology {id: 5, name:'Docker'})",
		"CREATE (n:Technology {id: 6, name:'Kubernetes'})",
		"CREATE (n:Technology {id: 7, name:'Python'})",
	}
	// Create indexes on developer and technology nodes
	indexes := []string{
		"CREATE INDEX ON :Developer(id);",
		"CREATE INDEX ON :Technology(id);",
	}

	// Create relationships between developers and technologies
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

	// Create a single list of all queries for sake of simplicity
	var allQueries []string = []string{"MATCH (n) DETACH DELETE n;"}
	allQueries = append(allQueries, developer_nodes...)
	allQueries = append(allQueries, technology_nodes...)
	allQueries = append(allQueries, indexes...)
	allQueries = append(allQueries, edges...)

	return allQueries
}

func ConnectDriverToDB() (neo4j.Driver, error) {
	// Memgraph communicates via Bolt protocol, using port 7687
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
	// Each session opens a new connection to the database and gets a thread from the thread pool, for multi-threaded access to Memgraph open multiple sessions
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
	// Connect to Memgraph
	driver, err := ConnectDriverToDB()
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer driver.Close()

	// Create mock dataset
	datasetCreationQueries := getDatasetCreationQueries()
	for _, query := range datasetCreationQueries {
		_, err := executeQuery(driver, query)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	// Create a Fiber app
	app := fiber.New()

	// Get developer called Andy and technologies he loves, http://localhost:3000/developer/Andy
	app.Get("/developer/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		query := fmt.Sprintf(`MATCH (dev:Developer {name:'%s'})-[loves:LOVES]->(tech:Technology) RETURN dev, loves, tech `, name)

		result, err := executeQuery(driver, query)
		if err != nil {
			return err
		}

		developer := Developer{}
		loves := Loves{}
		technologies := []Technology{}
		technology := Technology{}

		for result.Next() {
			record := result.Record().Values()
			for _, value := range record {
				switch v := value.(type) {
				case neo4j.Node:
					if value.(neo4j.Node).Labels()[0] == "Developer" {
						developer.Id = value.(neo4j.Node).Props()["id"].(int64)
						developer.Label = value.(neo4j.Node).Labels()[0]
						developer.Name = value.(neo4j.Node).Props()["name"].(string)

					} else if value.(neo4j.Node).Labels()[0] == "Technology" {
						technology.Id = value.(neo4j.Node).Props()["id"].(int64)
						technology.Label = value.(neo4j.Node).Labels()[0]
						technology.Name = value.(neo4j.Node).Props()["name"].(string)
						technologies = append(technologies, technology)
					} else {
						fmt.Println("Unknown Node type")
					}
				case neo4j.Relationship:
					if value.(neo4j.Relationship).Type() == "LOVES" {
						loves.Label = value.(neo4j.Relationship).Type()
					} else {
						fmt.Println("Unknown Relationship type")
					}
				default:
					fmt.Printf("I don't know about type %T!\n", v)
				}
			}
		}

		data := map[string]interface{}{
			"developer":    developer,
			"loves":        loves,
			"technologies": technologies,
		}

		return c.JSON(data)
	})

	// Get whole graph, including all nodes, and edges, http://localhost:3000/graph
	app.Get("/graph", func(c *fiber.Ctx) error {
		query := `MATCH (dev)-[loves]->(tech) RETURN dev, loves, tech`

		result, err := executeQuery(driver, query)
		if err != nil {
			return err
		}
		developer := Developer{}
		love := Loves{}
		technology := Technology{}
		developer_nodes := []Developer{}
		love_edges := []Loves{}
		technology_nodes := []Technology{}
		for result.Next() {
			record := result.Record().Values()
			for _, value := range record {
				switch v := value.(type) {
				case neo4j.Node:
					if value.(neo4j.Node).Labels()[0] == "Developer" {
						developer.Id = value.(neo4j.Node).Props()["id"].(int64)
						developer.Label = value.(neo4j.Node).Labels()[0]
						developer.Name = value.(neo4j.Node).Props()["name"].(string)
						developer_nodes = append(developer_nodes, developer)

					} else if value.(neo4j.Node).Labels()[0] == "Technology" {
						technology.Id = value.(neo4j.Node).Props()["id"].(int64)
						technology.Label = value.(neo4j.Node).Labels()[0]
						technology.Name = value.(neo4j.Node).Props()["name"].(string)
						technology_nodes = append(technology_nodes, technology)
					} else {
						fmt.Println("Unknown Node type")
					}
				case neo4j.Relationship:
					if value.(neo4j.Relationship).Type() == "LOVES" {
						love.Label = value.(neo4j.Relationship).Type()
						love_edges = append(love_edges, love)
					} else {
						fmt.Println("Unknown Relationship type")
					}
				default:
					fmt.Printf("I don't know about type %T!\n", v)
				}
			}
		}

		data := map[string]interface{}{
			"developer":  developer_nodes,
			"loves":      love_edges,
			"technology": technology_nodes,
		}
		return c.JSON(data)
	})
	log.Fatal(app.Listen(":3000"))
}
