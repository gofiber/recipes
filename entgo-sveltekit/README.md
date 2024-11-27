---
title: Entgo Sveltekit
keywords: [ent, sveltekit, tailwindcss, sqlite, rest]
description: A full-stack Todo application built using Sveltekit, Tailwind CSS, Entgo, and SQLite.
---

# Todo Application

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/entgo-sveltekit) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/entgo-sveltekit)

![image](https://github.com/ugurkorkmaz/gofiber-recipes/assets/40540244/08c6ee52-724a-4cf4-8352-9cf6f5b007ef)

This Todo application is a full-stack project built using Sveltekit, Tailwind CSS, Fiber, Entgo, and SQLite. It showcases the construction of a monolithic architecture for a full-stack application.

## Run the Project

To run the project, follow these steps:

1. Execute the following command to run all the necessary commands for building and running the application:

```bash
go run ./bin all
```
2. Once the build process is complete, you can start the application by running:
```bash
./app
```


## Available Commands
The following commands are available to manage the project:


| Command | Description |
| --- | --- |
| `go-run` | Run the Golang project. |
| `go-build` | Build the Golang project. |
| `go-test` | Run tests for the Golang project. |
| `svelte-run` | Run the SvelteKit project. |
| `svelte-build` | Build the SvelteKit project. |
| `generate-ent` | Generate entity files. |
| `all` | Run all commands (`generate-ent`, `svelte-build`, `go-test`, `go-build`). |

## Usage

To use this application, run the following command:

```bash
go run ./bin <command>
```


API Routes
----------

The Go Fiber application provides the following API routes:

| Method | Endpoint | Handler Function | Description |
| --- | --- | --- | --- |
| GET | /api/v1/todo/list | todoHandler.GetAllTodos | Get a list of all todos |
| GET | /api/v1/todo/get/:id | todoHandler.GetTodoByID | Get a specific todo by its ID |
| POST | /api/v1/todo/create | todoHandler.CreateTodo | Create a new todo |
| PUT | /api/v1/todo/update/:id | todoHandler.UpdateTodoByID | Update an existing todo by its ID |
| DELETE | /api/v1/todo/delete/:id | todoHandler.DeleteTodoByID | Delete a todo by its ID |

Go Dependencies
---------------

-   **Go Modules:** Go's built-in package manager used to manage dependencies for Go projects.
-   **Entgo:** A Golang Object Relational Mapping (ORM) tool used to define and generate database schemas.
-   **Fiber:** A fast and minimalist web framework for Golang.
-   **Sqlite:** A small, lightweight, embedded SQL database engine.

Npm Dependencies
----------------

-   **SvelteKit:** A JavaScript framework used to build modern web applications.
-   **Tailwind CSS:** A fast and customizable CSS styling library. Can be used in SvelteKit projects.

----------------

Author: [@ugurkorkmaz](https://github.com/ugurkorkmaz)

