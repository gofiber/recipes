---
title: Google OAuth2
keywords: [oauth2, google, authentication]
description: Implementing Google OAuth2 authentication.
---

# Fiber with Google OAuth2

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/oauth2-google) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/oauth2-google)

This example demonstrates how to implement Google OAuth2 authentication in a Fiber application.

## Prerequisites

- Go 1.16 or higher
- Go modules

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/oauth2-google
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Obtain OAuth credentials from [Google Developers Console](https://console.developers.google.com/).

4. Create a `.env` file in the root directory and add your Google OAuth credentials:
    ```env
    APP_PORT=3300
    GOOGLE_CLIENT_ID=your_client_id
    GOOGLE_CLIENT_SECRET=your_client_secret
    GOOGLE_REDIRECT_URL=http://localhost:3300/api/auth/google/callback
    ```

## Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3300`.

## Endpoints

| Method | URL                          | Description                                      |
| ------ | ---------------------------- | ------------------------------------------------ |
| GET    | /api/                        | Redirects to Google login URL                    |
| GET    | /api/auth/google/callback    | Handles Google OAuth2 callback and returns user's email |

## Example Requests

### Redirect to Google Login
```sh
curl -X GET http://localhost:3300/api/
```

### Google OAuth2 Callback
```sh
curl -X GET http://localhost:3300/api/auth/google/callback?state=state&code=code
```

## Packages Used

- [Godotenv](https://github.com/joho/godotenv)
- [Fiber](https://github.com/gofiber/fiber)
- [OAuth2](https://github.com/golang/oauth2)
