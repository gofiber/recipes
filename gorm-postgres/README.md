---
title: GORM + PostgreSQL
keywords: [gorm, postgres, database]
description: Using GORM with PostgreSQL database.
---

# GORM with PostgreSQL Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/gorm-postgres) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/gorm-postgres)

This project demonstrates how to set up a Go application using the Fiber framework with GORM and PostgreSQL.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- [GORM](https://gorm.io/) package
- PostgreSQL

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/gorm-postgres
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up PostgreSQL and create a database:
    ```sh
    createdb go-db
    ```

4. Configure the database connection via the `DB_DSN` environment variable:
    ```sh
    export DB_DSN="host=localhost user=postgres password='' dbname=go-db port=5432 sslmode=disable"
    ```
    If `DB_DSN` is not set, the application falls back to the default DSN above.

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the application at `http://localhost:3000`.

## Endpoints

| Method | URL        | Description                |
| ------ | ---------- | -------------------------- |
| GET    | /hello     | Returns a hello message    |
| GET    | /allbooks  | Retrieves all books        |
| GET    | /book/:id  | Retrieves a book by ID     |
| POST   | /book      | Creates a new book         |
| PUT    | /book/:id  | Updates an existing book   |
| DELETE | /book/:id  | Deletes a book             |

## Example Requests

### Get All Books
```sh
curl -X GET http://localhost:3000/allbooks
```

### Get Book by ID
```sh
curl -X GET http://localhost:3000/book/1
```

### Create a New Book
```sh
curl -X POST http://localhost:3000/book \
  -d '{"title": "New Book", "author": "Author Name"}' \
  -H "Content-Type: application/json"
```

### Update a Book
```sh
curl -X PUT http://localhost:3000/book/1 \
  -d '{"title": "Updated Book", "author": "Updated Author"}' \
  -H "Content-Type: application/json"
```

### Delete a Book
```sh
curl -X DELETE http://localhost:3000/book/1
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
