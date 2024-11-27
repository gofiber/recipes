---
title: OAuth2
keywords: [oauth2, golang, authentication, api]
description: Implementing OAuth2 authentication.
---

# OAuth2

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/oauth2) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/oauth2)

This project demonstrates how to implement OAuth2 authentication in a Go application.

## Prerequisites

Ensure you have the following installed:

- Golang
- [OAuth2](https://github.com/golang/oauth2) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/oauth2
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

## Environment Variables

Create a `.env` file in the root directory and add the following variables:

```shell
# CLIENT_ID is the OAuth2 client ID
CLIENT_ID=

# CLIENT_SECRET is the OAuth2 client secret
CLIENT_SECRET=

# REDIRECT_URL is the OAuth2 redirect URL
REDIRECT_URL=

# AUTH_URL is the OAuth2 authorization URL
AUTH_URL=

# TOKEN_URL is the OAuth2 token URL
TOKEN_URL=
```

## Example

Here is an example of how to set up an OAuth2 configuration:

```go
package main

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

func main() {
    conf := &oauth2.Config{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
        RedirectURL:  "your-redirect-url",
        Endpoint:     google.Endpoint,
    }

    // Your code here
}
```

## References

- [OAuth2 Package Documentation](https://pkg.go.dev/golang.org/x/oauth2)
- [Google OAuth2 Documentation](https://developers.google.com/identity/protocols/oauth2)
