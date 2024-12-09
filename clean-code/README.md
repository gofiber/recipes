---
title: Clean Code
keywords: [clean, code, fiber, postgres, go]
description: Implementing clean code in Go.
---

# Clean Code Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/clean-code) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/clean-code)

This is an example of a RESTful API built using the Fiber framework (https://gofiber.io/) and PostgreSQL as the database.

## Description of Clean Code

Clean code is a philosophy and set of practices aimed at writing code that is easy to understand, maintain, and extend. Key principles of clean code include:

- **Readability**: Code should be easy to read and understand.
- **Simplicity**: Avoid unnecessary complexity.
- **Consistency**: Follow consistent coding standards and conventions.
- **Modularity**: Break down code into small, reusable, and independent modules.
- **Testability**: Write code that is easy to test.

This Fiber app is a good example of clean code because:

- **Modular Structure**: The code is organized into distinct modules, making it easy to navigate and understand.
- **Clear Separation of Concerns**: Different parts of the application (e.g., routes, handlers, services) are clearly separated, making the codebase easier to maintain and extend.
- **Error Handling**: Proper error handling is implemented to ensure the application behaves predictably.

## Start

1. Build and start the containers:
    ```sh
    docker compose up --build
    ```

1. The application should now be running and accessible at `http://localhost:3000`.
   
## Endpoints

- `GET /api/v1/books`: Retrieves a list of all books.
  ```sh
  curl -X GET http://localhost:3000/api/v1/books
  ```

- `POST /api/v1/books`: Adds a new book to the collection.
  ```sh
  curl -X POST http://localhost:3000/api/v1/books \
       -H "Content-Type: application/json" \
       -d '{"title":"Title"}'
  ```
