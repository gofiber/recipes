---
title: Asynq
keywords: [asynq, redis, queue, worker, background jobs, tasks]
description: Enqueue background jobs from Fiber and process them with an Asynq worker.
---

# Fiber and Asynq example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/asynq) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/asynq)

## Description

This example shows how to run background jobs with [Asynq](https://github.com/hibiken/asynq) (a Redis-backed task queue) from a [Fiber](https://github.com/gofiber/fiber) HTTP server. The API enqueues a job and returns immediately; a separate worker process consumes the queue and does the slow work, with retries and priorities handled by Asynq.

## How it works

- The Fiber API exposes `POST /enqueue?email=<email>&user=<id>`.
- Each request enqueues a `email:welcome` task onto Redis and returns the task id — it never does the work inline.
- A separate **worker** process pulls tasks from Redis and runs them, retrying on failure with backoff.
- The task type and payload live in a shared `task` package, so the API and the worker can't drift apart.

The worker uses weighted queues (`critical` drained ~6x as often as `low`), which is the usual way to keep a noisy low-priority job from starving important ones.

## Requirements

- [Go](https://golang.org/dl/) 1.25 or higher
- A running Redis instance (or use the provided `docker-compose.yml`)

## Running the example

### With Docker Compose

```bash
docker compose up --build
```

This starts Redis, the API, and the worker together.

### Manually

Start Redis, then in two terminals:

```bash
# terminal 1 — API
make run-api

# terminal 2 — worker
make run-worker
```

Set `REDIS_ADDR` if Redis isn't on `localhost:6379`.

## Trying it out

```bash
curl -X POST "http://localhost:3000/enqueue?email=jane@example.com&user=42"
# {"enqueued":true,"task_id":"...","queue":"default"}
```

The worker terminal logs:

```
sending welcome email to jane@example.com (user 42)
```

A request missing `email` or `user` returns `400`.

## Notes

- The handler returns an error to signal a retry; returning `asynq.SkipRetry` (as it does for a malformed payload) tells Asynq not to bother retrying something that will never succeed.
- Enqueue options like `asynq.MaxRetry` and `asynq.Queue` are set per call, so different endpoints can enqueue onto different queues with different retry policies.
