// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber
// ❤️ Special thanks to @mthli for creating this example
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
	
	// Let’s Encrypt has rate limits: https://letsencrypt.org/docs/rate-limits/
	// It's recommended to use it's staging environment to test the code:
	// https://letsencrypt.org/docs/staging-environment/

	// Certificate manager
	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		// Replace with your domain
		HostPolicy: autocert.HostWhitelist("example.com"),
		// Folder to store the certificates
		Cache: autocert.DirCache("./certs"),
	}
	
	// TLS Config
	tls := &tls.Config{
		// Get Certificate from Let's Encrypt
		GetCertificate: m.GetCertificate,
		// By default NextProtos contains the "h2"
		// This has to be removed since Fasthttp does not support HTTP/2
		// Or it will cause a flood of PRI method logs
		// http://webconcepts.info/concepts/http-method/PRI
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}

	// Listen on a secure port
	app.Listen(443, tls)
}
