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

This recipe uses [Mustache templates](https://github.com/cbroglie/mustache) (via the `gofiber/template/mustache` engine) to render an [RSS 2.0](https://www.rssboard.org/rss-specification) feed. The template lives in `./xmls/example.xml`. A `channel` has a title, link and description, and each `item` is rendered by the `{{#Items}}` section.

Use `{{Name}}` (two braces) for text. It escapes `&`, `<`, `>` and quotes, so a title such as `Fiber recipes & news` still produces valid XML. `{{{Name}}}` (three braces) inserts the value unescaped and would break the feed.

**`xmls/example.xml`:**

```xml
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
    <channel>
        <title>{{Title}}</title>
        <link>{{Link}}</link>
        <description>{{Description}}</description>
        <language>{{Lang}}</language>
        {{#Items}}
        <item>
            <title>{{Title}}</title>
            <link>{{Link}}</link>
            <description>{{Description}}</description>
            <guid isPermaLink="true">{{Link}}</guid>
            <pubDate>{{PubDate}}</pubDate>
        </item>
        {{/Items}}
    </channel>
</rss>
```

**`main.go`:**

```go
package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/mustache/v3"
)

func main() {
	engineXML := mustache.New("./xmls", ".xml")
	if err := engineXML.Load(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/rss", func(c fiber.Ctx) error {
		// Set Content-Type to application/rss+xml
		c.Type("rss")

		published := time.Date(2024, time.January, 2, 15, 4, 5, 0, time.UTC)

		// Set rendered template to body. {{Name}} in the template escapes &, <, > and
		// quotes, so titles and descriptions cannot break the XML. ({{{Name}}} would not.)
		return engineXML.Render(c, "example", fiber.Map{
			"Lang":        "en",
			"Title":       "Fiber recipes & news",
			"Link":        "http://localhost:3000/",
			"Description": "An example RSS 2.0 feed rendered with Mustache",
			"Items": []fiber.Map{
				{
					"Title":       "Hello, RSS",
					"Link":        "http://localhost:3000/posts/1",
					"Description": "The first <post> in the feed",
					"PubDate":     published.Format(time.RFC1123Z),
				},
				{
					"Title":       "Second item",
					"Link":        "http://localhost:3000/posts/2",
					"Description": "Another entry",
					"PubDate":     published.Add(24 * time.Hour).Format(time.RFC1123Z),
				},
			},
		})
	})

	log.Fatal(app.Listen(":3000"))
}
```

### Testing with curl

```sh
curl http://localhost:3000/rss
```

## References

- [Fiber Documentation](https://docs.gofiber.io)
- [Fiber Template - Mustache](https://github.com/gofiber/template/tree/master/mustache)
- [cbroglie/mustache](https://pkg.go.dev/github.com/cbroglie/mustache)
