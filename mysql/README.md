---
title: MySQL + PostgreSQL
keywords: [mysql, postgresql, postgres, database, sql]
description: Connecting to MySQL or PostgreSQL with database/sql.
---

# MySQL and PostgreSQL Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/mysql) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/mysql)

This project demonstrates how to connect to a MySQL database in a Go application using the Fiber framework and the standard library's `database/sql` package, without an ORM. It exposes a small CRUD API for an `employees` table. [Using PostgreSQL](#using-postgresql) instead only needs a different driver, connection string and placeholder style.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- MySQL
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/mysql
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Create the database and the table the API works on:
    ```sql
    CREATE DATABASE fiber_demo;
    USE fiber_demo;
    CREATE TABLE employees (
      id     INT AUTO_INCREMENT PRIMARY KEY,
      name   VARCHAR(255) NOT NULL,
      salary DOUBLE NOT NULL,
      age    INT NOT NULL
    );
    ```

4. Update the `user`, `password` and `dbname` constants at the top of `main.go` to match your database.

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

## Endpoints

| Method | URL            | Description                 |
| ------ | -------------- | --------------------------- |
| GET    | /employee      | Retrieves all employees     |
| GET    | /employee/:id  | Retrieves an employee by ID |
| POST   | /employee      | Creates a new employee      |
| PUT    | /employee/:id  | Updates an existing employee |
| DELETE | /employee/:id  | Deletes an employee         |

## Example Requests

### Create an Employee
```sh
curl -X POST http://localhost:3000/employee -d '{"name": "Jane Doe", "salary": 52000, "age": 31}' -H "Content-Type: application/json"
```

### Get All Employees
```sh
curl -X GET http://localhost:3000/employee
```

### Get an Employee by ID
```sh
curl -X GET http://localhost:3000/employee/1
```

### Update an Employee
```sh
curl -X PUT http://localhost:3000/employee/1 -d '{"name": "Jane Doe", "salary": 56000, "age": 32}' -H "Content-Type: application/json"
```

### Delete an Employee
```sh
curl -X DELETE http://localhost:3000/employee/1
```

## Using PostgreSQL

The handlers use only `database/sql`, so switching to PostgreSQL is a matter of changing the driver, the connection string and a few SQL details:

| | MySQL (this recipe) | PostgreSQL |
| --- | --- | --- |
| Driver import | `_ "github.com/go-sql-driver/mysql"` | `_ "github.com/jackc/pgx/v5/stdlib"` |
| Driver name for `sql.Open` | `"mysql"` | `"pgx"` |
| Connection string | `root:password@/fiber_demo` | `host=localhost port=5432 user=postgres password=password dbname=fiber_demo sslmode=disable` |
| Query placeholders | `?` | `$1`, `$2`, ... |
| ID of an inserted row | `result.LastInsertId()` | Add `RETURNING id` to the `INSERT` and read it with `db.QueryRow(...).Scan(&id)`. The pgx driver does not support `LastInsertId`. |

Create the table with PostgreSQL types:

```sql
CREATE TABLE employees (
  id     SERIAL PRIMARY KEY,
  name   TEXT NOT NULL,
  salary DOUBLE PRECISION NOT NULL,
  age    INTEGER NOT NULL
);
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [MySQL Documentation](https://dev.mysql.com/doc/)
- [Go MySQL Driver Documentation](https://pkg.go.dev/github.com/go-sql-driver/mysql)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [pgx stdlib Documentation](https://pkg.go.dev/github.com/jackc/pgx/v5/stdlib)
