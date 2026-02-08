---
title: Swagger
keywords: [swagger, api, documentation, contrib]
description: Generate Swagger documentation for your application.
---

# Swagger API Documentation

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/swagger) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/swagger)

This project demonstrates how to integrate Swagger for API documentation in a Go application.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Swag](https://github.com/swaggo/swag) for generating Swagger docs

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/swagger
    ```

2. Install dependencies:
    ```sh
    go get -u github.com/swaggo/swag/cmd/swag
    go get -u github.com/swaggo/gin-swagger
    go get -u github.com/swaggo/files
    ```

## Generating Swagger Docs

1. Generate the Swagger documentation:
    ```sh
    swag init
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the Swagger UI:
    Open your browser and navigate to `http://localhost:8080/swagger/index.html`

## Example

Here is an example of how to document an API endpoint using Swag:

```go
// @Summary Show an account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} http.Response
// @Failure 404 {object} http.Response
// @Router /accounts/{id} [get]
func GetAccount(c *gin.Context) {
    // Your code here
}
```

## References

- [Swag Documentation](https://github.com/swaggo/swag)
- [Gin Swagger Middleware](https://github.com/swaggo/gin-swagger)
