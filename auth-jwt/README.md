---
title: Auth + JWT
keywords: [auth, jwt, gorm, fiber]
description: Simple JWT authentication.
---

# Auth JWT Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/auth-jwt) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/auth-jwt)

This example demonstrates a boilerplate setup for a Go Fiber application that uses JWT for authentication.

## Description

This project provides a starting point for building a web application with user authentication using JWT. It leverages Fiber for the web framework and GORM for ORM.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/auth-jwt
    ```

2. Set the environment variables in a `.env` file:
    ```env
    DB_PORT=5432
    DB_USER=example_user
    DB_PASSWORD=example_password
    DB_NAME=example_db
    SECRET=example_secret
    ```

3. Install the dependencies:
    ```bash
    go mod download
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

The API should now be running on `http://localhost:3000`.

## Database Management

You can manage the database via `psql` with the following command:
```bash
psql -U <DB_USER> -d <DB_NAME> -h localhost -p <DB_PORT>
```

Replace `<DB_USER>`, `<DB_NAME>`, and `<DB_PORT>` with the values from your `.env` file.

## API Endpoints

The following endpoints are available in the API:

- **POST /api/auth/register**: Register a new user.
- **POST /api/auth/login**: Authenticate a user and return a JWT.
- **GET /api/user/:id**: Get a user (requires a valid JWT).
- **POST /api/user**: Create a new user.
- **PATCH /api/user/:id**: Update a user (requires a valid JWT).
- **DELETE /api/user/:id**: Delete a user (requires a valid JWT).
- **GET /api/product**: Get all products.
- **GET /api/product/:id**: Get a product.
- **POST /api/product**: Create a new product (requires a valid JWT).
- **DELETE /api/product/:id**: Delete a product (requires a valid JWT).

## Example Usage

1. Register a new user:
    ```bash
    curl -X POST http://localhost:3000/api/auth/register -d '{"username":"testuser", "password":"testpassword", "email":"test@example.com"}' -H "Content-Type: application/json"
    ```

2. Login to get a JWT:
    ```bash
    curl -X POST http://localhost:3000/api/auth/login -d '{"username":"testuser", "password":"testpassword"}' -H "Content-Type: application/json"
    ```

3. Access a protected route:
    ```bash
    curl -H "Authorization: Bearer <JWT>" http://localhost:3000/api/user/1
    ```

Replace `<JWT>` with the token received from the login endpoint.

## Conclusion

This example provides a basic setup for a Go Fiber application with JWT authentication. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [GORM Documentation](https://gorm.io/docs/)
- [JWT Documentation](https://jwt.io/introduction/)
