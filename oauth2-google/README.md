# Fiber with Google OAuth2

### Implementation of Google OAuth2 with fiber, some packages used are mentioned below
Obtain OAuth credentials from https://console.developers.google.com/

## Endpoints 
- /api/ - redirects to login url

- /api/auth/google/callback - gives a callback to google and on success return's user's email

## Packages Used
- [Godotenv](https://github.com/joho/godotenv)
- [Fiber](https://github.com/gofiber/fiber)
- [OAuth2](https://github.com/golang/oauth2)