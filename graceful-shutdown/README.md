---
title: Graceful shutdown
keywords: [graceful, shutdown, os/signal, channel]
---

# Graceful shutdown in Fiber

```
fiberRecipes/graceful-shutdown on graceful-shutdown (f0834df) [?] via 🐹 v1.15.2 took 4s
❯ go run graceful-shutdown

 ┌───────────────────────────────────────────────────┐
 │                    Fiber v2.1.0                   │
 │               http://127.0.0.1:3000               │
 │                                                   │
 │ Handlers ............. 2  Threads ............. 8 │
 │ Prefork ....... Disabled  PID .............. 2540 │
 └───────────────────────────────────────────────────┘

^CGracefully shutting down...
Running cleanup tasks...
```

This shows how to implement a graceful shutdown with Fiber and the `os/signal` package.

## Explanation

This example relies on the use of channels, a data type in Go that allows you to send and receive data to/from specific places in an application (read more about them [here](https://tour.golang.org/concurrency/2)).

A channel is created, and registered with `signal.Notify` so that when the program receives an interrupt (for example, when `CTRL+C` is pressed), a notification is sent to the channel. Once this is received, `app.Shutdown` is called to close all active connections and return from `app.Listen`. After this point, cleanup functions can be run and the program eventually quits.
