---
title: Tableflip Example
keywords: [tableflip, golang, graceful upgrade]
---

# Tableflip Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/tableflip) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/tableflip)

This example demonstrates how to use [tableflip](https://github.com/cloudflare/tableflip) for graceful upgrades in a Go application.

## What is Tableflip?

Tableflip is a library that allows you to update the running code and/or configuration of a network service without disrupting existing connections. It achieves this by starting a new process, transferring clients to it, and then exiting the old process.

### Goals of Tableflip

- No old code keeps running after a successful upgrade.
- The new process has a grace period for initialization.
- Crashing during initialization is acceptable.
- Only a single upgrade is ever run in parallel.
- Tableflip works on Linux and macOS.

## Steps

1. **Build v0.0.1 Demo:**

    ```bash
    go build -o demo main.go
    ```

2. **Run the Demo and Create a GET Request to `127.0.0.1:8080/version`:**

    ```bash
    [PID: 123] v0.0.1
    ```

3. **Prepare a New Version:**

- Change the `main.go` to update the version to "v0.0.2".
- Rebuild the demo:

    ```bash
    go build -o demo main.go
    ```

4. **Kill the Old Process:**

    ```bash
    kill -s HUP 123
    ```

5. **Create the Request to the Version API Again:**

    ```bash
    [PID: 123] v0.0.2
    ```

The client is completely immune to server upgrades and reboots, and our application updates gracefully!
