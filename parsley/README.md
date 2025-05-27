---
title: Parsley
keywords: [parsley, dependency injection, di, inversion of control, ioc]
description: Using Parsley for dependency injection in an application.
---

# Fiber with Dependency Injection (via Parsley)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/parsley) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/parsley)

This example demonstrates integrating the [Parsley dependency injection framework](https://github.com/matzefriedrich/parsley) into a GoFiber web application. The goal is to showcase how dependency injection can create a clean, maintainable, and modular structure in your GoFiber projects.


## Prerequisites

* Go 1.23+


## Overview

In this example, we use [Parsley](https://github.com/matzefriedrich/parsley) to:

* **Bootstrap the application:** Set up and configure the Fiber app using Parsley’s DI container.
* **Register dependencies:** Define and register services and route handlers with the DI container.
* **Resolve dependencies:** Automatically resolve and inject them where needed.


### Key features

* **Modular configuration:** Services are registered in modules, allowing for a clean separation of concerns.
* **Automatic dependency injection:** Constructor-based dependency injection wires services together.
* **Simplified route management:** Route handlers are registered and managed via the DI container, making it easy to extend and maintain.


## How it works

* The `main` function bootstraps the application using Parsley’s `RunParsleyApplication` function.
* Modules define how services (such as the Fiber app and route handlers) are registered and configured.
* Route handlers are implemented as services that receive their dependencies (like the `Greeter` service) via constructor injection. The `Greeter` service is a simple example of how services can be injected and used within route handlers to handle requests.


## The recipe - step by step

This guide demonstrates integrating the Parsley dependency injection framework with the GoFiber web framework. You can either clone the GoFiber recipes repository and navigate to the **parsley** example, or replicate each module while following the article:

```sh
git clone https://github.com/gofiber/recipes.git
cd recipes/parsley
```

The main entry point of the application is in the `cmd/main.go`.

```go
package main

import (
    "context"

    "github.com/gofiber/recipes/parsley-app/internal"
    "github.com/gofiber/recipes/parsley-app/internal/modules"

    "github.com/matzefriedrich/parsley/pkg/bootstrap"
)

func main() {

    ctx := context.Background()

    // Runs a Fiber instance as a Parsley-enabled app
    bootstrap.RunParsleyApplication(ctx, internal.NewApp,
        modules.ConfigureFiber,
        modules.ConfigureGreeter)
}
```

In this file, the `RunParsleyApplication` function bootstraps the application. It initializes the Parsley application context and configures the GoFiber server with the necessary services and route handlers. Parsley's `bootstrap` package is generic and could also be used with other web application frameworks; the glue is the `NewApp` method, representing a constructor function that must return a `bootstrap.Application` instance.

The last parameter of the `RunParsleyApplication` function is an ellipsis parameter accepting `ModuleFunc` values representing service registration functions, which are invoked before calling the constructor function for `bootstrap.Application`. Here, the `ConfigureFiber` and `ConfigureGreeter` functions are specified; those are defined by the `modules` package.


### Configure and register the Fiber instance

The `ConfigureFiber` function sets up the Fiber application and registers it as a singleton service within the Parsley framework:

```go
package modules

import (
    "github.com/gofiber/fiber/v2"
    "github.com/matzefriedrich/parsley/pkg/registration"
    "github.com/matzefriedrich/parsley/pkg/types"
)

var _ types.ModuleFunc = ConfigureFiber

func ConfigureFiber(registry types.ServiceRegistry) error {
    registration.RegisterInstance(registry, fiber.Config{
        AppName:   "parsley-app-recipe",
        Immutable: true,
    })

    registry.Register(newFiber, types.LifetimeSingleton)
    registry.RegisterModule(RegisterRouteHandlers)

    return nil
}

func newFiber(config fiber.Config) *fiber.App {
    return fiber.New(config)
}

```

This configuration ensures that the Fiber instance is initialized and available for dependency injection.


### Define and register the application service(s)

The `Greeter` service generates greeting messages based on input parameters. In the recipe example application, this service is a dependency required by the handler of the `say-hello` route.

```go
package services

import "fmt"

type Greeter interface {
    SayHello(name string, polite bool) string
}

type greeter struct{}

func (g *greeter) SayHello(name string, polite bool) string {
    if polite {
        return fmt.Sprintf("Good day, %s!\n", name)
    }
    return fmt.Sprintf("Hi, %s\n", name)
}

func NewGreeter() Greeter {
    return &greeter{}
}
```

The `Greeter` service is registered by the `ConfigureGreeter` service registration module:

```go
package modules

import (
    "github.com/gofiber/recipes/parsley-app/internal/services"

    "github.com/matzefriedrich/parsley/pkg/types"
)

func ConfigureGreeter(registry types.ServiceRegistry) error {
    registry.Register(services.NewGreeterFactory, types.LifetimeTransient)
    return nil
}
```

This setup allows the `Greeter` service to be injected wherever needed within the application.


### Implement and register route handlers

Route handlers in this example are services that implement the `RouteHandler` interface, allowing them to register routes with the Fiber application.

```go
package route_handlers

import (
    "strconv"

    "github.com/gofiber/recipes/parsley-app/internal/services"

    "github.com/gofiber/fiber/v2"
)

type greeterRouteHandler struct {
    greeter services.Greeter
}

const defaultPoliteFlag = "true"

func (h *greeterRouteHandler) Register(app *fiber.App) {
    app.Get("/say-hello", h.HandleSayHelloRequest)
}

func (h *greeterRouteHandler) HandleSayHelloRequest(ctx *fiber.Ctx) error {

    name := ctx.Query("name")

    politeFlag := ctx.Query("polite", defaultPoliteFlag)
    polite, _ := strconv.ParseBool(politeFlag)

    msg := h.greeter.SayHello(name, polite)
    return ctx.Status(fiber.StatusOK).Send([]byte(msg))
}

var _ RouteHandler = &greeterRouteHandler{}

func NewGreeterRouteHandler(greeter services.Greeter) RouteHandler {
    return &greeterRouteHandler{
        greeter: greeter,
    }
}
```

This handler responds to GET requests at `/say-hello` with a greeting message, utilizing the `Greeter` service injected via the constructor function.


## Run the application

To start the application, execute:

```sh
go run ./cmd/main.go
```

Once running, you can test the `say-hello` endpoint via the browser, or from the terminal using `curl`. For this recipe, the default listening port is `5502`:

```sh
curl http://localhost:5502/say-hello?name=YourName&polite=true
```
