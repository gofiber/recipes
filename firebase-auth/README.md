---
title: Firebase Authentication
keywords: [firebase, authentication, middleware]
description: Firebase authentication integration.
---

# Go Fiber Firebase Authentication Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/firebase-auth) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/firebase-auth)

This example demonstrates how to protect Fiber routes with Firebase Authentication by verifying Firebase ID tokens.

## Requirements

- A Firebase project with Authentication enabled
- A Google Service Account credential JSON file (download from Firebase Console → Project Settings → Service Accounts)

## Setting Up

Copy `example.env` to `.env` and set the path to your service account credential file:

```
GOOGLE_SERVICE_ACCOUNT=path/to/serviceAccountKey.json
```

## Start

```bash
go run main.go
```

## Endpoints

| Method | Path            | Auth required | Description                          |
|--------|-----------------|---------------|--------------------------------------|
| GET    | /salut          | No            | Public greeting (French)             |
| POST   | /ciao           | No            | Public greeting (Italian)            |
| GET    | /salanthe       | No            | Public greeting (Sinhalese)          |
| GET    | /api/hello      | Yes           | Protected greeting (English)         |
| GET    | /api/ayubowan   | Yes           | Protected greeting with user claims  |

## curl Examples

### Public endpoint

```bash
curl http://localhost:3001/salut
```

### Protected endpoint — obtain a Firebase ID token first, then:

```bash
curl -H "Authorization: Bearer <YOUR_FIREBASE_ID_TOKEN>" \
     http://localhost:3001/api/hello
```

```bash
curl -H "Authorization: Bearer <YOUR_FIREBASE_ID_TOKEN>" \
     http://localhost:3001/api/ayubowan
```
