---
title: Todo App + Auth + GORM + Testcontainers
keywords: [todo app, gorm, authentication, testcontainers, postgres]
description: A Todo application with authentication using GORM and Postgres.
---

# Todo App with Auth using GORM and Testcontainers

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/local-development-testcontainers) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/local-development-testcontainers)

This project demonstrates a Todo application with authentication using GORM and Testcontainers.

The database is a Postgres instance created using the GoFiber's [Testcontainers Service module](https://github.com/gofiber/contrib/testcontainers). The instance is reused across multiple runs of the application, allowing to develop locally without having to wait for the database to be ready.

When using the `air` command to run the application, the database is automatically started alongside the Fiber application, and it's automatically stopped when the air command is interrupted.

## Prerequisites

Ensure you have the following installed and available in your `GOPATH`:

- Golang
- [Air](https://github.com/air-verse/air) for hot reloading

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/todo-app-testcontainers-postgres
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

1. Start the application:
    ```sh
    air
    ```

## Environment Variables

Create a `.env` file in the root directory and add the following variables:

```shell
# PORT returns the server listening port
# default: 8000
PORT=

# DB returns the name of the sqlite database
# default: postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

# TOKENKEY returns the jwt token secret
TOKENKEY=

# TOKENEXP returns the jwt token expiration duration.
# Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
# default: 10h
TOKENEXP=

# TESTCONTAINERS_RYUK_DISABLED disables the Ryuk container, to avoid removing the database container when the application is stopped.
# default: true
TESTCONTAINERS_RYUK_DISABLED=true
```
