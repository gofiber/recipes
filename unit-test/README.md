---
title: Unit Testing
keywords: [unit testing, testing, stretchr/testify]
description: Writing unit tests for a Go Fiber application.
---

# Unit Testing Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/unit-test) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/unit-test)

This example demonstrates how to write unit tests for a Go Fiber application using the `stretchr/testify` package.

## Description

This project provides a basic setup for unit testing in a Go Fiber application. It includes examples of how to structure tests, write test cases, and use the `stretchr/testify` package for assertions.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Project Structure

- `main.go`: The main application entry point.
- `main_test.go`: The test file containing unit tests.
- `go.mod`: The Go module file.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/unit-test
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

## Running the Tests

To run the tests, use the following command:
```bash
go test ./...
```

## Example Usage

The `main.go` file sets up a simple Fiber application with a single route. The `main_test.go` file contains unit tests for this application.

### `main.go`

This file sets up a basic Fiber application with a single route that returns "OK".

### `main_test.go`

This file contains unit tests for the Fiber application. It uses the `stretchr/testify` package for assertions.

```go
package main

import (
 "io"
 "net/http"
 "testing"

 "github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
 tests := []struct {
  description string
  route string
  expectedError bool
  expectedCode int
  expectedBody string
 }{
  {
   description: "index route",
   route: "/",
   expectedError: false,
   expectedCode: 200,
   expectedBody: "OK",
  },
  {
   description: "non existing route",
   route: "/i-dont-exist",
   expectedError: false,
   expectedCode: 404,
   expectedBody: "Cannot GET /i-dont-exist",
  },
 }

 app := Setup()

 for _, test := range tests {
  req, _ := http.NewRequest("GET", test.route, nil)
  res, err := app.Test(req, -1)
  assert.Equalf(t, test.expectedError, err != nil, test.description)
  if test.expectedError {
   continue
  }
  assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
  body, err := io.ReadAll(res.Body)
  assert.Nilf(t, err, test.description)
  assert.Equalf(t, test.expectedBody, string(body), test.description)
 }
}
```

## Unit Testing in General

Unit testing is a software testing method where individual units or components of a software are tested in isolation from the rest of the application. The purpose of unit testing is to validate that each unit of the software performs as expected. Unit tests are typically automated and written by developers as part of the development process.

### Benefits of Unit Testing

- **Early Bug Detection**: Unit tests help in identifying bugs early in the development cycle.
- **Documentation**: Unit tests can serve as documentation for the code.
- **Refactoring Support**: Unit tests provide a safety net when refactoring code.
- **Design**: Writing unit tests can lead to better software design.

## Unit Testing in Fiber

Fiber is an Express-inspired web framework written in Go. Unit testing in Fiber involves testing the individual routes and handlers to ensure they behave as expected. The `stretchr/testify` package is commonly used for writing assertions in Go tests.

### Writing Unit Tests in Fiber

1. **Setup the Application**: Create a function to setup the Fiber application. This function can be reused in the tests.
2. **Define Test Cases**: Create a structure to define the input and expected output for each test case.
3. **Perform Requests**: Use the `app.Test` method to perform HTTP requests and capture the response.
4. **Assertions**: Use the `stretchr/testify` package to write assertions and verify the response.

### The `app.Test` Method

The `app.Test` method in Fiber is used to simulate HTTP requests to the Fiber application and test the responses. This is particularly useful for unit tests as it allows testing the routes and handlers of the application without starting a real server.

#### Usage of the `app.Test` Method

The `app.Test` method takes two parameters:
1. **req**: An `*http.Request` object representing the HTTP request to be tested.
2. **timeout**: An `int` value specifying the maximum time in milliseconds that the request can take. A value of `-1` disables the timeout.

The method returns an `*http.Response` and an `error`. The `*http.Response` contains the application's response to the simulated request, and the `error` indicates if any error occurred during the request processing.

#### Example

Here is an example of how the `app.Test` method is used in a unit test:

```go
package main

import (
    "io"
    "net/http"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
    // Setup the app as it is done in the main function
    app := Setup()

    // Create a new HTTP request
    req, _ := http.NewRequest("GET", "/", nil)

    // Perform the request using app.Test
    res, err := app.Test(req, -1)

    // Verify that no error occurred
    assert.Nil(t, err)

    // Verify the status code
    assert.Equal(t, 200, res.StatusCode)

    // Read the response body
    body, _ := io.ReadAll(res.Body)

    // Verify the response body
    assert.Equal(t, "OK", string(body))
}
```

In this example, a GET request is sent to the root route (`"/"`) of the application. The response is verified to ensure that the status code is `200` and the response text is `"OK"`.

## Conclusion

This example provides a basic setup for unit testing in a Go Fiber application. It can be extended and customized further to fit the needs of more complex applications.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Testify Documentation](https://github.com/stretchr/testify)
- [Go Testing](https://golang.org/pkg/testing/)
