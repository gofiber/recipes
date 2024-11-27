---
title: Docker + Nginx
keywords: [docker, nginx, loadbalancer, reverse proxy]
description: Load balancing with Docker and Nginx.
---

# Docker + Nginx

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/docker-nginx-loadbalancer) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/docker-nginx-loadbalancer)

## Features

- **Docker and Nginx** : Deploy in docker using 5 replicas and load balancer with Nginx
- **Logger**: The application includes a request logger for monitoring HTTP requests.

## Endpoints

| Name         | Rute     | Parameters | State     | Protected | Method |
|--------------|----------| ---------- | --------- | --------- |--------|
| Hello        | /hello   | No         | Completed | No        | GET    |

## Getting Started

To get a local copy up and running, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Build the Docker image with docker compose
4. Run the Docker compose composition
  ```bash
  docker compose up --build
  ```
5. Access the application at `http://localhost:8080/hello`.

