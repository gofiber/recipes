package main

import (
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()
	// Enable request body streaming.
	// Note that this may have consequences if typical requests are larger
	// than the configured limit as it's uncertain if Fiber will gracefully
	// handle the streaming if it expects a complete body.
	app.Server().StreamRequestBody = true

	// You can test the route by using cURL:
	// curl -X POST --data-binary @/path/to/large/file localhost:3000
	app.Post("/", func(c *fiber.Ctx) error {
		reader := c.Context().RequestBodyStream()
		// Read 1MiB at a time
		buffer := make([]byte, 0, 1024*1024)
		for {
			length, err := io.ReadFull(reader, buffer[:cap(buffer)])
			// Cap the buffer based on the actual length read
			buffer = buffer[:length]
			if err != nil {
				// When the error is EOF, there are no longer any bytes to read
				// meaning the request is completed
				if err == io.EOF {
					break
				}

				// If the error is an unexpected EOF, the requested size to read
				// was larger than what was available. This is not an issue for
				// as long as the length returned above is used, or the buffer
				// is capped as it is above. Any error that is not an unexpected
				// EOF is an actual error, which we handle accordingly
				if err != io.ErrUnexpectedEOF {
					return err
				}
			}

			// You may now use the buffer to handle the chunk of length bytes
			log.Printf("Read %d bytes: %x ...", length, buffer[0])
		}
		return nil
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
