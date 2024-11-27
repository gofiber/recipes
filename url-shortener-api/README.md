---
title: URL Shortener
keywords: [url shortener, redis, api]
description: URL shortening service with a simple API.
---

# URL Shortener API

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/url-shortener-api) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/url-shortener-api)

This project provides a URL shortening service with a simple API.

## Tech Stack

- Golang
- Redis

## API Documentation

> API endpoint: `http://localhost:3000/api/v1/`

### API Payload

- `url` - Original URL
- `short` - Custom short URL (Optional)
- `expiry` - Time to expire: int (hours)

### API Response

- `url` - Original URL
- `short` - Custom short URL
- `expiry` - Time to expire: int (hours)
- `rate_limit` - Number of API calls remaining: int
- `rate_limit_reset` - Time to rate limit reset: int (minutes)

> API is rate limited to 10 calls every 30 minutes.
> These values can be changed in the `.env` file. Have fun.

## Setup

1. Start the containers:
    ```sh
    docker-compose up -d
    ```

2. Test the API:
   ![test.gif](test.gif)
