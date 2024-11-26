---
title: Firebase Functions
keywords: [firebase, functions, deployment, gcloud, cloud]
---
# Deploying GoFiber Application to Firebase Functions

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/firebase-functions) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/firebase-functions)

Welcome to this step-by-step guide on deploying a GoFiber application to Firebase Functions. If you’re looking to leverage the power of GoFiber, a fast and lightweight web framework for Go, and host your application on Firebase, you’re in the right place. In this tutorial, we’ll walk through the process of setting up your GoFiber app to run seamlessly on Firebase Functions.

## Prerequisites

1. Go installed on your machine.
2. Firebase CLI installed.
3. A Firebase project created.
4. Firestore and Cloud Functions enabled.

## Create a GoFiber App

Start by initializing your GoFiber application. Use the following commands in your terminal:

```bash
go mod init example.com/GofiberFirebaseBoilerplate
```

## Server Configuration

Create a server file `(src/server.go)` with a `CreateServer` function that sets up your GoFiber server.

```go
package src

import (
 "example.com/GofiberFirebaseBoilerplate/src/routes"

 "github.com/gofiber/fiber/v2"
)

func CreateServer() *fiber.App {
 version := "v1.0.0"

 app := fiber.New(fiber.Config{
  ServerHeader: "Gofiber Firebase Boilerplate",
  AppName:      "Gofiber Firebase Boilerplate " + version,
 })

 app.Get("/", func(c *fiber.Ctx) error {
  return c.SendString("Gofiber Firebase Boilerplate " + version)
 })

 routes.New().Setup(app)

 return app
}
```

## Routes Configuration

Now that your GoFiber application is initialized, let’s delve into setting up and configuring routes. This section is crucial for defining how your application handles incoming requests. Open the `src/routes/routes.go` file to manage your routes.

```go
package routes

import (
 "example.com/GofiberFirebaseBoilerplate/src/database"
 "example.com/GofiberFirebaseBoilerplate/src/repositories"

 "github.com/gofiber/fiber/v2"
)

type Routes struct {
 mainRepository *repositories.MainRepository
}

func New() *Routes {
 db := database.NewConnection()
 return &Routes{mainRepository: &repositories.MainRepository{DB: db}}
}

func (r *Routes) Setup(app *fiber.App) {
 app.Post("message", r.insertMessage)
}

func (r *Routes) insertMessage(c *fiber.Ctx) error {
 return c.SendString("ok")
}
```

## Database Configuration

Configure your Firestore database connection in the `src/database/database.go` file. Make sure to replace the placeholder credentials with your Firebase project's actual credentials.

```go
package database

import (
 "context"
 "encoding/json"
 "log"

 "cloud.google.com/go/firestore"
 firebase "firebase.google.com/go"

 "google.golang.org/api/option"
)

type Config struct {
 Host     string
 Port     string
 Password string
 User     string
 DBName   string
 SSLMode  string
}

func NewConnection() *firestore.Client {

 ctx := context.Background()

 sa := option.WithCredentialsJSON(credentials())
 app, err := firebase.NewApp(ctx, nil, sa)
 if err != nil {
  log.Fatalf("functions.init: NewApp %v\n", err)
 }

 db, err := app.Firestore(ctx)
 if err != nil {
  log.Fatalf("functions.init: Database init : %v\n", err)
 }

 return db
}

func credentials() []byte {
 // TODO: Replace with your Credentials
 data := map[string]interface{}{
  "type":                        "",
  "project_id":                  "",
  "private_key_id":              "",
  "private_key":                 "",
  "client_email":                "",
  "client_id":                   "",
  "auth_uri":                    "",
  "token_uri":                   "",
  "auth_provider_x509_cert_url": "",
  "client_x509_cert_url":        "",
  "universe_domain":             "",
 }

 bytes, err := json.Marshal(data)
 if err != nil {
  panic(err)
 }

 return bytes
}
```

## Repository Pattern

Implement the repository pattern in the `src/repositories/main.repository.go` file to interact with Firestore. This file includes an example of inserting a message into the database.

```go
package repositories

import (
 "context"

 "cloud.google.com/go/firestore"
 "example.com/GofiberFirebaseBoilerplate/src/models"
 "github.com/google/uuid"
)

type MainRepository struct {
 DB *firestore.Client
}

func (r *MainRepository) InsertMessage(body *models.MessageInputBody) error {
 id := uuid.New().String()
 _, err := r.DB.Collection("messages").Doc(id).Set(context.Background(), body)
 return err
}
```

## Model Definition

