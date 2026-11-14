---
title: OAuth2
keywords: [oauth2, golang, authentication, github, google, api]
description: Implementing GitHub OAuth2 authentication with GoFiber, with notes for Google.
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

## Using Google or another provider

The flow is the same for any OAuth2 provider: redirect the user with a random `state`, validate the `state` on the callback, then exchange the `code` for a token. This recipe calls GitHub's token endpoint by hand. For providers with an endpoint definition in [`golang.org/x/oauth2`](https://pkg.go.dev/golang.org/x/oauth2) (Google, GitHub, Microsoft and others), the library builds the authorization URL and performs the code exchange for you.

To use Google, create OAuth credentials in the [Google Cloud console](https://console.cloud.google.com/apis/credentials) and register `http://localhost:8080/oauth/redirect` as an authorized redirect URI. Then configure the client:

```go
import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:8080/oauth/redirect",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}
```

In `OAUTHBegin`, replace the hand-built GitHub URL with the library's:

```go
authURL := googleConfig.AuthCodeURL(state)
return ctx.Redirect().Status(fiber.StatusTemporaryRedirect).To(authURL)
```

In `OAUTHRedirect`, after the `state` check, replace the manual POST to the token endpoint with an exchange, and use the token to call Google's user info endpoint:

```go
token, err := googleConfig.Exchange(ctx.Context(), code)
if err != nil {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
}

resp, err := googleConfig.Client(ctx.Context(), token).Get("https://www.googleapis.com/oauth2/v1/userinfo")
```

Keep the `state` validation. It is what protects the callback from cross-site request forgery.

## References

- [GitHub OAuth Apps documentation](https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps)
- [GoFiber documentation](https://docs.gofiber.io)
