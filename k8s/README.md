---
title: Kubernetes
keywords: [kubernetes, cloud, deployment, gcloud, aws, azure]
description: Deploying applications to Kubernetes.
---

# Kubernetes Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/k8s) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/k8s)

This project demonstrates how to deploy a Go application using the Fiber framework on a Kubernetes cluster.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package
- Docker
- Kubernetes
- kubectl
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) (for local development)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/k8s
    ```

2. Install dependencies:
    ```sh
    go get
    ```

3. Build the Docker image:
    ```sh
    docker build -t fiber-k8s-example .
    ```

4. Start Minikube (if using Minikube):
    ```sh
    minikube start
    ```

5. Deploy the application to Kubernetes:
    ```sh
    kubectl apply -f deployment.yaml
    ```

## Running the Application

1. Check the status of the pods:
    ```sh
    kubectl get pods
    ```

2. Forward the port to access the application:
    ```sh
    kubectl port-forward svc/fiber-k8s-example 3000:3000
    ```

3. Access the application at `http://localhost:3000`.

## Example

Here is an example `main.go` file for the Fiber application:

```go
package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Kubernetes!")
    })

    log.Fatal(app.Listen(":3000"))
}
```

Here is an example `Dockerfile` for the application:

```Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /fiber-k8s-example

EXPOSE 3000

CMD ["/fiber-k8s-example"]
```

Here is an example `deployment.yaml` file for deploying the application to Kubernetes:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: fiber-k8s-example
spec:
  replicas: 2
  selector:
    matchLabels:
      app: fiber-k8s-example
  template:
    metadata:
      labels:
        app: fiber-k8s-example
    spec:
      containers:
      - name: fiber-k8s-example
        image: fiber-k8s-example:latest
        ports:
        - containerPort: 3000
---
apiVersion: v1
kind: Service
metadata:
  name: fiber-k8s-example
spec:
  type: NodePort
  selector:
    app: fiber-k8s-example
  ports:
    - protocol: TCP
      port: 3000
      targetPort: 3000
      nodePort: 30001
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Docker Documentation](https://docs.docker.com/)
