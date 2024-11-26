---
title: Sqlc
keywords: [database, sqlc, postgresql]
---

# Fiber with sqlc

> #### 🎯 [fiber](https://github.com/gofiber/fiber) + [sqlc](https://github.com/sqlc-dev/sqlc) Example

## 👀 Usage
#### 1. Run Postgres
```bash
$ docker compose build
```
```bash
$ docker compose up
```
#### 2. Wait 1-2 minutes
```console
[+] Running 2/0
 ✔ Network sqlc_default       Created                                                                             0.1s
 ✔ Container postgres         Created                                                                             0.0s
Attaching to postgres
postgres  |
postgres  | PostgreSQL Database directory appears to contain a database; Skipping initialization
postgres  |
postgres  |
postgres  | 2023-09-28 09:17:50.737 UTC [1] LOG:  starting PostgreSQL 16.0 (Debian 16.0-1.pgdg120+1) on aarch64-unknown-linux-gnu, compiled by gcc (Debian 12.2.0-14) 12.2.0, 64-bit
postgres  | 2023-09-28 09:17:50.737 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
postgres  | 2023-09-28 09:17:50.737 UTC [1] LOG:  listening on IPv6 address "::", port 5432
postgres  | 2023-09-28 09:17:50.740 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
postgres  | 2023-09-28 09:17:50.751 UTC [30] LOG:  database system was shut down at 2023-09-28 08:50:35 UTC
postgres  | 2023-09-28 09:17:50.770 UTC [1] LOG:  database system is ready to accept connections
```
#### 3. You have to migrate the database.
> ##### 🎯 It is a "database-first" ORM as opposed to "code-first" (like gorm/gorp). That means you must first create your database schema.
> ##### 🎯 I used [golang-migrate](https://github.com/golang-migrate/migrate) to proceed with the migrate.
###### 1. Make Migration files
```bash
$ migrate create -ext sql -dir ./database/migrations -seq create_initial_table
```
```console
sqlc/database/migrations/000001_create_initial_table.up.sql
sqlc/database/migrations/000001_create_initial_table.up.sql
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
# Go 1.17 and above:
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Go 1.16 and below:
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
#### 5. Reference
[sqlc document](https://docs.sqlc.dev/en/stable/)
