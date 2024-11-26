---
title: Svelte Netlify
keywords: [netlify, deploy, svelte]
---

# Deploying fiber on Netlify

[![Netlify Status](https://api.netlify.com/api/v1/badges/143c3c42-60f7-427a-b3fd-8ca3947a2d40/deploy-status)](https://app.netlify.com/sites/gofiber-svelte/deploys)

### Demo @ https://gofiber-svelte.netlify.app/

#### Based on the fiber-lambda API written by Fenny. Since the code hasn't been merged yet, I borrowed it into `adapter/adapter.go`

The app uses static pages under `public` directory. These are compiled using sveltejs and the complete template can be found [here](https://github.com/amalshaji/gofiber-sveltejs-netlify).


```toml
# netlify.toml

[build]
  command = "./build.sh"
  functions = "functions"
  publish = "public"

[build.environment]
  GO_IMPORT_PATH = "github.com/amalshaji/fiber-netlify"
  GO111MODULE = "on"

[[redirects]]
  from = "/api/*"
  to = "/.netlify/functions/gateway/:splat"
  status = 200
```

Deploying `net/http to Netlify` explains what these functions are doing. You can read it [here](https://blog.carlmjohnson.net/post/2020/how-to-host-golang-on-netlify-for-free/).

#### TL;DR
- build command builds the whole code to binary `cmd/gateway/gateway`
- we're building something called [netlify functions](https://functions.netlify.com/) (Please read)
- everything under public folder will be published(entrypoint: `index.html`)
- Netlify maps endpoints to `/.netlify/functions/gateway`, which is weird when you do requests, so we redirect it to `/api/*`
- status = 200 for server side redirects

#### Important
Netlify functions allows you to have up to 125,000 requests a month. This means you can have 2.89 requests per minute. Make sure you use `Cache` in you request handlers.
