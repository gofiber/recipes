package main

import (
	"fmt"
	"github.com/cloudflare/tableflip"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const version = "v0.0.1"

func main() {
	upg, _ := tableflip.New(tableflip.Options{})
	defer upg.Stop()

	// By prefixing PID to log, easy to interrupt from another process.
	log.SetPrefix(fmt.Sprintf("[PID: %d]", os.Getpid()))

	// Listen for the process signal to trigger the tableflip upgrade.
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)
		for range sig {
			upg.Upgrade()
		}
	}()

	// Listen must be called before Ready
	ln, _ := upg.Listen("tcp", "localhost:8080")
	defer ln.Close()

	app := fiber.New()
	app.Get("/version", func(c *fiber.Ctx) error {
		log.Println(version)
		return c.SendString(version)
	})

	go app.Listener(ln)

	// tableflip ready
	if err := upg.Ready(); err != nil {
		panic(err)
	}

	<-upg.Exit()
}
