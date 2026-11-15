---
title: MongoDB
keywords: [mongodb, database]
description: Connecting to a MongoDB database.
---

# MongoDB Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/mongodb) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/mongodb)

This project demonstrates how to connect to a MongoDB database in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- MongoDB
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/mongodb
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Start MongoDB using Docker:
    ```sh
    docker compose up -d
    ```

4. (Optional) Set the `MONGO_URI` environment variable. Defaults to `mongodb://localhost:27017/fiber_test`:
    ```sh
    export MONGO_URI="mongodb://localhost:27017/fiber_test"
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

## Endpoints

The application is a small employee API over the `employees` collection of the `fiber_test` database. `id` is the MongoDB ObjectID, as a hex string.

| Method | URL            | Description                  | Success |
| ------ | -------------- | ---------------------------- | ------- |
| GET    | /employee      | Lists all employees          | 200     |
| GET    | /employee/:id  | Gets an employee by ID       | 200     |
| POST   | /employee      | Creates an employee          | 201     |
| PUT    | /employee/:id  | Updates an employee          | 200     |
| DELETE | /employee/:id  | Deletes an employee          | 204     |

An invalid `:id` answers `400`, and an unknown one `404`.

## Example Requests

### Create an Employee
```sh
curl -X POST http://localhost:3000/employee \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "salary": 52000, "age": 31}'
```

### List Employees
```sh
curl http://localhost:3000/employee
```

### Get, Update and Delete an Employee
Use the `id` returned when you created the employee:

```sh
curl http://localhost:3000/employee/<id>

curl -X PUT http://localhost:3000/employee/<id> \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "salary": 56000, "age": 32}'

curl -X DELETE http://localhost:3000/employee/<id>
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [MongoDB Documentation](https://docs.mongodb.com)
- [MongoDB Go Driver Documentation](https://pkg.go.dev/go.mongodb.org/mongo-driver)
