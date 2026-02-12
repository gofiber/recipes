import { writable } from 'svelte/store';

export const gofiber_code = writable(
	`package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    app.Get("/", func (c fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":3000"))
}
`
);

export const robustRouting = writable(
	`app.Get("/", func (c fiber.Ctx) error {
    return c.SendString("GET request")
})

app.Get("/:param", func (c fiber.Ctx) error {
    return c.SendString("param: " + c.Params("param"))
})

app.Post("/", func (c fiber.Ctx) error {
    return c.SendString("POST request")
})
`
);

export const serverStaticFiles = writable(
	`app.Static("/", "./public")

// => http://localhost:3000/hello.html
// => http://localhost:3000/js/jquery.js
// => http://localhost:3000/css/style.css

// serve from multiple directories
app.Static("/", "./files")`
);
