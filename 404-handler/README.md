---
title: 404 Handler
keywords: [404, not found, handler, errorhandler, custom]
description: Custom 404 error page handling.
---

# Custom 404 Not Found Handler Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/404-handler) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/404-handler)

This example demonstrates how to implement a custom 404 Not Found handler using the [Fiber](https://gofiber.io) web framework in Go. The purpose of this example is to show how to handle requests to undefined routes gracefully by returning a 404 status code.

## Description

In web applications, it's common to encounter requests to routes that do not exist. Handling these requests properly is important to provide a good user experience and to inform the user that the requested resource is not available. This example sets up a simple Fiber application with a custom 404 handler to manage such cases.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Running the Example


1. `go mod init 404`
2. `touch main.go`
3. edit main.go
4. `go run main.go`


To run the example, use the following command:
```bash
go run main.go
```

The server will start and listen on `localhost:3000`.

## Example Routes

- **GET /hello**: Returns a simple greeting message.
- **Undefined Routes**: Any request to a route not defined will trigger the custom 404 handler.

## Custom 404 Handler

The custom 404 handler is defined to catch all undefined routes and return a 404 status code with a "Not Found" message.

## Code Overview

### `main.go`

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Fiber instance
    app := fiber.New()

    // Routes
    app.Get("/hello", hello)

    // 404 Handler
    app.Use(func(c *fiber.Ctx) error {
        return c.SendStatus(404) // => 404 "Not Found"
    })

    // Start server
    log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
    return c.SendString("I made a ☕ for you!")
}
```

## Conclusion

This example provides a basic setup for handling 404 Not Found errors in a Fiber application. It can be extended and customized further to fit the needs of more complex applications.



## Testing

1. `go get github.com/stretchr/testify/assert`
2. `go test -v`



When I submit my pull request, I am told to make the following suggested updates:

```
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Setup function to initialize the Fiber app
func setupApp() *fiber.App {
	app := fiber.New()

	// Routes
	app.Get("/hello", hello)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return app
}

func TestHelloRoute(t *testing.T) {
	// Initialize the app
	app := setupApp()

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	resp, _ := app.Test(req, -1) // -1 disables timeout

	// Check the response
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body
	body := make([]byte, resp.ContentLength)
	_, err := resp.Body.Read(body)
	if err != nil {
		t.Fatalf("Failed to read reponse body: %v", err)
	}
	
	defer resp.Body.Close()

	// Assert the response body
	assert.Equal(t, "I made a ☕ for you!", string(body))
}

func TestNotFoundRoute(t *testing.T) {
	// Initialize the app
	app := setupApp()

	// Create a test request for an unknown route
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	resp, _ := app.Test(req, -1) // -1 disables timeout

	// Check the response
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
```

However when I now run the test:


```
=== RUN   TestHelloRoute
    main_test.go:42: Failed to read reponse body: EOF
--- FAIL: TestHelloRoute (0.00s)
=== RUN   TestNotFoundRoute
--- PASS: TestNotFoundRoute (0.00s)
FAIL
exit status 1
FAIL    404     0.795s
```


## References

- [Fiber Documentation](https://docs.gofiber.io)
- [GitHub Repository](https://github.com/gofiber/fiber)




