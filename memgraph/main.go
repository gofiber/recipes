package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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

// getDatasetCreationQueries returns Cypher queries to build the mock dataset.
func getDatasetCreationQueries() []string {
	developer_nodes := []string{
		"CREATE (n:Developer {id: 1, name:'Andy'});",
		"CREATE (n:Developer {id: 2, name:'John'});",
		"CREATE (n:Developer {id: 3, name:'Michael'});",
	}

	technology_nodes := []string{
		"CREATE (n:Technology {id: 1, name:'Fiber'})",
		"CREATE (n:Technology {id: 2, name:'Memgraph'})",
		"CREATE (n:Technology {id: 3, name:'Go'})",
		"CREATE (n:Technology {id: 4, name:'Neo4j'})",
		"CREATE (n:Technology {id: 5, name:'Docker'})",
		"CREATE (n:Technology {id: 6, name:'Kubernetes'})",
		"CREATE (n:Technology {id: 7, name:'Python'})",
	}

	indexes := []string{
		"CREATE INDEX ON :Developer(id);",
		"CREATE INDEX ON :Technology(id);",
	}

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

	allQueries := []string{"MATCH (n) DETACH DELETE n;"}
	allQueries = append(allQueries, developer_nodes...)
	allQueries = append(allQueries, technology_nodes...)
	allQueries = append(allQueries, indexes...)
	allQueries = append(allQueries, edges...)
	return allQueries
}

func connectDriver(ctx context.Context) (neo4j.DriverWithContext, error) {
	boltURL := os.Getenv("BOLT_URL")
	if boltURL == "" {
		boltURL = "bolt://localhost:7687"
	}
	// bolt:// scheme means no TLS; no extra config needed
	return neo4j.NewDriverWithContext(boltURL, neo4j.BasicAuth("", "", ""))
}

func executeQuery(ctx context.Context, driver neo4j.DriverWithContext, query string, params map[string]any) (*neo4j.EagerResult, error) {
	return neo4j.ExecuteQuery(ctx, driver, query, params, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithWritersRouting())
}

func nodeToDeveloper(node neo4j.Node) Developer {
	return Developer{
		Id:    node.Props["id"].(int64),
		Name:  node.Props["name"].(string),
		Label: node.Labels[0],
	}
}

func nodeToTechnology(node neo4j.Node) Technology {
	return Technology{
		Id:    node.Props["id"].(int64),
		Name:  node.Props["name"].(string),
		Label: node.Labels[0],
	}
}

func main() {
	ctx := context.Background()

	driver, err := connectDriver(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close(ctx)

	// Seed mock dataset
	for _, query := range getDatasetCreationQueries() {
		if _, err := executeQuery(ctx, driver, query, nil); err != nil {
			log.Fatal(err)
		}
	}

	app := fiber.New()

	// GET /developer/:name — return developer and technologies they love
	app.Get("/developer/:name", func(c fiber.Ctx) error {
		name := c.Params("name")

		result, err := executeQuery(c.Context(), driver,
			`MATCH (dev:Developer {name: $name})-[loves:LOVES]->(tech:Technology) RETURN dev, loves, tech`,
			map[string]any{"name": name},
		)
		if err != nil {
			return err
		}

		developer := Developer{}
		loves := Loves{}
		technologies := []Technology{}

		for _, record := range result.Records {
			if dev, ok := record.Get("dev"); ok {
				developer = nodeToDeveloper(dev.(neo4j.Node))
			}
			if l, ok := record.Get("loves"); ok {
				loves.Label = l.(neo4j.Relationship).Type
			}
			if tech, ok := record.Get("tech"); ok {
				technologies = append(technologies, nodeToTechnology(tech.(neo4j.Node)))
			}
		}

		return c.JSON(map[string]any{
			"developer":    developer,
			"loves":        loves,
			"technologies": technologies,
		})
	})

	// GET /graph — return entire graph
	app.Get("/graph", func(c fiber.Ctx) error {
		result, err := executeQuery(c.Context(), driver,
			`MATCH (dev)-[loves]->(tech) RETURN dev, loves, tech`,
			nil,
		)
		if err != nil {
			return err
		}

		developerNodes := []Developer{}
		loveEdges := []Loves{}
		technologyNodes := []Technology{}

		for _, record := range result.Records {
			if dev, ok := record.Get("dev"); ok {
				developerNodes = append(developerNodes, nodeToDeveloper(dev.(neo4j.Node)))
			}
			if l, ok := record.Get("loves"); ok {
				loveEdges = append(loveEdges, Loves{Label: l.(neo4j.Relationship).Type})
			}
			if tech, ok := record.Get("tech"); ok {
				technologyNodes = append(technologyNodes, nodeToTechnology(tech.(neo4j.Node)))
			}
		}

		return c.JSON(map[string]any{
			"developer":  developerNodes,
			"loves":      loveEdges,
			"technology": technologyNodes,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
