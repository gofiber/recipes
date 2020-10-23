package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("processing...")
		time.Sleep(5 * time.Second)
		return c.SendString("Hello world!")
	})

	// server listening
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt is sent, notify the channel

	// when an interrupt is received
	// waiting until done then shutdown
	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

}
