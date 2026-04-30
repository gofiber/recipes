package main

import (
	"crypto"
	"crypto/tls"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/pkcs12"
)

func initFiberApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("This page is being served over TLS using a PKCS12 store type!")
	})

	return app
}

func initTLSConfig(path string, password string) (*tls.Certificate, error) {
	pkcs12Data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	key, cert, err := pkcs12.Decode(pkcs12Data, password)
	if err != nil {
		return nil, err
	}

	tlsCert := tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key.(crypto.PrivateKey),
		Leaf:        cert,
	}

	return &tlsCert, nil
}

func main() {
	path := "./security/server.p12"
	password := os.Getenv("PKCS12_PASSWORD")
	if password == "" {
		password = "changeit"
	}

	tlsCert, err := initTLSConfig(path, password)
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{*tlsCert}}

	app := initFiberApp()
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listener(ln))
}
