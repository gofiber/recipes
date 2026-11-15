---
title: Sqlc
keywords: [database, sqlc, postgresql]
description: Using Sqlc to generate Go code from SQL queries.
---

# Fiber with sqlc

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/sqlc) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/sqlc)

> #### 🎯 [fiber](https://github.com/gofiber/fiber) + [sqlc](https://github.com/sqlc-dev/sqlc) Example

## Usage

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
> ##### 🎯 It is a "database-first" ORM as opposed to "code-first" (like gorm/gorp). That means you must first create your database schema.
> ##### 🎯 I used [golang-migrate](https://github.com/golang-migrate/migrate) to proceed with the migrate.
> ##### 🎯 Install the `migrate` command with `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`.
> ##### 🎯 The repository already contains this migration, so you can skip step 1 and run the migration (step 2).
###### 1. Make Migration files
```bash
$ migrate create -ext sql -dir ./database/migrations -seq create_initial_table
```
```console
sqlc/database/migrations/000001_create_initial_table.up.sql
sqlc/database/migrations/000001_create_initial_table.down.sql
```
###### 2. Migrate
```bash
$ migrate -path database/migrations -database "postgresql://user:password@localhost:5432/fiber_demo?sslmode=disable" -verbose up
```
```console
2023/09/28 20:00:00 Start buffering 1/u create_initial_table
2023/09/28 20:00:00 Read and execute 1/u create_initial_table
2023/09/28 20:00:00 Finished 1/u create_initial_table (read 24.693541ms, ran 68.30925ms)
2023/09/28 20:00:00 Finished after 100.661625ms
2023/09/28 20:00:00 Closing source and database
```
###### 3. Rollback Migrate
```bash
$ migrate -path database/migrations -database "postgresql://user:password@localhost:5432/fiber_demo?sslmode=disable" -verbose down
```
```console
2023/09/28 20:00:00 Are you sure you want to apply all down migrations? [y/N]
y
2023/09/28 20:00:00 Applying all down migrations
2023/09/28 20:00:00 Start buffering 1/d create_initial_table
2023/09/28 20:00:00 Read and execute 1/d create_initial_table
2023/09/28 20:00:00 Finished 1/d create_initial_table (read 39.681125ms, ran 66.220125ms)
2023/09/28 20:00:00 Finished after 1.83152475s
```
#### 4. Use sqlc
###### 1. Install
```bash
# Go 1.25 and above:
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Legacy Go versions (<1.25):
go get github.com/sqlc-dev/sqlc/cmd/sqlc
```
###### 2. Create a configuration file
###### Example
###### sqlc.yaml
```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "database/query"
    schema: "database/migrations"
    gen:
      go:
        package: "sqlc"
        out: "database/sqlc"
        # Lower-case JSON keys.
        emit_json_tags: true
        # Nullable columns as pointers. With the default sql.NullString/sql.NullTime types,
        # a column serializes as {"String":"...","Valid":true} and cannot be decoded from
        # {"content":"text"}. (emit_pointers_for_null_types would do this too, but it only
        # works with the pgx driver, and this recipe uses database/sql.)
        overrides:
          - column: "post.content"
            go_type:
              type: "string"
              pointer: true
          - column: "post.created_at"
            go_type:
              import: "time"
              type: "Time"
              pointer: true
          - column: "post.updated_at"
            go_type:
              import: "time"
              type: "Time"
              pointer: true
```
###### author.sql
```sql
-- name: GetAuthors :many
SELECT * FROM author;

-- name: GetAuthor :one
SELECT * FROM author WHERE id = $1;

-- name: NewAuthor :one
INSERT INTO author (email, name) VALUES ($1, $2) RETURNING *;

-- name: UpdateAuthor :one
UPDATE author SET email = $1, name = $2 WHERE id = $3 RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM author WHERE id = $1;
```
###### post.sql
```sql
-- name: GetPosts :many
SELECT * FROM post;

-- name: GetPost :one
SELECT * FROM post WHERE id = $1;

-- name: NewPost :one
INSERT INTO post (title, content, author) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdatePost :one
UPDATE post SET title = $1, content = $2, author = $3 WHERE id = $4 RETURNING *;

-- name: DeletePost :exec
DELETE FROM post WHERE id = $1;

```
###### 3. Generate
```bash
$ sqlc generate
```
```text
sqlc/
├── author.sql.go
├── db.go
├── models.go
├── post.sql.go
```
#### 5. Run the application
> ##### 🎯 The application reads the database settings from the environment (it does not read a `.env` file), so load the values of `.env.example` first.
```bash
$ set -a; source .env.example; set +a
$ go run .
```
```console
Connected to database
```
The server listens on port 3000, see the endpoints below.
#### 6. Reference
[sqlc document](https://docs.sqlc.dev/en/stable/)

## API Endpoints

All paths are below `/api/v1`, for example `http://localhost:3000/api/v1/posts`. JSON keys are lower-case (`id`, `title`, `created_at`, ...), and the nullable `content`, `created_at` and `updated_at` columns are a plain value or `null`.

### Authors

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/authors` | List all authors |
| GET | `/api/v1/authors/:id` | Get an author by ID |
| POST | `/api/v1/authors` | Create a new author |
| PUT | `/api/v1/authors/:id` | Update an author |
| DELETE | `/api/v1/authors/:id` | Delete an author |

### Posts

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/v1/posts` | List all posts |
| GET | `/api/v1/posts/:id` | Get a post by ID |
| POST | `/api/v1/posts` | Create a new post |
| PUT | `/api/v1/posts/:id` | Update a post |
| DELETE | `/api/v1/posts/:id` | Delete a post |

The answers to a request that cannot be served are `400` for an id that is not a 32-bit number, `404` when the author or post of a `GET` or `PUT` request does not exist, `422` when a post names an author that does not exist, and `409` when you delete an author who still has posts. A `DELETE` of an id that does not exist is answered with `200` like any other, because the `:exec` queries do not report whether a row was deleted. Any other failure is a `500` with the text `internal server error`. The details of the database error are only written to the server log, because the driver messages name tables and constraints.

### Example

```bash
curl -X POST http://localhost:3000/api/v1/authors -H "Content-Type: application/json" \
  -d '{"email": "ann@example.com", "name": "Ann"}'

curl -X POST http://localhost:3000/api/v1/posts -H "Content-Type: application/json" \
  -d '{"title": "Hello", "content": "plain text", "author": 1}'

curl http://localhost:3000/api/v1/posts/1
```
