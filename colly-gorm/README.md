---
title: Colly Gorm
keywords: [colly, gorm, postgresql]
---

# Simple Web Scraping Colly App with Fiber

This is a basic Go application using the Fiber framework to create scraping tasks in colly.

## How to Run

1. Clone the repository.
2. Navigate to the project directory.
3. Run `docker compose up --build`.
4. Visit `http://127.0.0.1:3000/api/healthchecker` in a web browser or use a tool like `curl` to test it.
5. Send `GET` request to `http://127.0.0.1:3000/scrape/coursera` to start scraping Coursera courses. And `http://127.0.0.1:3000/scrape/quotes` to scrape `quotes.toscrape.com`.


## What It Does

- Scrapes data from websites and stores in PostgreSQL database.
