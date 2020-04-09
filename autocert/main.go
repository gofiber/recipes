package main

import (
	"crypto/tls"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hi TLS")
	})

	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"),
		Cache:      autocert.DirCache("./certs"),
	}

	tls := &tls.Config{
		GetCertificate: m.GetCertificate,
	}

	app.Listen(443, tls)
}
