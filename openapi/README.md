---
title: OpenAPI
keywords: [openAPI, api, documentation, huma]
description: Generate OpenAPI 3 documentation and JSON schema for your application.
---

# OpenAPI Documentation

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/openapi) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/openapi)

This project demonstrates how to add OpenAPI 3 documentation to a Go application using [Huma](https://github.com/danielgtaylor/huma).

This project got inspired by the [swagger recipe](https://github.com/gofiber/recipes/tree/master/swagger).

## Prerequisites

Ensure you have the following installed:

- Golang
- PostgreSQL, reachable with the credentials you put in `.env`

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/gofiber/recipes.git
   cd recipes/openapi
   ```

2. Download Go modules:
   ```sh
   go mod tidy
   ```

3. Copy the example environment file and set your PostgreSQL connection details (`DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_NAME`, `DB_PORT`):
   ```sh
   cp .env.example .env
   ```

## Running the Application

1. Start the application:

   ```sh
   go run main.go
   ```

2. Access the API Documentation:
   Open your browser and navigate to `http://localhost:3000/docs`.

3. OpenAPI Specs:
   - OpenAPI 3.1 JSON: `http://localhost:3000/openapi.json`.
   - OpenAPI 3.1 YAML: `http://localhost:3000/openapi.yaml`.
   - OpenAPI 3.0.3 JSON: `http://localhost:3000/openapi-3.0.json`.
   - OpenAPI 3.0.3 YAML: `http://localhost:3000/openapi-3.0.yaml`.

4. Generating TypeScript schema:

   ```sh
   npx openapi-typescript http://localhost:3000/openapi.json -o schema.ts
   ```

## Example

Here is a minimal example of adding huma to a existing Fiber codebase:

### `routes.go`

```go
import (
   ...
   "github.com/gofiber/fiber/v3"
   "github.com/danielgtaylor/huma/v2"
   "github.com/danielgtaylor/huma/v2/adapters/humafiber"
)
func New() *fiber.App {
   app := fiber.New()
   api := humafiber.New(app, huma.DefaultConfig("Book API", "1.0.0"))

   // app.Get("/books", handlers.GetAllBooks) // 👈 your existing code
   huma.Get(api, "/books", handlers.GetAllBooks) // 👈 huma version
   return app
}
```

The recipe itself registers its routes in a `/v1` group (`huma.NewGroup(api, "/v1")`), so its paths are `/v1/books` and `/v1/books/{id}`.

### `handlers/book.go`

```go
// func GetAllBooks(c fiber.Ctx) error {} // 👈 your existing code

// 👇 huma version
func GetAllBooks(ctx context.Context, _ *struct{}) (*GetAllBooksResponse, error) {
   return &GetAllBooksResponse{Body: books}, nil
}
```

## Enhancing Documentation

You can use `huma.Register` to add more information to the OpenAPI specification, such as descriptions with Markdown, examples, tags, and more.

```go
// huma.Get(group, "/books/{id}", handlers.GetBookByID)

huma.Register(api, huma.Operation{
   OperationID: "get-book-by-id",
   Method:      http.MethodGet,
   Path:        "/book/{id}",
   Summary:     "Get a book",
   Description: "Get a book by book ID.",
   Tags:        []string{"Books"},
}, handlers.GetBookByID)
```

## References

- [Huma Documentation](https://github.com/danielgtaylor/huma)
- [Huma Fiber Adapter](https://huma.rocks/features/bring-your-own-router)
- [Enhancing Documentation](https://huma.rocks/tutorial/your-first-api/#enhancing-documentation)
