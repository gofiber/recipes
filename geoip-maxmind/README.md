---
title: GeoIP + MaxMind
keywords: [geoip, maxmind, databases]
description: Geolocation with GeoIP and MaxMind databases.
---

# GeoIP (with MaxMind databases)

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/geoip-maxmind) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/geoip-maxmind)

This is an alternative method to resolve IP addresses to real-world location data using MaxMind GeoLite2 City databases.

## Prerequisites
Before you run this, you must first download a database from the MaxMind website - https://dev.maxmind.com/geoip/geoip2/geolite2/. To do this, you may need to register for a free account.

The database you need to download is the one with the edition ID `GeoLite2-City`. Place it in this folder and run

```
go run geoip-maxmind
```

## Usage
Make a request to `http://127.0.0.1:3000/geo/178.62.56.160`, for example. You can omit an IP address to use your current IP address, or replace to use another. If the IP address is invalid, a HTTP 400 is returned.

The response fields can be modified from the `ipLookup` struct, found in the `handlers/handlers.go` file.

### Example response

```json
{
  "City": {
    "GeoNameID": 2643743,
    "Names": {
      "de": "London",
      "en": "London",
      "es": "Londres",
      "fr": "Londres",
      "ja": "ロンドン",
      "pt-BR": "Londres",
      "ru": "Лондон",
      "zh-CN": "伦敦"
    }
  },
  "Country": {
    "IsoCode": "GB"
  },
  "Location": {
    "AccuracyRadius": 50
  }
}
```
