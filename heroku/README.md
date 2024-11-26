---
title: Heroku
keywords: [heroku, deployment]
---

# Heroku Deployment Example

This project demonstrates how to deploy a Go application using the Fiber framework on Heroku.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/heroku
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Log in to Heroku:
    ```sh
    heroku login
    ```

4. Create a new Heroku application:
    ```sh
    heroku create
    ```

5. Add a `Procfile` to the project directory with the following content:
    ```
    web: go run main.go
    ```

6. Deploy the application to Heroku:
    ```sh
    git add .
    git commit -m "Deploy to Heroku"
    git push heroku master
    ```

## Running the Application

1. Open the application in your browser:
    ```sh
    heroku open
    ```

## Example

Here is an example `main.go` file for the Fiber application:

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Heroku!")
    })

    log.Fatal(app.Listen(":" + getPort()))
}

func getPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    return port
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Heroku Documentation](https://devcenter.heroku.com/)
