---
title: Sessions + SQLite3
keywords: [sessions, sqlite3, storage]
---

# Sessions - SQLite3

This example demonstrates how to use Fiber sessions with the SQLite3 storage package. Run `localhost:3000` from multiple browsers to see active sessions for different users.

## Prerequisites

- Go 1.16 or higher
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

3. Create the sessions table in SQLite3:
    ```sql
    CREATE TABLE sessions (
        key TEXT PRIMARY KEY,
        data BLOB,
        expiry INTEGER,
        u TEXT
    );
    ```

## Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

## Explanation

This example uses the SQLite3 storage package to persist user sessions. The storage package can create the sessions table for you at initialization, but for the purpose of this example, the table is created manually with an additional "u" column to better query all user-related sessions.
