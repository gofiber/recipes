---
title: Monitoring with Apitally
keywords: [api, monitoring, apitally, fiber, go]
description: A simple REST API with monitoring and request logging using Apitally.
---

# Monitoring with Apitally

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/monitoring-with-apitally) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/monitoring-with-apitally)

This project showcases a simple REST API built with the Fiber framework in Go, featuring monitoring and request logging via Apitally.

[Apitally](https://apitally.io/fiber) is a lightweight monitoring and analytics tool that helps developers track API usage, performance, and errors with minimal setup.

## Prerequisites

Ensure you have Golang installed.

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/monitoring-with-apitally
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Obtain a client ID from [Apitally](https://apitally.io/fiber) by signing up and creating a new app in the dashboard.

## Running the application

1. Start the application:
    ```sh
    APITALLY_CLIENT_ID=your-client-id go run main.go
    ```

2. Make requests to the API:
    ```sh
    curl -X GET -H "Authorization: Bearer d7e123f5a2b9c4e8d6a7b2c1f5e9d3a4" http://localhost:3000/v1/books
    ```

## Dashboard

The Apitally dashboard will show the requests you've made to the API.

It provides detailed insights into the API's usage, errors, and performance. Individual requests can be inspected in the request log. You can also set up custom alerts.

![Apitally screenshots](https://assets.apitally.io/screenshots/overview.png)

## References

- [Apitally Documentation](https://docs.apitally.io/frameworks/fiber)
- [Fiber Documentation](https://docs.gofiber.io)
