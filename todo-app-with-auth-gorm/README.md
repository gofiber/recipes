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

## Running the Application

1. Start the application:
    ```sh
    air
    ```

## Environment Variables

Create a `.env` file in the root directory and add the following variables:

```shell
# PORT returns the server listening port
# default: 5000
PORT=

# DB returns the name of the sqlite database
# default: gotodo.db
DB=

# TOKENKEY returns the jwt token secret
TOKENKEY=

# TOKENEXP returns the jwt token expiration duration.
# Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
# default: 10h
TOKENEXP=
```
