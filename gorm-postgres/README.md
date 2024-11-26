---
title: GORM + PostgreSQL
keywords: [gorm, postgres, database]
---

# GORM with PostgreSQL Example

This project demonstrates how to set up a Go application using the Fiber framework with GORM and PostgreSQL.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- [GORM](https://gorm.io/) package
- PostgreSQL

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/gorm-postgres
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Set up PostgreSQL and create a database:
    ```sh
    createdb mydb
    ```

4. Update the database connection string in the code if necessary.

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the application at `http://localhost:3000`.

## Example

Here is an example `main.go` file for the Fiber application with GORM and PostgreSQL:

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"size:255"`
    Email string `gorm:"uniqueIndex"`
}

func main() {
    dsn := "host=localhost user=youruser password=yourpassword dbname=mydb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    db.AutoMigrate(&User{})

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, GORM with PostgreSQL!")
    })

    app.Post("/users", func(c *fiber.Ctx) error {
        user := new(User)
        if err := c.BodyParser(user); err != nil {
            return c.Status(400).SendString(err.Error())
        }
        db.Create(user)
        return c.JSON(user)
    })

    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
