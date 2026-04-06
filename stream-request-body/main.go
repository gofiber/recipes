package main

import (
	"io"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Use an external setup function in order
	// to configure the app in tests as well
	app := setup()

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// setup creates and configures a Fiber app with all routes.
func setup() *fiber.App {
	// Fiber instance
	app := fiber.New()
	// Enable request body streaming.
	// Note that this may have consequences if typical requests are larger
	// than the configured limit as it's uncertain if Fiber will gracefully
	// handle the streaming if it expects a complete body.
	app.Server().StreamRequestBody = true

	// You can test the route by using cURL:
	// curl -X POST --data-binary @/path/to/large/file localhost:3000
	app.Post("/", func(c fiber.Ctx) error {
		reader := c.RequestCtx().RequestBodyStream()
		if reader == nil {
			return nil
		}
		// Read 1MiB at a time
		buffer := make([]byte, 0, 1024*1024)
		for {
			length, err := io.ReadFull(reader, buffer[:cap(buffer)])
			// Cap the buffer based on the actual length read
			buffer = buffer[:length]
			if length > 0 {
				// Process the chunk - e.g., write to file, parse data, etc.
				log.Printf("Read %d bytes", length)
			}
			if err != nil {
				// When the error is EOF or ErrUnexpectedEOF, there are no
				// longer any bytes to read meaning the request is completed.
				// ErrUnexpectedEOF means the last chunk was smaller than the
				// buffer, which is normal for the final (or only) chunk.
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					break
				}
				return err
			}
		}
		return nil
	})

	return app
}
