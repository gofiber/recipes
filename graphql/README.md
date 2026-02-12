---
title: GraphQL
keywords: [graphql]
description: Setting up a GraphQL server.
---

# GraphQL Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/graphql) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/graphql)

This project demonstrates how to set up a GraphQL server in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- [gqlgen](https://github.com/99designs/gqlgen) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/graphql
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Initialize gqlgen:
    ```sh
    go run github.com/99designs/gqlgen init
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the GraphQL playground at `http://localhost:3000/graphql`.

## Example

Here is an example `main.go` file for the Fiber application with GraphQL:

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v3"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
)

func main() {
    app := fiber.New()

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver{}}))

    app.All("/graphql", func(c fiber.Ctx) error {
        srv.ServeHTTP(c.Context().ResponseWriter(), c.Context().Request)
        return nil
    })

    app.Get("/", func(c fiber.Ctx) error {
        playground.Handler("GraphQL playground", "/graphql").ServeHTTP(c.Context().ResponseWriter(), c.Context().Request)
        return nil
    })

    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [gqlgen Documentation](https://gqlgen.com/)
- [GraphQL Documentation](https://graphql.org/)
