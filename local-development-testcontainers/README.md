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
- [Docker](https://docs.docker.com/get-docker/), which Testcontainers uses to start the PostgreSQL container
- [Air](https://github.com/air-verse/air) for hot reloading
- [Godotenv](https://github.com/joho/godotenv) CLI, which `.air.conf` uses to load `.env`

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/local-development-testcontainers
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

Copy the example file and review it. The application refuses to start without `TOKEN_KEY`:

```sh
cp .env.example .env
```

| Variable | Description | Default |
| --- | --- | --- |
| `PORT` | Server listening port | `8000` |
| `DB` | PostgreSQL connection string | `postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable` |
| `TOKEN_KEY` | Secret used to sign the JWTs. Generate a random one with `openssl rand -hex 32` | required |
| `TOKEN_EXP` | Token lifetime, a [`time.ParseDuration`](https://pkg.go.dev/time#ParseDuration) string | `10h` |
| `TESTCONTAINERS_RYUK_DISABLED` | Disables the Ryuk container so the database container is not removed when the application stops | `true` in `.env.example` |
