---
title: RabbitMQ
keywords: [rabbitmq, amqp, messaging, queue]
description: Using RabbitMQ with Fiber to publish messages to a queue.
---

# Fiber and RabbitMQ example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/rabbitmq) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/rabbitmq)

## Description

This example demonstrates how to integrate [RabbitMQ](https://www.rabbitmq.com/) with a [Fiber](https://github.com/gofiber/fiber) HTTP server. The API exposes a `/send` endpoint that publishes messages to a RabbitMQ queue. A separate worker process consumes messages from the queue and prints them to the console.

## How it works

![Architecture diagram](https://user-images.githubusercontent.com/11155743/112727736-f8ca2200-8f34-11eb-8d40-12d9f381bd05.png)

- The Fiber API server connects to RabbitMQ and exposes `GET /send?msg=<text>`.
- Each request publishes the `msg` query parameter as a message to the `TestQueue` queue.
- The worker process subscribes to `TestQueue` and logs received messages.

## Prerequisites

- [Go](https://golang.org/) 1.21+
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)

## Environment variables

| Variable       | Default                                  | Description              |
|----------------|------------------------------------------|--------------------------|
| `RABBITMQ_URL` | `amqp://user:password@localhost:5672/`   | RabbitMQ connection URL  |

## Setup

### Option A: Docker Compose (recommended)

Start all services (RabbitMQ, worker, API) with a single command:

```bash
docker compose up --build
```

### Option B: Manual setup

1. Start RabbitMQ:

```bash
make docker.network
make docker.rabbitmq
```

2. Wait ~30 seconds for RabbitMQ to be ready.

3. Start the worker (in a separate terminal):

```bash
make docker.worker
```

4. Run the API server:

```bash
make run
# or
RABBITMQ_URL=amqp://user:password@localhost:5672/ go run main.go
```

## Endpoints

### `GET /send`

Publishes a message to the `TestQueue` RabbitMQ queue.

| Parameter | Type   | Required | Description          |
|-----------|--------|----------|----------------------|
| `msg`     | string | yes      | Message to publish   |

**Success response** `200 OK`:

```json
{"status": "message sent"}
```

**Error response** `400 Bad Request`:

```json
{"error": "msg parameter required"}
```

## curl examples

Send a message:

```bash
curl "http://127.0.0.1:3000/send?msg=Hello%20World"
```

Missing parameter (returns 400):

```bash
curl "http://127.0.0.1:3000/send"
```

## Worker output

When a message is received, the worker prints:

```console
2021/03/27 16:32:35 Successfully connected to RabbitMQ instance
2021/03/27 16:32:35 [*] - Waiting for messages
2021/03/27 16:32:35 [*] - Run Fiber API server and go to http://127.0.0.1:3000/send?msg=<YOUR TEXT HERE>
2021/03/27 16:33:24 Received message: Hello World
```

## RabbitMQ management dashboard

The RabbitMQ management UI is available at [http://localhost:15672](http://localhost:15672) (default credentials: `user` / `password`).

![RabbitMQ dashboard](https://user-images.githubusercontent.com/11155743/112728092-8fe3a980-8f36-11eb-9d79-be8eab26358b.png)
