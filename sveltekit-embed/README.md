---
title: Sveltekit Embed
keywords: [sveltekit, tailwindcss, embed]
---

# Fiber Sveltekit Embed App

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/sveltekit-embed) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/sveltekit-embed)

![image](https://github.com/gofiber/recipes/assets/40540244/2aa084b8-9bbc-46f3-9759-930857429f05)

This application is a full-stack project built using Sveltekit, Tailwind CSS, Fiber. It showcases the construction of a monolithic architecture for a full-stack application.

## Run the Project

To run the project, follow these steps:

1. Execute the following command to run all the necessary commands for building and running the application:

```bash
make all
```
2. Once the build process is complete, you can start the application by running:
```bash
./app
```


## Available Commands
The following commands are available to manage the project:


| Command | Description |
| --- | --- |
| `info` | Info command. Displays the available commands and the purpose of the application. |
| `go-build` | Builds the Golang project and creates an `app` file. |
| `svelte-build` | Builds the SvelteKit project. It first installs the dependencies and then performs the project build. |
| `all` | Runs both `svelte-build` and `go-build` commands sequentially. |

## Usage

To use this application, run the following command:

```bash
make <command>
```


API Routes
----------

The Go Fiber application provides the following API routes:

| Route | Description |
| --- | --- |
| `/*` | Serves static files from the specified directory (`template.Dist()`). If a file is not found, it serves `index.html`. |

Go Dependencies
---------------

-   **Go Modules:** Go's built-in package manager used to manage dependencies for Go projects.
-   **Fiber:** A fast and minimalist web framework for Golang.

Npm Dependencies
----------------

-   **SvelteKit:** A JavaScript framework used to build modern web applications.
-   **Tailwind CSS:** A fast and customizable CSS styling library. Can be used in SvelteKit projects.
-   **Skeleton UI:** This is a fully featured UI Toolkit for building reactive interfaces quickly using Svelte and Tailwind.

----------------

Author: [@ugurkorkmaz](https://github.com/ugurkorkmaz)

