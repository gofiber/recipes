package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testOauth2/config"
	"testOauth2/models"
	"testOauth2/router"
	"time"

	"github.com/antigloss/go/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/segmentio/encoding/json"
)

// Assets contains embedded frontend ressources
//go:embed www/*
var Assets embed.FS

func main() {
	var app *fiber.App
	var err error

	models.ClientID = config.Config("CLIENT_ID")
	models.ClientSecret = config.Config("CLIENT_SECRET")
	models.MySessionStore = session.New(session.Config{
		CookieSecure: true,
	})

	// define system log
	models.SYSLOG, err = logger.New(&logger.Config{
		LogDir:            "./logs",
		LogFileMaxSize:    200,
		LogFileMaxNum:     500,
		LogFileNumToDel:   50,
		LogFilenamePrefix: "SYS",
		LogLevel:          logger.LogLevelTrace,
		LogDest:           logger.LogDestFile, //| logger.LogDestConsole,
		Flag:              logger.ControlFlagLogDate | logger.ControlFlagLogLineNum,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting app ...")
	models.SYSLOG.Info("Starting app ...")

	// load only the contents of the subfolder www
	if subFS, err := fs.Sub(Assets, "www"); err != nil {
		subFS = nil
		models.SYSLOG.Fatal(err)
	} else {
		engine := html.NewFileSystem(http.FS(subFS), ".html")
		// engine.Reload(true)       // if the templates need constant reparsing
		engine.Debug(true)        // make the engine declare parsed templates
		engine.Delims("{{", "}}") // define delimiters to use in the templates

		// instantiate the application
		app = fiber.New(fiber.Config{
			Prefork:          false,            // run in a single thread
			ServerHeader:     "OAUTH2 tester",  // name the server
			DisableKeepalive: false,            // <-- must keep alive to have web sockets working
			JSONEncoder:      json.Marshal,     // use a better JSON library
			ReadTimeout:      time.Second * 20, // set a timeout to be able to shutdown the application
			Views:            engine,           // declare templating engine
		})

		// prepare routes
		router.SetupRoutes(app)
	}

	fmt.Println("Running app ...")
	models.SYSLOG.Info("Running app ...")

	// run server in a separate goroutine
	go func() {
		if err := app.Listen(":8080"); err != nil {
			models.SYSLOG.Fatal(err)
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
