---
title: AWS Elastic Beanstalk
keywords: [aws, elastic beanstalk, deploy, amazon, aws-eb]
---

# AWS Elastic Beanstalk Example

This example demonstrates how to deploy a Go Fiber application to AWS Elastic Beanstalk.

## Description

This project provides a starting point for deploying a Go Fiber application to AWS Elastic Beanstalk. It includes necessary configuration files and scripts to build and deploy the application.

## Requirements

- [AWS CLI](https://aws.amazon.com/cli/)
- [Elastic Beanstalk CLI](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/eb-cli3-install.html)
- [Go](https://golang.org/dl/) 1.18 or higher
- [Git](https://git-scm.com/downloads)

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/aws-eb
    ```

2. Initialize Elastic Beanstalk:
    ```bash
    eb init
    ```

3. Create an Elastic Beanstalk environment:
    ```bash
    eb create
    ```

4. Deploy the application:
    ```bash
    eb deploy
    ```

## Build Process

The build process is defined in the `Buildfile` and `build.sh` scripts.

- `Buildfile`:
    ```ruby
    make: ./build.sh
    ```

- `build.sh`:
    ```bash
    #!/bin/bash -xe
    # Get dependencies
    go get -u github.com/gofiber/fiber/v2

    # Build the binary
    go build -o application application.go

    # Modify permissions to make the binary executable.
    chmod +x application
    ```

## Application Code

The main application code is in `application.go`:
```go
package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
)

func main() {
    // Initialize the application
    app := fiber.New()

    // Hello, World!
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Listen and Serve on 0.0.0.0:$PORT
    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }

    log.Fatal(app.Listen(":" + port))
}
```

## .gitignore

The `.gitignore` file includes configurations to ignore Elastic Beanstalk specific files:
```plaintext
# Elastic Beanstalk Files
.elasticbeanstalk/*
!.elasticbeanstalk/*.cfg.yml
!.elasticbeanstalk/*.global.yml
```

## Conclusion

This example provides a basic setup for deploying a Go Fiber application to AWS Elastic Beanstalk. It can be extended and customized further to fit the needs of more complex applications.

## References

- [AWS Elastic Beanstalk Documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/Welcome.html)
- [Fiber Documentation](https://docs.gofiber.io)
