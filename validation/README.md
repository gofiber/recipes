---
title: Validation
keywords: [validation, input, go-playground, validator]
---

# Validation with [Fiber](https://gofiber.io)

This example demonstrates how to use [go-playground/validator](https://github.com/go-playground/validator) for input validation in a Go Fiber application.

## Description

This project provides a basic setup for validating request data in a Go Fiber application using the `go-playground/validator` package. It includes the necessary configuration and code to perform validation on incoming requests.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Project Structure

- `main.go`: The main application entry point.
- `config/env.go`: Configuration file for environment variables.
- `go.mod`: The Go module file.
- `.env`: Environment variables file.

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

3. Create a `.env` file in the root directory with the following content:
    ```dotenv
    PORT=":8080"
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

The application should now be running on `http://localhost:8080`.

## Example Usage

1. Send a POST request to `http://localhost:8080/validate` with a JSON payload:
    ```json
    {
        "name": "John Doe",
        "email": "john.doe@example.com",
        "age": 30
    }
    ```

2. The server will validate the request data and respond with a success message if the data is valid, or an error message if the data is invalid.

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
