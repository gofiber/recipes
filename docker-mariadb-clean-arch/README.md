# Docker MariaDB Clean Architecture

A sample REST application with Fiber to showcase Clean Architecture with MariaDB as a dependency with Docker.

## Prerequisites

- Docker Compose for running the application.
- Shell that supports `sh`, `make`, and `curl` for end-to-end testing. UNIX systems or WSL should work fine.
- Postman if you want to test this API with GUI.

## Application

This application is a slightly complex example of a REST API that have four major endpoints. A public user can access the `User`, `Auth`, and `Misc` major endpoints, but they cannot access the `City` endpoint (as it is protected). If one wants to access said endpoint, they have to log in first via the `Auth` endpoint, and only after that they can access the `City` endpoint. This application uses MariaDB as a database (dockerized), and JWT as an authentication mechanism. This application also showcases how to perform 1-to-many relational mapping in Clean Architecture, and also the implementation of `JOIN` SQL clause in Go in general.

## Clean Architecture

Clean architecture will be explained here.

## System Architecture

System architecture diagram / graph will be explained here.

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

Endpoints classified here are endpoints to perform authentication. In my opinion, this is framework-layer / implementation detail, so there is no 'domain' regarding this endpoint and you can use this endpoint as an enhancement to other endpoints. Authentication in this API is done using JSON Web Tokens.

- `POST /api/v1/auth/login` to log in as the user with ID of 1 in the database. Will return JWT and said JWT will be stored in a cookie.
- `POST /api/v1/auth/logout` to log out. This route removes the JWT from the cookie.
- `GET /api/v1/auth/private` to access a private route which displays information about the current (valid) JWT.

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

- Test with Postman (set the request URL to `localhost:8080`) or with the created end-to-end testing script. Keep in mind that the end-to-end script is only available for the first run. If you are trying to run it the second time, you might not be able to get all of the perfect results (because of the auto-increment in the MariaDB). Please run `make stop` first if you want to run the test suite again.

```bash
make test
```

- Teardown or stop the container. This will also delete the Docker volume created and will also delete the created image.

```bash
make stop
```

You're done!

## FAQ

Some frequently asked questions that I found scattered on the Internet. Keep in mind that the answers are mostly subjective.

**Q: Is this the right way to do Clean Architecture?**

A: Nope. There are many ways to perform clean architecture - this example being one of them. Some projects might be better than this example.

**Q: Why is authentication an implementation detail?**

A: Authentication is an implementation detail because it does not interact with the use-case or the repository / interface layer. Authentication is a bit strange that it can be implemented in any other routes as a middleware. Keep in mind that this is my subjective opinion.

**Q: Is this the recommended way to structure Fiber projects?**

A: Nope. Just like any other Gophers, I recommend you to start your project by using a single `main.go` file. Some projects do not require complicated architectures. After you start seeing the need to branch out, I recommend you to [split your code based on functional responsibilities](https://rakyll.org/style-packages/). If you need an even more strict structure, then you can try to adapt Clean Architecture or any other architectures that you see fit, such as Onion, Hexagonal, etcetera.

**Q: Is this only for Fiber?**

A: Nope. You can simply adjust `handler.go` and `middleware.go` files in order to change the presentation / frameworks-and-drivers layer to something else. You can use `net/http`, `gin-gonic`, `echo`, and many more. If you want to change or add your database, you just need to adjust the `repository.go` file accordingly. If you want to change your business logic, simply change the `service.go` file. Keep in mind that the lower layer you change, the more you have to change everything (see image above).

**Q: Is this production-ready?**

A: I try to make this as production-ready as possible 😉

## Further Improvements

Several further improvements that could be implemented in this project:

- Add more tests and mocks, especially unit tests (Clean Architecture is the best for performing unit tests).
- Add more API endpoints.
- Add a caching mechanism to the repository layer, such as Redis.
- Try to integrate S3 backend to the repository layer (MinIO is a good choice).
- Maybe add a `domain` folder in the `internal` package where we can leave the entities there?

## References

References.
