---
title: Clean Architecture
keywords: [clean, architecture, fiber, mongodb, go]
description: Implementing clean architecture in Go.
---

# Clean Architecture Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/clean-architecture) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/clean-architecture)

This example demonstrates a Go Fiber application following the principles of Clean Architecture.

## Description

This project provides a starting point for building a web application with a clean architecture. It leverages Fiber for the web framework, MongoDB for the database, and follows the Clean Architecture principles to separate concerns and improve maintainability.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [MongoDB](https://www.mongodb.com/try/download/community)
- [Git](https://git-scm.com/downloads)

## Project Structure

- `api/`: Contains the HTTP handlers, routes, and presenters.
- `pkg/`: Contains the core business logic and entities.
- `cmd/`: Contains the main application entry point.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/clean-architecture
    ```

2. Set the environment variables in a `.env` file:
    ```env
    DB_URI=mongodb://localhost:27017
    DB_NAME=example_db
    ```

3. Install the dependencies:
    ```bash
    go mod download
    ```

4. Run the application:
    ```bash
    go run cmd/main.go
    ```

The API should now be running on `http://localhost:3000`.

## API Endpoints

The following endpoints are available in the API:

- **GET /books**: List all books.
- **POST /books**: Add a new book.
- **PUT /books**: Update an existing book.
- **DELETE /books**: Remove a book.

## Example Usage

1. Add a new book:
    ```bash
    curl -X POST http://localhost:3000/books -d '{"title":"Book Title", "author":"Author Name"}' -H "Content-Type: application/json"
    ```

2. List all books:
    ```bash
    curl http://localhost:3000/books
    ```

3. Update a book:
    ```bash
    curl -X PUT http://localhost:3000/books -d '{"id":"<book_id>", "title":"Updated Title", "author":"Updated Author"}' -H "Content-Type: application/json"
    ```

4. Remove a book:
    ```bash
    curl -X DELETE http://localhost:3000/books -d '{"id":"<book_id>"}' -H "Content-Type: application/json"
    ```

Replace `<book_id>` with the actual ID of the book.

## Clean Architecture Principles

Clean Architecture is a software design philosophy that emphasizes the separation of concerns, making the codebase more maintainable, testable, and scalable. In this example, the Go Fiber application follows Clean Architecture principles by organizing the code into distinct layers, each with its own responsibility.

### Layers in Clean Architecture

1. **Entities (Core Business Logic)**
  - Located in the `pkg/entities` directory.
  - Contains the core business logic and domain models, which are independent of any external frameworks or technologies.

2. **Use Cases (Application Logic)**
  - Located in the `pkg/book` directory.
  - Contains the application-specific business rules and use cases. This layer orchestrates the flow of data to and from the entities.

3. **Interface Adapters (Adapters and Presenters)**
  - Located in the `api` directory.
  - Contains the HTTP handlers, routes, and presenters. This layer is responsible for converting data from the use cases into a format suitable for the web framework (Fiber in this case).

4. **Frameworks and Drivers (External Interfaces)**
  - Located in the `cmd` directory.
  - Contains the main application entry point and any external dependencies like the web server setup.

### Example Breakdown

- **Entities**: The `entities.Book` struct represents the core business model for a book.
- **Use Cases**: The `book.Service` interface defines the methods for interacting with books, such as `InsertBook`, `UpdateBook`, `RemoveBook`, and `FetchBooks`.
- **Interface Adapters**: The `handlers` package contains the HTTP handlers that interact with the `book.Service` to process HTTP requests and responses.
- **Frameworks and Drivers**: The `cmd/main.go` file initializes the Fiber application and sets up the routes using the `routes.BookRouter` function.

### Code Example

#### `entities/book.go`
```go
package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
    ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title  string             `json:"title"`
    Author string             `json:"author"`
}
```

#### `pkg/book/service.go`
```go
package book

import "clean-architecture/pkg/entities"

type Service interface {
    InsertBook(book *entities.Book) (*entities.Book, error)
    UpdateBook(book *entities.Book) (*entities.Book, error)
    RemoveBook(id primitive.ObjectID) error
    FetchBooks() ([]*entities.Book, error)
}
```

#### `api/handlers/book_handler.go`
```go
package handlers

import (
    "clean-architecture/pkg/book"
    "clean-architecture/pkg/entities"
    "clean-architecture/api/presenter"
    "github.com/gofiber/fiber/v2"
    "net/http"
    "errors"
)

func AddBook(service book.Service) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var requestBody entities.Book
        err := c.BodyParser(&requestBody)
        if err != nil {
            c.Status(http.StatusBadRequest)
            return c.JSON(presenter.BookErrorResponse(err))
        }
        if requestBody.Author == "" || requestBody.Title == "" {
            c.Status(http.StatusInternalServerError)
            return c.JSON(presenter.BookErrorResponse(errors.New("Please specify title and author")))
        }
        result, err := service.InsertBook(&requestBody)
        if err != nil {
            c.Status(http.StatusInternalServerError)
            return c.JSON(presenter.BookErrorResponse(err))
        }
        return c.JSON(presenter.BookSuccessResponse(result))
    }
}
```

#### `cmd/main.go`
```go
package main

import (
    "clean-architecture/api/routes"
    "clean-architecture/pkg/book"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    bookService := book.NewService() // Assume NewService is a constructor for the book service
    routes.BookRouter(app, bookService)
    app.Listen(":3000")
}
```

By following Clean Architecture principles, this example ensures that each layer is independent and can be modified or replaced without affecting the other layers, leading to a more maintainable and scalable application.

## Conclusion

This example provides a basic setup for a Go Fiber application following Clean Architecture principles. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [MongoDB Documentation](https://docs.mongodb.com/)
- [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)
