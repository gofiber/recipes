---
title: Todo App + Auth + GORM
keywords: [todo app, gorm, authentication]
description: A Todo application with authentication using GORM.
---

# Todo App with Auth using GORM

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/todo-app-with-auth-gorm) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/todo-app-with-auth-gorm)

This project demonstrates a Todo application with authentication using GORM.

## Prerequisites

Ensure you have the following installed and available in your `GOPATH`:

- Golang
- [Air](https://github.com/air-verse/air) for hot reloading
- [Godotenv](https://github.com/joho/godotenv) for loading `.env` file

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/todo-app-with-auth-gorm
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Create the environment file and add a secret for signing the JWTs. `air` loads `.env`, and the application refuses to start without `TOKEN_KEY`:
    ```sh
    cp .env.example .env
    echo "TOKEN_KEY=$(openssl rand -hex 32)" >> .env
    ```

## Running the Application

1. Start the application:
    ```sh
    air
    ```

## Environment Variables

| Variable | Description | Default |
| --- | --- | --- |
| `PORT` | Server listening port | `5000` |
| `DB` | Name of the SQLite database file | `gotodo.db` |
| `TOKEN_KEY` | Secret used to sign the JWTs. Generate a random one with `openssl rand -hex 32` | required |
| `TOKEN_EXP` | Token lifetime, a [`time.ParseDuration`](https://pkg.go.dev/time#ParseDuration) string | `10h` |
