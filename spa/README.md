---
title: Single Page Application (SPA)
keywords: [spa, static, client-side routing]
description: Serving a Single Page Application (SPA) with Fiber, including the index.html fallback for client-side routes.
---

# Single Page Application (SPA)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/spa) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/spa)

This project demonstrates how to serve a Single Page Application (SPA) from a Go backend built with the Fiber framework. Fiber serves the files of the built frontend from the `dist` directory under `/web`. Every other path below `/web`, such as the client-side route `/web/about`, gets `dist/index.html`, so reloading the page or opening a link directly still loads the app and its router takes over.

The `dist` directory contains a small hand-written SPA (a single `index.html` with a tiny client-side router and no build step), so the recipe runs as it is. Replace it with the output of the build of your own frontend, for example React, Vue or Svelte.

## Prerequisites

Ensure you have the following installed:

- Golang

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/spa
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

1. Start the Fiber application from this directory, because it reads `dist` relative to the working directory:
    ```sh
    go run main.go
    ```

2. Open `http://localhost:3000/web`. The links in the page change the URL without loading a new page. Reload the page at `http://localhost:3000/web/about`, or open that address directly, and it still loads because Fiber answers with `index.html`.

## Using Your Own Frontend

Build your frontend into `dist` and make `/web/` its base path, so the asset URLs in the built `index.html` start with `/web/`. With [Vite](https://vite.dev), for example, set the base path in `vite.config.js` and run the build, which writes to `dist` by default:

```js
// vite.config.js
export default {
  base: "/web/",
};
```

```sh
npm run build
```

Other bundlers have a similar option, for example `--public-url /web/` for Parcel. Your router needs the same base path, for example `basename="/web"` for React Router.

## Example

Here is the `main.go` file that serves the SPA with Fiber v3 using the `static` middleware:

```go
// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://docs.gofiber.io
// 📝 Github Repository: https://github.com/gofiber/fiber

package main

import (
	"log"

	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Serve Single Page Application on "/web".
	// A path that is not a file in dist, such as the client-side route /web/about,
	// gets index.html so client-side routing works. The static middleware has set
	// the status to 404 by then, so reset it to 200.
	app.Get("/web*", static.New("dist", static.Config{
		NotFoundHandler: func(ctx fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).SendFile("./dist/index.html")
		},
	}))

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
```

The `static` middleware sets the status to `404` before it calls `NotFoundHandler`, so the handler resets it to `200`. Without that, the browser still shows the page, but crawlers, proxies and monitoring see an error for every client-side route.

A missing file such as `/web/missing.js` gets `index.html` with status `200` as well, as on most SPA hosts. To answer those with a real `404`, return `fiber.ErrNotFound` from the handler when `filepath.Ext(ctx.Path())` is not empty. Client-side routes that contain a dot would then get a `404` too.

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Fiber Static Middleware](https://docs.gofiber.io/middleware/static)
