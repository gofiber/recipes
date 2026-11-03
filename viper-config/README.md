---
title: Viper Configuration
keywords: [viper, config, configuration, environment, yaml]
description: Manage layered application configuration in Fiber using Viper with YAML and environment variables.
---

# Viper Configuration

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/viper-config) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/viper-config)

This recipe demonstrates how to manage layered configurations in a [Fiber](https://github.com/gofiber/fiber) application using [Viper](https://github.com/spf13/viper).

## Features

- **Layered Defaults**: Pre-configured defaults for application settings.
- **YAML Configuration File**: Reads `config.yaml` automatically when present.
- **Environment Variable Overrides**: Environment variables prefixed with `APP_` override file and default values.
- **Type-safe Unmarshaling**: Maps configuration values directly into a Go struct with `mapstructure` tags.

## Prerequisites

- [Go](https://go.dev/) 1.26 or newer.

## Setup & Running

1. Navigate to the recipe directory:

```bash
cd viper-config
```

2. Run the application:

```bash
go run main.go
```

3. In a separate terminal, test the endpoints:

```bash
curl http://localhost:3000/
# Output: {"env":"staging","port":3000,"service":"Fiber Viper (from config.yaml)","version":"1.0.0"}

curl http://localhost:3000/healthz
# Output: OK
```

`env` and `service` come from `config.yaml`, `port` and `version` from the defaults in `main.go`.

4. Test overriding settings with environment variables:

Stop the running server (`Ctrl+C`), then start with custom environment variables:

```bash
APP_PORT=8080 APP_ENV=production APP_SERVICE_NAME="Fiber Production" go run main.go
```

In a separate terminal:

```bash
curl http://localhost:8080/
# Output: {"env":"production","port":8080,"service":"Fiber Production","version":"1.0.0"}
```

Viper's `Unmarshal` only sees environment values for keys it already knows, which is why every key in `loadConfig` has a default.
