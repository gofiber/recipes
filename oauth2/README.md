---
title: OAuth2
keywords: [oauth2, golang, authentication, github, api]
description: Implementing GitHub OAuth2 authentication with GoFiber.
---

# OAuth2 (GitHub)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/oauth2) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/oauth2)

This project demonstrates how to implement GitHub OAuth2 authentication in a GoFiber application.

## Prerequisites

- Go 1.21+
- A [GitHub OAuth App](https://github.com/settings/developers)
  - Set **Authorization callback URL** to `http://localhost:8080/oauth/redirect`

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/oauth2
    ```

2. Copy the example env file and fill in your credentials:
    ```sh
    cp .env.example .env
    ```

3. Install dependencies:
    ```sh
    go mod download
    ```

## Running the Application

```sh
go run app.go
```

Then open `http://localhost:8080` in your browser.

## Environment Variables

Create a `.env` file in the root directory (see `.env.example`):

```shell
# GitHub OAuth2 App credentials
CLIENT_ID=your_github_client_id
CLIENT_SECRET=your_github_client_secret
```

## OAuth2 Flow

```
Browser → GET /oauth/begin
        → generates CSRF state, stores in session
        → redirects to https://github.com/login/oauth/authorize

GitHub  → GET /oauth/redirect?code=...&state=...
        → validates CSRF state
        → exchanges code for access token via GitHub API
        → stores token in session
        → redirects to /welcome.html

GET /protected → OAUTHProtected middleware checks session token
```

## Example: GitHub OAuth2 token exchange

```go
// POST https://github.com/login/oauth/access_token
// with client_id, client_secret, and code
// Response:
// {"access_token":"gho_...","token_type":"bearer","scope":""}
```

## References

- [GitHub OAuth Apps documentation](https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps)
- [GoFiber documentation](https://docs.gofiber.io)
