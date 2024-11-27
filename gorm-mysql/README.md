---
title: GORM MySQL
keywords: [gorm, mysql, database, rest, api]
description: Using GORM with MySQL database.
---

# GORM MySQL Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/gorm-mysql) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/gorm-mysql)

This is a sample program demonstrating how to use GORM as an ORM to connect to a MySQL database with the Fiber web framework.

## Prerequisites

- Go 1.16 or higher
- MySQL database
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/gorm-mysql
   ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Configure the database connection in the `config.json` file:
    ```json
    {
      "DB_Username": "your_db_username",
      "DB_Password": "your_db_password",
      "DB_Name": "your_db_name",
      "DB_Host": "localhost",
      "DB_Port": "3306"
    }
    ```

## Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

## Endpoints

| Method | URL       | Description                |
| ------ | --------- | -------------------------- |
| GET    | /hello    | Returns a hello message    |
| GET    | /allbooks | Retrieves all books        |
| GET    | /book/:id | Retrieves a book by ID     |
| POST   | /book     | Creates a new book         |
| PUT    | /book     | Updates an existing book   |
| DELETE | /book     | Deletes a book             |

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
curl -X POST http://localhost:3000/book -d '{"title": "New Book", "author": "Author Name"}' -H "Content-Type: application/json"
```

### Update a Book
```sh
curl -X PUT http://localhost:3000/book -d '{"id": 1, "title": "Updated Book", "author": "Updated Author"}' -H "Content-Type: application/json"
```

### Delete a Book
```sh
curl -X DELETE http://localhost:3000/book -d '{"id": 1}' -H "Content-Type: application/json"
```
