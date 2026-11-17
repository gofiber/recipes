---
title: JWT
keywords: [jwt, json web token, authentication]
description: Using JSON Web Tokens (JWT) for authentication.
---

# Fiber with JWT

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/jwt) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/jwt)

This example demonstrates how to use JSON Web Tokens (JWT) for authentication in a Fiber application.

## Prerequisites

- Go 1.27 or higher
- Go modules
- PostgreSQL. The product endpoints store their data there.

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/jwt
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Copy the sample environment file and fill it in. `SECRET` signs and verifies the tokens (for example `openssl rand -hex 32`). The application refuses to start without it, and it also needs a `.env` file to be present:
    ```sh
    cp .env.sample .env
    ```
    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=your_password
    DB_NAME=your_database
    SECRET=your_secret
    ```

## Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

## Endpoints

| Method | URL               | Auth | Description                                    |
| ------ | ----------------- | ---- | ---------------------------------------------- |
| GET    | /api              | no   | Status check                                   |
| POST   | /api/auth/login   | no   | Returns a JWT for the demo user                |
| GET    | /api/product      | no   | Lists all products                             |
| GET    | /api/product/:id  | no   | Gets a product by ID                           |
| POST   | /api/product      | JWT  | Creates a product                              |
| DELETE | /api/product/:id  | JWT  | Deletes a product                              |

The login handler accepts one hard-coded demo user, `ender` with the password `ender`. A real application checks the credentials against its own user store.

## Example Requests

### Login
```sh
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"identity": "ender", "password": "ender"}'
```

The token is in the `data` field of the response.

### Create a Product
```sh
curl -X POST http://localhost:3000/api/product \
  -H "Authorization: Bearer <your_jwt_token>" \
  -H "Content-Type: application/json" \
  -d '{"title": "Widget", "description": "A widget", "amount": 3}'
```

### Get a Product
```sh
curl http://localhost:3000/api/product/1
```

A request to a protected route without a token answers `400`, and one with an invalid or expired token answers `401`.
