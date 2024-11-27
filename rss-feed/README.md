---
title: RSS Feed
keywords: [rss, feed]
description: Generating an RSS feed.
---

# RSS Feed

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/rss-feed) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/rss-feed)

This project demonstrates how to create an RSS feed in a Go application using the Fiber framework.

## Prerequisites

Ensure you have the following installed:

- Golang
- [Fiber](https://github.com/gofiber/fiber) package

## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/gofiber/recipes.git
    cd recipes/rss-feed
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

1. Start the application:
    ```sh
    go run main.go
    ```

## Example

Here is an example of how to create an RSS feed in a Fiber application:

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gorilla/feeds"
    "time"
)

func main() {
    app := fiber.New()

    app.Get("/rss", func(c *fiber.Ctx) error {
        feed := &feeds.Feed{
            Title:       "Example RSS Feed",
            Link:        &feeds.Link{Href: "http://example.com/rss"},
            Description: "This is an example RSS feed",
            Author:      &feeds.Author{Name: "John Doe", Email: "john@example.com"},
            Created:     time.Now(),
        }

        feed.Items = []*feeds.Item{
            {
                Title:       "First Post",
                Link:        &feeds.Link{Href: "http://example.com/post/1"},
                Description: "This is the first post",
                Author:      &feeds.Author{Name: "John Doe", Email: "john@example.com"},
                Created:     time.Now(),
            },
            {
                Title:       "Second Post",
                Link:        &feeds.Link{Href: "http://example.com/post/2"},
                Description: "This is the second post",
                Author:      &feeds.Author{Name: "Jane Doe", Email: "jane@example.com"},
                Created:     time.Now(),
            },
        }

        rss, err := feed.ToRss()
        if err != nil {
            return err
        }

        c.Set("Content-Type", "application/rss+xml")
        return c.SendString(rss)
    })

    app.Listen(":3000")
}
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Gorilla Feeds Documentation](https://pkg.go.dev/github.com/gorilla/feeds)
