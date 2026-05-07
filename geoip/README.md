---
title: GeoIP
keywords: [geoip, ip-api, geolocation]
description: Geolocation using ip-api.com.
---

# GeoIP Example

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/geoip) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/geoip)

This recipe demonstrates how to build a GeoIP lookup service with [Fiber](https://github.com/gofiber/fiber). It proxies requests to the [ip-api.com](http://ip-api.com) JSON API and caches responses for 10 minutes using Fiber's built-in cache middleware.

> **Note:** This recipe depends on the free [ip-api.com](http://ip-api.com) service. The free tier is limited to **1000 requests per minute** from a single IP address. For higher traffic, consider a paid plan or a self-hosted alternative such as [geoip-maxmind](../geoip-maxmind/).

## Prerequisites

- Go 1.21+
- Internet access (ip-api.com is called at runtime — no local database required)

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/geoip
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

```sh
go run main.go
```

The server starts on port `3000` by default. Set the `PORT` environment variable to override:

```sh
PORT=8080 go run main.go
```

Open `http://localhost:3000` in a browser to use the web UI.

## Example

Look up geolocation data for an IP address via the `/geo` endpoint:

```sh
curl "http://localhost:3000/geo?ip=178.62.56.160"
```

Example response:

```json
{
  "status": "success",
  "country": "United Kingdom",
  "countryCode": "GB",
  "region": "ENG",
  "regionName": "England",
  "city": "London",
  "zip": "EC1A",
  "lat": 51.5085,
  "lon": -0.1257,
  "timezone": "Europe/London",
  "isp": "DigitalOcean, LLC",
  "org": "DigitalOcean, LLC",
  "as": "AS14061 DigitalOcean, LLC",
  "query": "178.62.56.160"
}
```

Omit the `ip` query parameter to look up the caller's own IP address:

```sh
curl "http://localhost:3000/geo"
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [ip-api.com Documentation](http://ip-api.com/docs)
- [ip-api.com Rate Limits](http://ip-api.com/docs/api:json#usage_limits)
