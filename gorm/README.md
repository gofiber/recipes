---
title: GORM
keywords: [gorm, sqlite, api, rest]
description: Using GORM with SQLite database.
---

# GORM Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/gorm) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/gorm)

This is a sample program demonstrating how to use GORM as an ORM to connect to a SQLite database with the Fiber web framework.

## Prerequisites

- Go 1.18 or higher
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/gorm-example
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
