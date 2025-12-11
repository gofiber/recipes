package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/graphql-go/graphql"
)

type Input struct {
	Query         string                 `query:"query"`
	OperationName string                 `query:"operationName"`
	Variables     map[string]interface{} `query:"variables"`
}

func main() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	app := fiber.New()

	// curl 'http://localhost:9090/?query=query%7Bhello%7D'
	app.Get("/", func(ctx fiber.Ctx) error {
		var input Input
		if err := ctx.Bind().Query(&input); err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString("Cannot parse query parameters: " + err.Error())
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  input.Query,
			OperationName:  input.OperationName,
			VariableValues: input.Variables,
		})

		ctx.Set("Content-Type", "application/graphql-response+json")
		return ctx.JSON(result)
	})

	// curl 'http://localhost:9090/' --header 'content-type: application/json' --data-raw '{"query":"query{hello}"}'
	app.Post("/", func(ctx fiber.Ctx) error {
		var input Input
		if err := ctx.Bind().Body(&input); err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString("Cannot parse body: " + err.Error())
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  input.Query,
			OperationName:  input.OperationName,
			VariableValues: input.Variables,
		})

		ctx.Set("Content-Type", "application/graphql-response+json")
		return ctx.JSON(result)
	})

	log.Fatal(app.Listen(":9090"))
}
