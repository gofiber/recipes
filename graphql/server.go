package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/recipes/graphql/graph"
	"github.com/gofiber/recipes/graphql/graph/generated"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := fiber.New()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	gqlHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.ServeHTTP(w, r)
	})
	playground := playground.Handler("GraphQL playground", "/query")

	app.All("/query", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(gqlHandler)(c.Context())
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playground)(c.Context())
		return nil
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(app.Listen(":" + port))
}
