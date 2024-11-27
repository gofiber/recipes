---
title: Parsley
keywords: [parsley, dependency injection, di, inversion of control, ioc]
description: Using Parsley for dependency injection in an application.
---

# Fiber with Dependency Injection (via Parsley)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/parsley) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/parsley)

This example demonstrates integrating the [Parsley dependency injection framework](https://github.com/matzefriedrich/parsley) into a GoFiber web application. The goal is to showcase how dependency injection can create a clean, maintainable, and modular structure in your GoFiber projects.

## Overview

In this example, we use Parsley to:

* **Bootstrap the Application:** Set up and configure the Fiber app using Parsley’s DI container.
* **Register Dependencies:** Define and register services and route handlers with the DI container.
* **Resolve Dependencies:** Automatically resolve and inject them where needed.

## Key Features

* **Modular Configuration:** Services are registered in modules, allowing for a clean separation of concerns.
* **Automatic Dependency Injection:** Constructor-based dependency injection wire services together.
* **Simplified Route Management:** Route handlers are registered and managed via the DI container, making it easy to extend and maintain.

## How It Works

* The `main` function bootstraps the application using Parsley’s `RunParsleyApplication` function.
* Modules define how services (such as the Fiber app and route handlers) are registered and configured.
* Route handlers are implemented as services that receive their dependencies (like the `Greeter` service) via constructor injection.
The `Greeter` service is a simple example of how services can be injected and used within route handlers to handle HTTP requests.

## Running the Example

To run this example:

* Clone the repository and navigate to the example directory.
* Run `go run main.go` to start the application.
* Access the application by navigating to `http://localhost:5502/say-hello?name=YourName`. This will return a greeting message, demonstrating the integration of Parsley with GoFiber.
