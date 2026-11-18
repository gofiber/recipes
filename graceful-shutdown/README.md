---
title: Graceful shutdown
keywords: [graceful, shutdown, os/signal, channel]
description: Graceful shutdown of applications.
---

# Graceful shutdown in Fiber

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/graceful-shutdown) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/graceful-shutdown)

```
$ go run main.go

    _______ __
   / ____(_) /_  ___  _____
  / /_  / / __ \/ _ \/ ___/
 / __/ / / /_/ /  __/ /
/_/   /_/_.___/\___/_/          v3.7.0
--------------------------------------------------
INFO Server started on:         http://127.0.0.1:3000 (bound on host 0.0.0.0 and port 3000)
INFO Total handlers:            2
INFO Prefork:                   Disabled
INFO PID:                       2199
INFO Total process count:       1

^CGracefully shutting down...
Running cleanup tasks...
Fiber was successfully shutdown.
```

This shows how to implement a graceful shutdown with Fiber and the `os/signal` package.

## Explanation

This example relies on the use of channels, a data type in Go that allows you to send and receive data to/from specific places in an application (read more about them [here](https://tour.golang.org/concurrency/2)).

A channel is created, and registered with `signal.Notify` so that when the program receives an interrupt (for example, when `CTRL+C` is pressed), a notification is sent to the channel. Once this is received, `app.ShutdownWithTimeout` is called to stop accepting new connections and wait up to 5 seconds for the active ones to finish, which also makes `app.Listen` return. After this point, cleanup functions can be run and the program eventually quits.
