---
title: Validation
keywords: [validation, input, go-playground, validator]
description: Input validation using go-playground/validator.
---

# Validation with [Fiber](https://gofiber.io)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/validation) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/validation)

This example demonstrates how to use [go-playground/validator](https://github.com/go-playground/validator) for input validation in a Go Fiber application.

## Description

This project provides a basic setup for validating request data in a Go Fiber application using the `go-playground/validator` package. It includes the necessary configuration and code to perform validation on incoming requests.

## Requirements

- [Go](https://golang.org/dl/) 1.27 or higher
- [Git](https://git-scm.com/downloads)

## Project Structure

- `main.go`: The main application entry point.
- `config/env.go`: Configuration file for environment variables.
- `go.mod`: The Go module file.
- `.env.example`: Example environment variables file. Copy it to `.env`.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/validation
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

3. Copy the example environment file (it sets `PORT=":8080"`):
    ```bash
    cp .env.example .env
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

The application should now be running on `http://localhost:8080`.

## Example Usage

The `POST /test` endpoint reads a JSON body into a `User` and validates it. `id` and `firstname` are required, and `password` must be at least 10 characters long.

1. Send a valid request:
    ```bash
    curl -X POST http://localhost:8080/test \
      -H "Content-Type: application/json" \
      -d '{"id": 1, "firstname": "Ada", "password": "correct-horse"}'
    ```
    The server answers `200` with the validated fields. It never echoes the password:
    ```json
    {"firstname":"Ada","id":1}
    ```

2. Send an invalid request:
    ```bash
    curl -X POST http://localhost:8080/test \
      -H "Content-Type: application/json" \
      -d '{"id": 1, "firstname": "", "password": "short"}'
    ```
    The server answers `400` with a message that names each failing field and rule.

## Code Overview

### `main.go`

The main Go file sets up the Fiber application, handles HTTP requests, and performs validation using the `go-playground/validator` package.

### `config/env.go`

The configuration file for loading environment variables.

```go
package config

import "os"

// Config func to get env value
func Config(key string) string {
    return os.Getenv(key)
}
```

## Conclusion

This example provides a basic setup for validating request data in a Go Fiber application using the `go-playground/validator` package. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Validator Documentation](https://github.com/go-playground/validator)
