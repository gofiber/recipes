---
title: Firebase Authentication
keywords: [firebase, authentication, middleware]
---

# Go Fiber Firebase Authentication Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/firebase-auth) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/firebase-auth)

This example use [gofiber-firebaseauth middleware](https://github.com/sacsand/gofiber-firebaseauth) to authenticate the endpoints. Find the documentation for middleware here for more configurations options [docs](https://github.com/sacsand/gofiber-firebaseauth)

## Setting Up

* Clone the repo and set your firebase credentials in your .env file
 Need Configured Firebase Authentication App and Google Service Account Credential (JSON file contain credential). You can get all these config from Firebase Console.

```
SERVICE_ACCOUNT_JSON = "path to service account credential json"
```

## Start

```
go build
go run main
```
