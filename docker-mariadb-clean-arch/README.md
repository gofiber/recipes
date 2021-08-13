# Docker MariaDB Clean Architecture

A sample REST application with Fiber to showcase Clean Architecture with MariaDB as a dependency with Docker.

## Prerequisites

- Docker Compose for running the application.
- Shell that supports `sh`, `make`, and `curl` for end-to-end testing. UNIX systems or WSL should work fine.
- Postman if you want to test this API with GUI.

## Application

Application will be explained here.

## Clean Architecture

Clean architecture will be explained here.

## System Architecture

System architecture diagram / graph will be explained here.

## FAQ

FAQ section here.

## API Endpoints / Features

This API is divided into four 'major endpoints', which are miscellaneous, users, authentication, and cities.

### Miscellaneous

Endpoints classified here are miscellaneous endpoints.

- `GET /api/v1` for health check.

### Users

Endpoints classified here are endpoints to perform operation on 'User' domain.

- `GET /api/v1/users` to get all users.
- `POST /api/v1/users` to create a user.
- `GET /api/v1/users/<userID>` to get a user.
- `PUT /api/v1/users/<userID>` to update a user.
- `DELETE /api/v1/users/<userID>` to delete a user.

### Authentication

Endpoints classified here are endpoints to perform authentication. In my opinion, this is framework-layer / implementation detail, so there is no 'domain' regarding this endpoint and you can use this endpoint as an enhancement to other endpoints.

- `POST /api/v1/auth/login` to log in.
- `POST /api/v1/auth/logout` to log out.
- `GET /api/v1/auth/private` to access a private route which displays information about the current (valid) JWT token.

### Cities

Endpoints classified here are endpoints to perform operation on `City` domain. **Endpoints here are protected**, so if you are going to use this endpoint, make sure you are logged in first (or at least have a valid JWT).

- `GET /api/v1/cities` to get all cities.
- `POST /api/v1/cities` to create a new city.
- `GET /api/v1/cities/<cityID>` to get a city.
- `PUT /api/v1/cities/<cityID>` to update a city.
- `DELETE /api/v1/cities/<cityID>` to delete a city.

## Installation

In order to run this application, you just need to do the following commands.

- Clone the repository.

```bash
git clone git@github.com:gofiber/recipes.git
```

- Switch to this repository.

```bash
cd recipes/docker-mariadb-clean-arch
```

- Run immediately with Docker.

```bash
make start
```

- Test with Postman or with the created end-to-end testing script.

```bash
make test
```

- Teardown or stop the container. This will also delete the Docker volume created and will also delete the created image.

```bash
make stop
```

You're done!

## References

References.
