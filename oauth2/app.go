package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"oauth2/config"
	"oauth2/models"
	"oauth2/router"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/template/html/v3"
	"github.com/joho/godotenv"
	"github.com/segmentio/encoding/json"
)

// Assets contains embedded frontend resources
//
//go:embed www/*
var Assets embed.FS

func main() {
	var app *fiber.App

	// Load .env once at startup
	if err := godotenv.Load(".env"); err != nil {
		slog.Warn("could not load .env file", "err", err)
	}

	models.ClientID = config.Config("CLIENT_ID")
	models.ClientSecret = config.Config("CLIENT_SECRET")
	models.MySessionStore = session.NewStore(session.Config{
		CookieSecure: true,
	})

	fmt.Println("Starting app ...")
	models.SYSLOG.Info("Starting app ...")

	// load only the contents of the subfolder www
	subFS, err := fs.Sub(Assets, "www")
	if err != nil {
		slog.Error("could not create sub FS", "err", err)
		os.Exit(1)
	}

	engine := html.NewFileSystem(http.FS(subFS), ".html")
	// engine.Reload(true)       // if the templates need constant reparsing
	engine.Debug(true)        // make the engine declare parsed templates
	engine.Delims("{{", "}}") // define delimiters to use in the templates

	// instantiate the application
	app = fiber.New(fiber.Config{
		ServerHeader:     "OAUTH2 tester",  // name the server
		DisableKeepalive: false,            // <-- must keep alive to have web sockets working
		JSONEncoder:      json.Marshal,     // use a better JSON library
		ReadTimeout:      time.Second * 20, // set a timeout to be able to shutdown the application
		Views:            engine,           // declare templating engine
	})

	// prepare routes
	router.SetupRoutes(app)

	fmt.Println("Running app ...")
	models.SYSLOG.Info("Running app ...")

	// run server in a separate goroutine
	go func() {
		if err := app.Listen(":8080", fiber.ListenConfig{EnablePrefork: false}); err != nil {
			slog.Error("server error", "err", err)
			os.Exit(1)
		}
	}()

	// Create channel to send a notification when a signal is being received
	c := make(chan os.Signal, 1)

	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	models.SYSLOG.Info("Gracefully shutting down ...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	models.SYSLOG.Info("Running cleanup tasks...")

	// Your other cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	// etc.
	models.MySessionStore.Storage.Close()

	models.SYSLOG.Info("DONE.")
	fmt.Println("DONE.")
}
