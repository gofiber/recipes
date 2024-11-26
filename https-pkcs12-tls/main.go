package main

import (
	"crypto"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/pkcs12"
)

func initFiberApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
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
	password := "changeit"

	tlsCert, error := initTLSConfig(path, password)

	if error != nil {
		fmt.Println("Unable to initialize TLS configuration object. Check your configuration and try again. Program will STOP.")
	} else {
		config := &tls.Config{Certificates: []tls.Certificate{*tlsCert}}

		app := initFiberApp()
		ln, err := tls.Listen("tcp", ":443", config)
		if err != nil {
			panic(err)
		}

		log.Fatal(app.Listener(ln))
	}
}
