---
title: GeoIP
keywords: [geoip, maxmind, ip]
description: Geolocation with GeoIP.
---

# GeoIP Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/geoip) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/geoip)

This project demonstrates how to set up a GeoIP lookup service in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- [MaxMind GeoIP2](https://github.com/oschwald/geoip2-golang) package
- GeoIP2 database file (e.g., `GeoLite2-City.mmdb`)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/geoip
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Download the GeoIP2 database file and place it in the project directory.

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

2. Access the application at `http://localhost:3000`.

## Example

Here is an example `main.go` file for the Fiber application with GeoIP lookup:

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/oschwald/geoip2-golang"
    "net"
)

func main() {
    app := fiber.New()

    db, err := geoip2.Open("GeoLite2-City.mmdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    app.Get("/geoip/:ip", func(c *fiber.Ctx) error {
        ip := c.Params("ip")
        parsedIP := net.ParseIP(ip)
        record, err := db.City(parsedIP)
        if err != nil {
            return c.Status(http.StatusInternalServerError).SendString(err.Error())
        }
        return c.JSON(record)
    })

    log.Fatal(app.Listen(":3000"))
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [MaxMind GeoIP2 Documentation](https://pkg.go.dev/github.com/oschwald/geoip2-golang)
- [GeoIP2 Database](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data)
