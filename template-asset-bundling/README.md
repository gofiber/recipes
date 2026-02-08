---
title: Template Asset Bundling
keywords: [template, tailwindcss, parcel]
description: Setting up a Go application with template rendering and asset bundling.
---

# Template Asset Bundling

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/template-asset-bundling) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/template-asset-bundling)

This is a quick example of how to do asset bundling using [gofiber/template](https://github.com/gofiber/template), [Tailwind CSS](https://tailwindcss.com), and [Parcel](https://parceljs.org).

## Prerequisites

Ensure you have the following installed:

- Golang
- Node.js
- npm

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/template-asset-bundling
    ```

2. Install dependencies:
    ```sh
    npm install
    ```

## Usage

### Building Assets

1. Build the assets:
    ```sh
    npm run build
    ```

2. Watch assets for changes:
    ```sh
    npm run dev
    ```

### Running the Application

1. Start the Fiber application:
    ```sh
    go run main.go
    ```
