# Fiber and RabbitMQ example

1. Run Docker container with RabbitMQ:

```bash
make docker.rabbitmq
```

2. Wait 2-3 minutes for the RabbitMQ container to be ready to use.
3. Run Docker container with worker:

```bash
make docker.worker
```

4. Start Fiber API server (_on another console_):

```bash
make run
```

5. Go to [127.0.0.1:3000/send?msg=Hello!](http://127.0.0.1:3000/send?msg=Hello!) and see received message on worker's console, like this:

```console
2021/03/27 16:32:35 Successfully connected to RabbitMQ instance
2021/03/27 16:32:35 [*] - Waiting for messages
2021/03/27 16:32:35 [*] - Run Fiber API server and go to http://127.0.0.1:3000/send?msg=<YOUR TEXT HERE>
2021/03/27 16:33:24 Received message: Hello!
```

## How it works?

![Screenshot](https://user-images.githubusercontent.com/11155743/112727736-f8ca2200-8f34-11eb-8d40-12d9f381bd05.png)
