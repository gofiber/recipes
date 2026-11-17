---
title: Sveltekit Embed
keywords: [sveltekit, tailwindcss, embed]
description: A full-stack application built using Sveltekit and Tailwind CSS.
---

# Fiber Sveltekit Embed App

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/sveltekit-embed) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/sveltekit-embed)

![image](https://github.com/gofiber/recipes/assets/40540244/2aa084b8-9bbc-46f3-9759-930857429f05)

This application is a full-stack project built using Sveltekit, Tailwind CSS, Fiber. It showcases the construction of a monolithic architecture for a full-stack application.

## Prerequisites

- Go
- Node.js 22.17 or newer (see `engines` in `template/package.json`). The makefile runs pnpm through `npx`, so nothing else has to be installed.

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
| `install-template-dependencies` | Installs the dependencies of the SvelteKit project with the pinned pnpm version and the frozen lockfile. |
| `svelte-build` | Builds the SvelteKit project into `template/dist`. It first installs the dependencies and then performs the project build. |
| `all` | Runs both `svelte-build` and `go-build` commands sequentially. |
| `build` | Same as `all`. This is the target CI runs. |
| `clean` | Removes the `app` file and the build output in `template/dist`. |

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
| `/*` | Serves the SvelteKit build that is embedded in the binary (`template.Dist()`). If a path is not a file in the build, such as a client-side route, it serves `index.html` with status 200, so the SvelteKit router can handle it. A missing file such as `/_app/missing.js` gets `index.html` as well. |

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

