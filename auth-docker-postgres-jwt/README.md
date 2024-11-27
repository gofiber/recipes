---
title: Auth + Docker + Postgres + JWT
keywords: [auth, docker, postgres, jwt]
description: Authentication with Docker, Postgres, and JWT.
---

# Auth Docker Postgres JWT Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/auth-docker-postgres-jwt) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/auth-docker-postgres-jwt)

This example demonstrates a boilerplate setup for a Go Fiber application that uses Docker, PostgreSQL, and JWT for authentication.

## Description

This project provides a starting point for building a web application with user authentication using JWT. It leverages Docker for containerization and PostgreSQL as the database.

## Requirements

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/dl/) 1.18 or higher

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/auth-docker-postgres-jwt
    ```

2. Set the environment variables in a `.env` file:
    ```env
    DB_PORT=5432
    DB_USER=example_user
    DB_PASSWORD=example_password
    DB_NAME=example_db
    SECRET=example_secret
    ```

3. Build and start the Docker containers:
    ```bash
    docker-compose build
    docker-compose up
    ```

The API and the database should now be running.

## Database Management

You can manage the database via `psql` with the following command:
```bash
docker-compose exec db psql -U <DB_USER>
```

Replace `<DB_USER>` with the value from your `.env` file.

## API Endpoints

The following endpoints are available in the API:

- **POST /api/auth/register**: Register a new user.
- **POST /api/auth/login**: Authenticate a user and return a JWT.
- **GET /api/user/:id**: Get a user (requires a valid JWT).
- **PATCH /api/user/:id**: Update a user (requires a valid JWT).
- **DELETE /api/user/:id**: Delete a user (requires a valid JWT).

## Example Usage

1. Register a new user:
    ```bash
    curl -X POST http://localhost:3000/api/auth/register -d '{"username":"testuser", "password":"testpassword"}' -H "Content-Type: application/json"
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

This example provides a basic setup for a Go Fiber application with Docker, PostgreSQL, and JWT authentication. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Docker Documentation](https://docs.docker.com)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [JWT Documentation](https://jwt.io/introduction/)
