---
title: Hexagonal Architecture
keywords: [hexagonal, architecture, mongodb]
description: A Hexagonal Software Architecture in Golang and MongoDB.
---

# A Hexagonal Software Architecture in Golang and MongoDB

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/hexagonal) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/hexagonal)

This project presents a simple product catalogue microservice to demonstrate the principles of a hexagonal software architecture. The microservice exposes a RESTful API that allows consuming applications to perform CRUD operations on a product catalogue. The microservice is developed in Golang, and the product catalogue data is persisted in a MongoDB repository.

![Hexagonal Architecture](Hexagonal-Arch.png)

## Prerequisites

- Go 1.27 or higher
- MongoDB
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/hexagonal
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Configure the MongoDB connection in `config/config.yaml`. The server listens on `server.port` (default `:8080`, a bare number such as `8083` works too):
    ```yaml
    database:
      url: "your_mongodb_uri"
      db: "your_db_name"
      timeout: 30
    server:
      port: ":8080"
    ```

## Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:8080`.

## Endpoints

A product has a `code` (its identifier), a `name` and a `price`.

| Method | URL                  | Description                            |
| ------ | -------------------- | -------------------------------------- |
| GET    | /products            | Retrieves all products                 |
| GET    | /products/:code      | Retrieves a product by code            |
| POST   | /products            | Creates a new product                  |
| PUT    | /products            | Updates the product with the same code |
| DELETE | /products/:code      | Deletes a product                      |

## Example Requests

### Create a New Product
```sh
curl -X POST http://localhost:8080/products -d '{"code": "p-1", "name": "New Product", "price": 100}' -H "Content-Type: application/json"
```

### Get All Products
```sh
curl -X GET http://localhost:8080/products
```

### Get Product by Code
```sh
curl -X GET http://localhost:8080/products/p-1
```

### Update a Product
```sh
curl -X PUT http://localhost:8080/products -d '{"code": "p-1", "name": "Updated Product", "price": 150}' -H "Content-Type: application/json"
```

### Delete a Product
```sh
curl -X DELETE http://localhost:8080/products/p-1
```

## Hexagonal Architecture

Hexagonal architecture, also known as ports and adapters architecture, is a design pattern used to create loosely coupled application components that can be easily connected to their software environment by means of ports and adapters. This architecture allows an application to be equally driven by users, programs, automated tests, or batch scripts, and to be developed and tested in isolation from its eventual runtime devices and databases.

## Additional Information

For more information on hexagonal architecture, you can refer to the following resources:
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
- [Hexagonal Architecture in Golang](https://medium.com/@matryer/hexagonal-architecture-in-go-2b5e0df2d8f8)
