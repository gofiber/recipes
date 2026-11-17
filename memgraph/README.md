---
title: Memgraph
keywords: [memgraph, graph, database]
description: Using Memgraph.
---

# Fiber and Memgraph

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/memgraph) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/memgraph)

This is a cookbook recipe for setting up Fiber backend and Memgraph database. 🚀

## Prerequisites

Go is an obvious prerequisite. Make sure it is installed and configured properly.

After that you need two Go packages: Fiber and Neo4j driver for Go. You can install them with the following commands:

```
go get -u github.com/gofiber/fiber/v3
go get github.com/neo4j/neo4j-go-driver/v5
```

## Run Memgraph

The easiest way to run Memgraph is to use Docker.
Once docker is installed on your machine, you can run Memgraph with the following command:

```
docker run --name memgraph -p 7687:7687 -p 7444:7444 -v mg_lib:/var/lib/memgraph memgraph/memgraph
```

Port `3000` is deliberately not published: the Fiber app in this recipe listens on it. Memgraph Lab, the web UI, is a separate image; see the [Memgraph documentation](https://memgraph.com/docs) if you want it, and map it to a different host port.

Alternatively, `docker compose up --build` starts both Memgraph and the app. Its healthcheck uses `mgconsole`, which the image contains. The app's Dockerfile starts with `FROM --platform=$BUILDPLATFORM`, which needs BuildKit, the default builder of Docker Engine 23 and later and of Docker Desktop. With an older Docker Engine set `DOCKER_BUILDKIT=1`, and with `docker-compose` v1 also `COMPOSE_DOCKER_CLI_BUILD=1`.

## Run the recipe

After you have installed all the prerequisites, you can run the recipe with the following command:

```
cd memgraph
go run ./main.go
```

This will do the following:

1. Connect Fiber backend to Memgraph database
2. Generate mock data to populate the database. **Every start first deletes all nodes in the database** (`MATCH (n) DETACH DELETE n`), so do not point the recipe at a database whose data you want to keep
3. Define two request handlers: one for getting the graph and one for getting developer nodes

## Test the recipe

Once Fiber app is running, you can test the recipe by sending a GET request to the following endpoints:

```
http://localhost:3000/graph
http://localhost:3000/developer/Andy
```

## Additional resources

For extra information use the documentation on the following links:
- Fiber: https://docs.gofiber.io/
- Memgraph: https://memgraph.com/docs
