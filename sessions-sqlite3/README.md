---
title: Sessions + SQLite3
keywords: [sessions, sqlite3, storage]
description: Using SQLite3 as a storage engine for user sessions.
---

# Sessions - SQLite3

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/sessions-sqlite3) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/sessions-sqlite3)

This example uses the SQLite3 storage package to persist user sessions. While the storage package can automatically create the sessions table at initialization, the application creates it itself at startup (see `main.go`) to add an additional "u" column. This custom column serves several purposes:

- Enables efficient querying of sessions by user identifier
- Allows tracking of multiple sessions per user
- Facilitates session cleanup for specific users

The default table schema only stores session data and expiry, making it difficult to associate sessions with specific users. The "u" column solves this limitation.

## Prerequisites

- Go 1.27 or higher
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/sessions-sqlite3
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Running the Application

The SQLite driver (`github.com/mattn/go-sqlite3`) uses cgo, so you need a C compiler such as `gcc` installed. On the first start the application creates the database file `./fiber.db` and the `sessions` table with its extra `u` column, so there is nothing to set up in SQLite beforehand.

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

## Explanation

This example uses the SQLite3 storage package to persist user sessions. The storage package can create the sessions table for you at initialization, but for the purpose of this example, `main.go` creates the table itself with an additional "u" column to better query all user-related sessions.

Fiber's session store saves each session under the key `session:<id>` in the `k` column, and encodes the session data itself. The example therefore:

- sets the `u` column with plain SQL, using the same `session:` key prefix (`sessionKeyPrefix` in `main.go`);
- lists a user's sessions by selecting the `k` column where `u` matches, and loads each one with `store.GetByID`;
- deletes another session of the same user with `store.Delete`.

It never decodes the stored `v` column by hand, because that format is internal to Fiber and can change between versions.

## Security Notes

- The UID passed from the front-end is used directly to identify the user. In production, **always validate the UID** server-side (e.g., verify credentials against your database) instead of trusting client-supplied values.
- Session type assertions (`s.Get("uid").(string)`) are guarded with an `ok` check to prevent panics on malformed or missing session data.
- The session cookie is `HttpOnly`. The front end escapes every value it inserts as HTML, because the user agent of a session is chosen by whoever logs in.
