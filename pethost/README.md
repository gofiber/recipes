---
title: Pethost
keywords: [pethost, deploy, cloud, docker]
description: Deploying to Pethost with one command.
---

# Pethost Deployment Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/pethost) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/pethost)

This project demonstrates how to deploy a Go application using the Fiber framework on [Pethost](https://pethost.dev).

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- [Pethost CLI](https://pethost.dev/docs/cli/)

## Setup

1. Clone the repository:

    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/pethost
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Deploy the application:

    ```sh
    pethost deploy
    ```

    The folder has no `Dockerfile`, so the command writes one that runs `go build`, deploys the application, and prints its HTTPS address.

## Running the Application

1. Open the application in your browser using the address the command printed.

## Example

See `./main.go` for the full application code. It exposes:

- `GET /` — welcome message

The application listens on port 3000. No port has to be configured: Pethost finds the port the running application listens on.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Pethost guide: deploy a Fiber app](https://pethost.dev/blog/deploy-fiber-app/)
- [Pethost CLI](https://pethost.dev/docs/cli/)
