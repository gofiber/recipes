---
title: Template Asset Bundling
keywords: [template, tailwindcss, parcel]
description: Setting up a Go application with template rendering and asset bundling.
---

# Template Asset Bundling

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/template-asset-bundling) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/template-asset-bundling)

This example demonstrates how to integrate asset bundling into a Go web application using [Fiber](https://github.com/gofiber/fiber), [gofiber/template](https://github.com/gofiber/template) for HTML template rendering, [Tailwind CSS](https://tailwindcss.com) for utility-first styling, and [Parcel](https://parceljs.org) as a zero-configuration asset bundler. Parcel processes and hashes the CSS assets, which are then served as static files by Fiber.

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) 1.21+
- [Node.js](https://nodejs.org/) 18+ and npm

## Project Structure

```
template-asset-bundling/
├── app.go               # Fiber application entry point
├── handlers/
│   └── handlers.go      # Route handlers (Home, About, NotFound)
├── views/
│   ├── layouts/
│   │   └── main.html    # Base layout template
│   ├── partials/        # Reusable template partials
│   ├── index.html       # Home page template
│   ├── about.html       # About page template
│   └── 404.html         # Not found template
├── assets/
│   └── app.css          # Tailwind CSS source (input)
├── public/
│   └── assets/          # Compiled assets output (git-ignored)
├── package.json         # Node dependencies and npm scripts
├── tailwind.config.js   # Tailwind CSS configuration
└── .postcssrc           # PostCSS configuration for Parcel
```

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/template-asset-bundling
    ```

2. Install Node.js dependencies:
    ```sh
    npm install
    ```

3. Install Go dependencies:
    ```sh
    go mod download
    ```

## Running

### Development

Run the asset watcher and the Go server in separate terminals:

```sh
# Terminal 1 — watch and rebuild assets on change
npm run dev

# Terminal 2 — start the Fiber server (template hot-reload enabled)
go run app.go
```

Open [http://localhost:3000](http://localhost:3000) in your browser.

### Production

Build optimized assets first, then run the server with `APP_ENV=production` to disable template hot-reloading:

```sh
npm run build
APP_ENV=production go run app.go
```

## How It Works

- `npm run dev` runs Parcel in watch mode, compiling `assets/app.css` (Tailwind source) into hashed output files under `public/assets/`.
- The `getCssAsset` template function walks `public/assets/` at render time to find the correct hashed filename and injects the `<link>` tag automatically.
- In development (`APP_ENV` is not `production`), `engine.Reload(true)` re-parses templates on every request so changes are reflected without restarting the server.