Define a message input model in src/models/message_input_body.go to structure the data you'll be working with.

```go
package models

type MessageInputBody struct {
 From    string `json:"from"`
 To      string `json:"to"`
 Message string `json:"message"`
}
```

## Functions for Cloud Integration

In `functions.go`, convert Google Cloud Function requests to Fiber and route them to your application. This file includes functions to facilitate the integration of Google Cloud Functions and GoFiber.

```go
package app

import (
 "bytes"
 "context"
 "fmt"
 "io"
 "log"
 "net"
 "net/http"
 "strings"

 "github.com/gofiber/fiber/v2"
 "github.com/valyala/fasthttp/fasthttputil"
)

// CloudFunctionRouteToFiber route cloud function http.Handler to *fiber.App
// Internally, google calls the function with the /execute base URL
func CloudFunctionRouteToFiber(fiberApp *fiber.App, w http.ResponseWriter, r *http.Request) error {
 return RouteToFiber(fiberApp, w, r, "/execute")
}

// RouteToFiber route http.Handler to *fiber.App
func RouteToFiber(fiberApp *fiber.App, w http.ResponseWriter, r *http.Request, rootURL ...string) error {
 ln := fasthttputil.NewInmemoryListener()
 defer ln.Close()

 // Copy request
 body, err := io.ReadAll(r.Body)
 if err != nil {
  return err
 }

 url := fmt.Sprintf("%s://%s%s", "http", "0.0.0.0", r.RequestURI)
 if len(rootURL) > 0 {
  url = strings.Replace(url, rootURL[0], "", -1)
 }

 proxyReq, err := http.NewRequest(r.Method, url, bytes.NewReader(body))

 if err != nil {
  return err
 }

 proxyReq.Header = r.Header

 // Create http client
 client := http.Client{
  Transport: &http.Transport{
   DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
    return ln.Dial()
   },
  },
 }

 // Serve request to internal HTTP client
 go func() {
  log.Fatal(fiberApp.Listener(ln))
 }()

 // Call internal Fiber API
 response, err := client.Do(proxyReq)
 if err != nil {
  return err
 }

 // Copy response and headers
 for k, values := range response.Header {
  for _, v := range values {
   w.Header().Set(k, v)
  }
 }
 w.WriteHeader(response.StatusCode)

 io.Copy(w, response.Body)
 response.Body.Close()

 return nil
}
```

## Main Application Entry

In `main.go`, initialize your GoFiber app and start the server. This file also includes an exported Cloud Function handler for deployment.

```go
package app

import (
 "fmt"
 "net/http"
 "strings"

 "example.com/GofiberFirebaseBoilerplate/src"
 "github.com/gofiber/fiber/v2"
)

var app *fiber.App

func init() {
 app = src.CreateServer()
}

// Start start Fiber app with normal interface
func Start(addr string) error {
 if -1 == strings.IndexByte(addr, ':') {
  addr = ":" + addr
 }

 return app.Listen(addr)
}

// MyCloudFunction Exported http.HandlerFunc to be deployed to as a Cloud Function
func MyCloudFunction(w http.ResponseWriter, r *http.Request) {
 err := CloudFunctionRouteToFiber(app, w, r)
 if err != nil {
  fmt.Fprintf(w, "err : %v", err)
  return
 }
}
```

## Development

For local development, utilize the `cmd/main.go` file. If you prefer hot reloading, the `.air.toml` configuration file is included for use Air.

## cmd/main.go

```go
package main

import (
 "log"
 "os"

 app "example.com/GofiberFirebaseBoilerplate"
)

func main() {

 port := "3001"
 if envPort := os.Getenv("PORT"); envPort != "" {
  port = envPort
 }

 if err := app.Start(port); err != nil {
  log.Fatalf("app.Start: %v\n", err)
 }
}
```

## .air.toml

```go
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ./cmd"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  post_cmd = []
  pre_cmd = []
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
```

## Deployment

Deploy your Cloud Function using the following commands, replacing `<YourProjectID>` with your Firebase project ID:

```bash
gcloud config set project <YourProjectID>
gcloud functions deploy MyCloudFunction --runtime go120 --trigger-http
```

## Conclusion

Congratulations! You’ve successfully configured and deployed a GoFiber application on Firebase Functions. This powerful combination allows you to build fast and efficient serverless applications. Experiment further with GoFiber features and Firebase integrations to unlock the full potential of your serverless architecture. Happy coding!

## Medium Post
https://medium.com/@kmltrk07/how-to-deploy-gofiber-app-to-firebase-functions-8d4d537a4464
