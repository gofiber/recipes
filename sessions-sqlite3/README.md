---
title: Sessions + SQLite3
keywords: [sessions, sqlite3, storage]
---

# Sessions - SQLite3

This example uses the SQLite3 storage package to persist user sessions. While the storage package can automatically create the sessions table at initialization, we create it manually to add an additional "u" column. This custom column serves several purposes:

- Enables efficient querying of sessions by user identifier
- Allows tracking of multiple sessions per user
- Facilitates session cleanup for specific users

The default table schema only stores session data and expiry, making it difficult to associate sessions with specific users. The "u" column solves this limitation.

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
