---
title: Colly Gorm
keywords: [colly, gorm, postgresql]
description: Web scraping with Colly and GORM.
---

# Simple Web Scraping Colly App with Fiber

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/colly-gorm) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/colly-gorm)

A Go application using [Fiber](https://gofiber.io), [Colly v2](https://go-colly.org/), and [GORM](https://gorm.io) to scrape websites and persist data in PostgreSQL.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) and Docker Compose

## How to Run

1. Clone the repository.
2. Navigate to the project directory: `cd colly-gorm`
3. Copy the example env file: `cp app/app.env.example app/app.env`
4. Start the stack: `docker compose up --build`

## Project Structure

```
colly-gorm/
├── app/
│   ├── app.env.example              # Environment variable template
│   ├── Dockerfile
│   ├── go.mod
│   ├── cmd/
│   │   └── api/
│   │       └── main.go              # App entry point, Fiber routes
│   └── internals/
│       ├── consts/
│       │   └── consts.go            # Config loading via Viper
│       └── services/
│           ├── database/
│           │   ├── database.go      # GORM connection
│           │   └── models.go        # Quote and Course models
│           └── scrapers/
│               ├── toscrape.go      # Quotes scraper
│               └── coursera_courses.go  # Coursera scraper
├── db/
│   └── create_db.sql                # DB initialization
└── docker-compose.yml
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/healthchecker` | Health check — returns service status |
| `GET` | `/scrape/quotes` | Triggers async scraping of [quotes.toscrape.com](http://quotes.toscrape.com) and stores results in PostgreSQL |
| `GET` | `/scrape/coursera` | Triggers async scraping of [coursera.org/browse](https://www.coursera.org/browse) and stores course data in PostgreSQL |

Scraping jobs run asynchronously; the endpoint returns immediately while scraping continues in the background.

## Database Models

**Quote**
- `author` — quote author
- `quote` — quote text

**Course**
- `title`, `description`, `creator`, `url`, `rating`

## Environment Variables

See `app/app.env.example`:

```env
POSTGRES_HOST=colly_db
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=colly
```

## What It Does

- Registers Colly HTML callbacks before visiting pages (correct callback order).
- Scrapes data from websites and stores it in a PostgreSQL database via GORM.
- Uses Fiber middleware (logger, CORS) applied globally before sub-app routing.
