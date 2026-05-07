package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"sync"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

// appPort is the port that the server will listen on
const appPort = "3000"

// index is the HTML template that will be served to the client on the index page (`/`)
const index = `<!DOCTYPE html>
<html>
<body>

<h1>SSE Messages</h1>
<div id="result"></div>

<script>
if(typeof(EventSource) !== "undefined") {
  var source = new EventSource("/sse");
  source.onmessage = function(event) {
    document.getElementById("result").innerHTML += event.data + "<br>";
  };
} else {
  document.getElementById("result").innerHTML = "Sorry, your browser does not support server-sent events...";
}
</script>

</body>
</html>
`

// indexTmpl is parsed once at startup to avoid per-request overhead.
var indexTmpl = template.Must(template.New("index").Parse(index))

var (
	sseMessageQueue []string
	mu              sync.Mutex
)

func main() {
	// Fiber instance
	app := fiber.New()

	// CORS for external resources
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Cache-Control"},
	}))

	app.Get("/", func(c fiber.Ctx) error {
		c.Response().Header.SetContentType(fiber.MIMETextHTMLCharsetUTF8)

		buf := new(bytes.Buffer)
		if err := indexTmpl.Execute(buf, nil); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).Send(buf.Bytes())
	})

	app.Get("/sse", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		return c.Status(fiber.StatusOK).SendStreamWriter(func(w *bufio.Writer) {
			var i int
			for {
				// Check for client disconnect before doing any work.
				select {
				case <-c.Context().Done():
					return
				default:
				}

				i++

				var msg string

				// If there are messages that have been sent to the `/publish` endpoint
				// then use these first, otherwise just send the current time.
				mu.Lock()
				if len(sseMessageQueue) > 0 {
					msg = fmt.Sprintf("%d - message recieved: %s", i, sseMessageQueue[0])
					sseMessageQueue = sseMessageQueue[1:]
				} else {
					msg = fmt.Sprintf("%d - the time is %v", i, time.Now())
				}
				mu.Unlock()

				fmt.Fprintf(w, "data: Message: %s\n\n", msg)

				if err := w.Flush(); err != nil {
					// Refreshing page in web browser will establish a new
					// SSE connection, but only (the last) one is alive, so
					// dead connections must be closed here.
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
					return
				}
				time.Sleep(2 * time.Second)
			}
		})
	})

	// Publish endpoint adds messages to the queue that will be sent to the client
	// via the `/sse` endpoint in FIFO order. If there are no messages in the queue
	// then the current time will be sent to the client instead.
	app.Put("/publish", func(c fiber.Ctx) error {
		type Message struct {
			Message string `json:"message"`
		}

		payload := new(Message)

		if err := c.Bind().Body(payload); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		mu.Lock()
		sseMessageQueue = append(sseMessageQueue, payload.Message)
		mu.Unlock()

		return c.SendString("Message added to queue\n")
	})

	// Start server
	log.Fatal(app.Listen(":" + appPort))
}

// fiber:context-methods migrated
