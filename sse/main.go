package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"
)

// appPort is the port that the server will listen on
const appPort = "7331"

// index is the HTML template that will be served to the client on the index page (`/`)
const index = `<!DOCTYPE html>
<html>
<body>

<h1>SSE Messages</h1>
<div id="result"></div>

<script>
if(typeof(EventSource) !== "undefined") {
  var source = new EventSource("http://127.0.0.1:{{.Port}}/sse");
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

func main() {
	// Fiber instance
	app := fiber.New()

	// CORS for external resources
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Cache-Control",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Response().Header.SetContentType(fiber.MIMETextHTMLCharsetUTF8)

		tpl, err := template.New("index").Parse(index)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		data := struct {
			Port string
		}{
			Port: appPort,
		}

		buf := new(bytes.Buffer)
		err = tpl.Execute(buf, data)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).Send(buf.Bytes())
	})

	app.Get("/sse", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")
			var i int
			for {
				i++
				msg := fmt.Sprintf("%d - the time is %v", i, time.Now())
				fmt.Fprintf(w, "data: Message: %s\n\n", msg)
				fmt.Println(msg)

				err := w.Flush()
				if err != nil {
					// Refreshing page in web browser will establish a new
					// SSE connection, but only (the last) one is alive, so
					// dead connections must be closed here.
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)

					break
				}
				time.Sleep(2 * time.Second)
			}
		}))

		return nil
	})

	// Start server
	log.Fatal(app.Listen(":" + appPort))
}
