---
title: Form Data
keywords: [form-data, binding, urlencoded, struct, slice, map]
description: Parsing application/x-www-form-urlencoded bodies into structs, slices and maps.
---

# Form Data with [Fiber](https://gofiber.io)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/form-data) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/form-data)

This example demonstrates how to parse `application/x-www-form-urlencoded` request bodies into Go structs, slices and maps using Fiber v3's `Bind().Body()`.

## Description

This project shows four ways a urlencoded form body can be parsed: a flat struct, a repeated field into a `[]string`, a nested struct plus a slice of structs, and dynamic keys into a map via `net/url.ParseQuery`.

## Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Project Structure

- `main.go`: The main application entry point.
- `go.mod`: The Go module file.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/form-data
    ```

2. Install the dependencies:
    ```bash
    go mod download
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

The application should now be running on `http://localhost:3000`.

## Example Usage

1. Send a POST request to `http://localhost:3000/struct` with a flat form:
    ```bash
    curl -X POST http://localhost:3000/struct -d "name=Jane&email=jane@example.com&age=30"
    ```
    ```json
    {"name":"Jane","email":"jane@example.com","age":30}
    ```

2. Send a POST request to `http://localhost:3000/slice` with a repeated field:
    ```bash
    curl -X POST http://localhost:3000/slice -d "tags=go&tags=web"
    ```
    ```json
    {"tags":["go","web"]}
    ```

3. Send a POST request to `http://localhost:3000/nested` with a nested struct and a slice of structs:
    ```bash
    curl -X POST http://localhost:3000/nested -d "customer=Jane&address.street=Main+St&address.city=Springfield&items.0.name=Widget&items.0.qty=2&items.1.name=Gadget&items.1.qty=1"
    ```
    ```json
    {"customer":"Jane","address":{"street":"Main St","city":"Springfield"},"items":[{"name":"Widget","qty":2},{"name":"Gadget","qty":1}]}
    ```

4. Send a POST request to `http://localhost:3000/map` with keys that have no matching struct field:
    ```bash
    curl -X POST http://localhost:3000/map -d "color=blue&size=M&size=L"
    ```
    ```json
    {"color":["blue"],"size":["M","L"]}
    ```

5. A field that cannot convert to its declared Go type returns a 400 with the bind error:
    ```bash
    curl -X POST http://localhost:3000/struct -d "name=Jane&email=jane@example.com&age=notanumber"
    ```
    ```json
    {"error":"bind \"age\" from body: schema: error converting value for \"age\""}
    ```

## Code Overview

### `main.go`

The main Go file sets up the Fiber application and four routes, each binding a urlencoded form body into a different Go shape:

- `POST /struct` binds a flat form into a struct with `Bind().Body()`.
- `POST /slice` binds a repeated field (`tags=go&tags=web`) into a `[]string`.
- `POST /nested` binds a dotted key path (`address.street`) into a nested struct, and a dotted, indexed key path (`items.0.qty`) into a slice of structs.
- `POST /map` parses the raw body with `net/url.ParseQuery` into a `map[string][]string`, for keys with no matching struct field. `Bind().Body()` cannot cover this case: gofiber/schema only decodes into a pointer to a struct.

## Conclusion

This example provides a minimal reference for binding `application/x-www-form-urlencoded` bodies in a Go Fiber application, covering scalar fields, repeated fields, nested structs, slices of structs, and dynamic keys.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [gofiber/schema](https://github.com/gofiber/schema)
