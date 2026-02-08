---
title: Colly Gorm
keywords: [colly, gorm, postgresql]
description: Web scraping with Colly and GORM.
---

# Simple Web Scraping Colly App with Fiber

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/colly-gorm) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/colly-gorm)

This is a basic Go application using the Fiber framework to create scraping tasks in colly.

## How to Run

1. Clone the repository.
2. Navigate to the project directory.
3. Run `docker compose up --build`.
4. Visit `http://127.0.0.1:3000/api/healthchecker` in a web browser or use a tool like `curl` to test it.
5. Send `GET` request to `http://127.0.0.1:3000/scrape/coursera` to start scraping Coursera courses. And `http://127.0.0.1:3000/scrape/quotes` to scrape `quotes.toscrape.com`.


## What It Does

- Scrapes data from websites and stores in PostgreSQL database.
