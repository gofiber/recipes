---
title: Bootstrap
keywords: [bootstrap, gorm, validator, env]
description: A starter project with Fiber, GORM, a validator and environment configuration.
---

# Bootstrap

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/bootstrap) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/bootstrap)

A starter project ("bootstrap" as in bootstrapping a new application, not the Bootstrap CSS framework) for rapid development using Go-Fiber / Gorm / Validator.

## Components
* Fiber
  * Html Engine Template
  * Logger
  * Monitoring
* Gorm
  * PGSQL Driver
* Validator
* Env File

## Router
API Router `/api` with rate limiter middleware (the default limiter allows 5 requests per minute per client and answers further requests with `429 Too Many Requests`)
Http Router `/` with CORS and CSRF middleware

The monitoring page at `/dashboard` is not protected. It shows process and system metrics to anyone who can reach the server, and `.env.example` binds the server to `0.0.0.0`. Add authentication (for example the `basicauth` middleware) or remove the route before the app is reachable from other hosts.

## Setup

1. Copy the example env file over:
    ```
    cp .env.example .env
    ```

2. Modify the env file you just copied `.env` with the correct credentials for your database. Make sure the database you entered in `DB_NAME` has been created. The application stops with a panic when there is no `.env` file in the directory you start it from.

3. Run the API:
    ```
    go run main.go
    ```
Your api should be running at `http://localhost:4000/` if the port is in use you may modify it in the `.env` you just created.
