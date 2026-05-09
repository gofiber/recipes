---
title: Swagger
keywords: [swagger, api, documentation, contrib]
description: Generate Swagger documentation for your application.
---

# Swagger API Documentation

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/swagger) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/swagger)

This project demonstrates how to integrate Swagger for API documentation in a [Fiber](https://github.com/gofiber/fiber) application using [`gofiber/contrib/swaggerui`](https://github.com/gofiber/contrib/tree/main/swaggerui) and [`swaggo/swag`](https://github.com/swaggo/swag).

## Prerequisites

- Go 1.21+
- [Swag CLI](https://github.com/swaggo/swag) for generating Swagger docs
- PostgreSQL (connection configured via environment variables)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/swagger
    ```

2. Install the Swag CLI:
    ```sh
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

3. Configure the database connection via environment variables:
    ```sh
    export DB_USER=postgres
    export DB_PASSWORD=secret
    export DB_HOST=localhost
    export DB_NAME=books
    export DB_PORT=5432
    ```

## Generating Swagger Docs

Generate (or regenerate) the Swagger documentation from the annotations in your source code:

```sh
swag init
```

This writes `docs/docs.go`, `docs/swagger.json`, and `docs/swagger.yaml`.

## Running the Application

```sh
go run main.go
```

The application starts on port **3000**.

## Accessing the Swagger UI

Open your browser and navigate to:

```
http://localhost:3000/docs/docs
```

The Swagger UI is served by [`gofiber/contrib/v3/swaggerui`](https://github.com/gofiber/contrib/tree/main/swaggerui) and reads the spec from `./docs/swagger.json`.

## API Endpoints

All routes are prefixed with `/api/v1`.

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/books` | List all books |
| GET | `/api/v1/books/:id` | Get a book by ID |
| POST | `/api/v1/books` | Register a new book |
| DELETE | `/api/v1/books/:id` | Delete a book by ID |

## Example

Annotate your Fiber handler functions with Swag comments to generate docs:

```go
// GetBookByID returns a book by ID
// @Summary Get book by ID
// @Description Get a single book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} ResponseHTTP{data=models.Book}
// @Failure 404 {object} ResponseHTTP{}
// @Router /v1/books/{id} [get]
func GetBookByID(c fiber.Ctx) error {
    // Your code here
}
```

## curl Examples

```sh
# List all books
curl http://localhost:3000/api/v1/books

# Get book by ID
curl http://localhost:3000/api/v1/books/1

# Create a book
curl -X POST http://localhost:3000/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":"The Go Programming Language","author":"Donovan & Kernighan"}'

# Delete a book
curl -X DELETE http://localhost:3000/api/v1/books/1
```

## References

- [Fiber](https://github.com/gofiber/fiber)
- [gofiber/contrib/swaggerui](https://github.com/gofiber/contrib/tree/main/swaggerui)
- [Swag Documentation](https://github.com/swaggo/swag)
