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
    docker build -t fiber-k8s-example:latest .
    ```
    The image is built for the architecture of your machine. If the cluster runs on another one, for example an `amd64` cluster and an Apple silicon laptop, name it with `--platform`:
    ```sh
    docker build --platform linux/amd64 -t fiber-k8s-example:latest .
    ```
    The Dockerfile compiles the Go program for the requested platform on your machine. The commands of its last stage run in an image of that platform, so building for another architecture also needs Docker's emulation, which Docker Desktop includes.

    The Dockerfile starts with `FROM --platform=$BUILDPLATFORM`, which needs BuildKit, the default builder of Docker Engine 23 and later and of Docker Desktop. With an older Docker Engine set `DOCKER_BUILDKIT=1`.

4. Start Minikube (if using Minikube):
    ```sh
    minikube start
    ```

5. Make the image available to the cluster. Minikube runs its own container runtime, so it cannot see images built by your local Docker. Load the image into it:
    ```sh
    minikube image load fiber-k8s-example:latest
    ```
    With any other cluster, push the image to a registry the cluster can pull from and set that image name in `my-service.yaml`.

6. Deploy the application to Kubernetes:
    ```sh
    kubectl apply -f my-service.yaml
    ```

## Running the Application

1. Check the status of the pods and wait until both are `Running` and ready (`1/1`):
    ```sh
    kubectl get pods
    ```

2. Forward the port to access the application:
    ```sh
    kubectl port-forward svc/fiber-k8s-example 3000:3000
    ```
    The Service is of type `LoadBalancer`. On a cloud cluster it gets an external IP (see `kubectl get svc fiber-k8s-example`). On Minikube that IP stays `<pending>` unless you run `minikube tunnel`, so use the port-forward.

3. Access the application at `http://localhost:3000`. It answers `Hello, World!`. The readiness and liveness probes request `GET /healthz` on port 3000 of each pod.

4. Remove everything when you are done:
    ```sh
    kubectl delete -f my-service.yaml
    ```

## Example

Here is the `main.go` file of the Fiber application. It serves `/` and a `/healthz` endpoint for the probes, and it shuts down gracefully when Kubernetes sends `SIGTERM` (for example when a pod is replaced or scaled down):

```go
package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Printf("Listen error: %v\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Gracefully shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("Shutdown error: %v\n", err)
	}

	log.Println("Fiber shutdown complete.")
}
```

Here is the `Dockerfile` for the application:

```Dockerfile
FROM --platform=$BUILDPLATFORM golang:1.27-alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /k8s-example
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /bin/k8s-example .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -S k8s-example && adduser -S k8s-example -G k8s-example
USER k8s-example
WORKDIR /home/k8s-example
COPY --from=builder /bin/k8s-example ./
EXPOSE 3000
ENTRYPOINT ["./k8s-example"]
```

Here is the `my-service.yaml` file for deploying the application to Kubernetes:

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
        # The image built from this recipe's Dockerfile. On a cluster other than Minikube,
        # push it to a registry and use the full image name here, for example
        # registry.example.com/your-team/fiber-k8s-example:1.0.0
        image: fiber-k8s-example:latest
        # Use the image that is already on the node instead of pulling it from a registry.
        # Needed for an image loaded into Minikube with `minikube image load`.
        imagePullPolicy: IfNotPresent
        resources:
          requests:
            memory: "32Mi"
            cpu: "100m"
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 3000
        readinessProbe:
          httpGet:
            path: /healthz
            port: 3000
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          httpGet:
            path: /healthz
            port: 3000
          initialDelaySeconds: 10
          periodSeconds: 20

---
apiVersion: v1
kind: Service
metadata:
  name: fiber-k8s-example
spec:
  type: LoadBalancer
  selector:
    app: fiber-k8s-example
  ports:
  - port: 3000
    targetPort: 3000
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Docker Documentation](https://docs.docker.com/)
