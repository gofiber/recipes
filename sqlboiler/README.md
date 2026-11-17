---
title: Sqlboiler
keywords: [sqlboiler, database, docker]
description: Using Sqlboiler ORM.
---

# Fiber with sqlboiler

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/sqlboiler) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/sqlboiler)

> #### 🎯 [Fiber](https://github.com/gofiber/fiber) + [Sqlboiler](https://github.com/aarondl/sqlboiler) Example

## 👀 Usage
#### 1. Run Postgres

```bash
$ docker compose up -d
```
> The compose file mounts `./data/postgres/` at `/var/lib/postgresql`. Postgres 18 keeps its data in a versioned subdirectory of that path, and a mount at `/var/lib/postgresql/data` makes the container exit.
#### 2. Wait until Postgres is ready
The first start takes a few seconds. Check the log, which has to end with the line below:
```bash
$ docker compose logs postgresql
```
```console
postgres  | 2026-11-15 11:23:44.061 UTC [1] LOG:  database system is ready to accept connections
```
#### 3. You have to migrate the database.
> ###### 🎯 It is a "database-first" ORM as opposed to "code-first" (like gorm/gorp). That means you must first create your database schema.
> ###### 🎯 I used [golang-migrate](https://github.com/golang-migrate/migrate) to proceed with the migrate.
> ###### 🎯 Install the `migrate` command with `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`.
> ###### 🎯 The repository already contains this migration, so you can skip step 1 and run the migration (step 2).
###### 1. Make Migration files
```bash
$ migrate create -ext sql -dir ./migrations -seq create_initial_table
```
```console
sqlboiler/migrations/000001_create_initial_table.up.sql
sqlboiler/migrations/000001_create_initial_table.down.sql
```
###### 2. Migrate
```bash
$ migrate -path migrations -database "postgresql://user:password@localhost:5432/fiber_demo?sslmode=disable" -verbose up
```
```console
2023/09/22 20:00:00 Start buffering 1/u create_initial_table
2023/09/22 20:00:00 Read and execute 1/u create_initial_table
2023/09/22 20:00:00 Finished 1/u create_initial_table (read 24.693541ms, ran 68.30925ms)
2023/09/22 20:00:00 Finished after 100.661625ms
2023/09/22 20:00:00 Closing source and database
```
###### 3. Rollback Migrate
```bash
$ migrate -path migrations -database "postgresql://user:password@localhost:5432/fiber_demo?sslmode=disable" -verbose down
```
```console
2023/09/22 20:00:00 Are you sure you want to apply all down migrations? [y/N]
y
2023/09/22 20:00:00 Applying all down migrations
2023/09/22 20:00:00 Start buffering 1/d create_initial_table
2023/09/22 20:00:00 Read and execute 1/d create_initial_table
2023/09/22 20:00:00 Finished 1/d create_initial_table (read 39.681125ms, ran 66.220125ms)
2023/09/22 20:00:00 Finished after 1.83152475s
```
#### 4. Use sqlboiler
###### 1. Install
```bash
# Go 1.25 and above:
$ go install github.com/aarondl/sqlboiler/v4@latest
$ go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-psql@latest
```
> ###### 🎯 `sqlboiler` runs the `sqlboiler-psql` driver, so `$(go env GOPATH)/bin` has to be in your `PATH`.
###### 2. Create a configuration file
> ###### 🎯 The configuration file should be named sqlboiler.toml
###### Example
```toml
output   = "models"
wipe     = true
no-tests = true
no-hooks = true

[psql]
  dbname = "fiber_demo"
  host   = "localhost"
  port   = 5432
  user   = "user"
  pass   = "password"
  sslmode   = "disable"
```
> ###### 🎯 The Postgres container has no SSL, and `sqlboiler` uses `sslmode = "require"` by default. Without `sslmode = "disable"` it stops with `pq: SSL is not enabled on the server`.
> ###### 🎯 Leave `schema` out unless your tables are not in the default `public` schema. A schema that does not exist ends with `no tables found in database`.
###### 3. Create models
> ###### 🎯 After creating a configuration file that points at the database we want to generate models for, we can invoke the sqlboiler command line utility.
```bash
$ sqlboiler psql
```
```text
models/
├── author.go
├── boil_queries.go
├── boil_table_names.go
├── boil_types.go
├── boil_view_names.go
├── post.go
├── psql_upsert.go
├── schema_migrations.go
```
#### 5. Run the application
> ###### 🎯 The application reads the database settings from the environment (it does not read a `.env` file), so load the values of `.env.example` first.
```bash
$ set -a; source .env.example; set +a
$ go run .
```
```console
Connected to database
```
#### 6. Try the API
The server listens on port 3000.

| Method | Route | Description |
| --- | --- | --- |
| `GET` | `/api/v1/authors` | List the authors |
| `GET` | `/api/v1/authors/:id` | Get an author |
| `POST` | `/api/v1/authors` | Create an author (`name`, `email`) |
| `PUT` | `/api/v1/authors/:id` | Rename an author (`name`) |
| `DELETE` | `/api/v1/authors/:id` | Delete an author who has no posts |
| `GET` | `/api/v1/posts` | List the posts |
| `GET` | `/api/v1/posts/:id` | Get a post |
| `POST` | `/api/v1/posts` | Create a post (`title`, `content`, `author` is the id of an existing author) |
| `PUT` | `/api/v1/posts/:id` | Update the `title` and `content` of a post |
| `DELETE` | `/api/v1/posts/:id` | Delete a post |

The answers to a request that cannot be served are `400` for an id that is not a 32-bit number, `404` when the author or post does not exist, `422` when a post names an author that does not exist, and `409` when you delete an author who still has posts. Any other failure is a `500` with the text `internal server error`. The details of the database error are only written to the server log, because the driver messages name tables and constraints.

```bash
$ curl -X POST localhost:3000/api/v1/authors -H 'Content-Type: application/json' -d '{"name":"Ada","email":"ada@example.com"}'
```
```json
{"id":1,"email":"ada@example.com","name":"Ada"}
```
```bash
$ curl -X POST localhost:3000/api/v1/posts -H 'Content-Type: application/json' -d '{"title":"Hello","content":"First post","author":1}'
```
```json
{"id":1,"title":"Hello","content":"First post","created_at":"2026-11-15T11:23:08.477667322Z","updated_at":"2026-11-15T11:23:08.477667322Z","author":1}
```
```bash
$ curl localhost:3000/api/v1/posts/1
```
```json
{"id":1,"title":"Hello","content":"First post","created_at":"2026-11-15T11:23:08.477667Z","updated_at":"2026-11-15T11:23:08.477667Z","author":1}
```
