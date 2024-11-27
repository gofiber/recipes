---
title: RabbitMQ
keywords: [rabbitmq, amqp, messaging, queue]
description: Using RabbitMQ.
---

# Fiber and RabbitMQ example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/rabbitmq) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/rabbitmq)

1. Create Docker network:

```bash
make docker.network
```

2. Run Docker container with RabbitMQ:

```bash
make docker.rabbitmq
```

3. Wait 2-3 minutes for the RabbitMQ container to be ready to use.
4. Run Docker container with worker:

```bash
make docker.worker
```

5. Start Fiber API server (_on another console_):

```bash
make run
```

6. Go to [127.0.0.1:3000/send?msg=Hello!](http://127.0.0.1:3000/send?msg=Hello!) and see received message on worker's console, like this:

```console
2021/03/27 16:32:35 Successfully connected to RabbitMQ instance
2021/03/27 16:32:35 [*] - Waiting for messages
2021/03/27 16:32:35 [*] - Run Fiber API server and go to http://127.0.0.1:3000/send?msg=<YOUR TEXT HERE>
2021/03/27 16:33:24 Received message: Hello!
```

7. Also, you can see useful RabbitMQ dashboard at [localhost:15672](http://localhost:15672):

![Screenshot](https://user-images.githubusercontent.com/11155743/112728092-8fe3a980-8f36-11eb-9d79-be8eab26358b.png)

## How it works?

![Screenshot](https://user-images.githubusercontent.com/11155743/112727736-f8ca2200-8f34-11eb-8d40-12d9f381bd05.png)
