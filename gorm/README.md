---
title: GORM
keywords: [gorm, sqlite, mysql, postgres, api, rest]
description: Using GORM with SQLite, MySQL or PostgreSQL.
---

# GORM Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/gorm) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/gorm)

This is a sample program demonstrating how to use GORM as an ORM to connect to a SQLite database with the Fiber web framework.

## Prerequisites

- Go 1.25 or higher
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/gorm
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

## Endpoints

| Method | URL              | Description                |
| ------ | ---------------- | -------------------------- |
| GET    | /api/v1/book     | Retrieves all books        |
| GET    | /api/v1/book/:id | Retrieves a book by ID     |
| POST   | /api/v1/book     | Creates a new book         |
| DELETE | /api/v1/book/:id | Deletes a book             |

## Example Requests

### Get All Books
```sh
curl -X GET http://localhost:3000/api/v1/book
```

### Get Book by ID
```sh
curl -X GET http://localhost:3000/api/v1/book/1
```

### Create a New Book
```sh
curl -X POST http://localhost:3000/api/v1/book -d '{"title": "New Book", "author": "Author Name"}' -H "Content-Type: application/json"
```

### Delete a Book
```sh
curl -X DELETE http://localhost:3000/api/v1/book/1
```

## Using MySQL or PostgreSQL

SQLite keeps this recipe runnable without a database server, but GORM talks to every database through the same API. Only the driver and the connection string change. Install the driver you need:

```sh
go get gorm.io/driver/mysql
go get gorm.io/driver/postgres
```

Then replace the `gorm.Open` call in `initDatabase` in `main.go`.

### MySQL

```go
import "gorm.io/driver/mysql"

// parseTime=True is required for time.Time fields and charset=utf8mb4 for full UTF-8 support.
// See https://github.com/go-sql-driver/mysql#dsn-data-source-name for the DSN format.
dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

### PostgreSQL

```go
import "gorm.io/driver/postgres"

dsn := "host=localhost user=postgres password='' dbname=go-db port=5432 sslmode=disable"
database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

In both cases, create the database first (for example `createdb go-db` for PostgreSQL). To keep credentials out of the source, read the connection string from the environment, for example `dsn := os.Getenv("DB_DSN")`. `AutoMigrate` then creates the `books` table on startup, as it does for SQLite.
