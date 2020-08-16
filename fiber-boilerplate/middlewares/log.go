package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/libraries"
	"net/http"
	"runtime/debug"
	"time"
)

//Middleware requestid + logger + recover for request traceability
func LogMiddleware(c *fiber.Ctx) {
	start := time.Now()
	rid := c.Get(fiber.HeaderXRequestID)
	if rid == "" {
		rid = uuid.New().String()
		c.Set(fiber.HeaderXRequestID, rid)
	}

	fields := &libraries.LogFields{
		RequestID: rid,
		RemoteIP:  c.IP(),
		Method:    c.Method(),
		Host:      c.Hostname(),
		Path:      c.Path(),
		Protocol:  c.Protocol(),
	}

	defer func() {
		rvr := recover()

		if rvr != nil {
			err, ok := rvr.(error)
			if !ok {
				err = fmt.Errorf("%v", rvr)
			}

			fields.Error = err
			fields.Stack = debug.Stack()

			c.Status(http.StatusInternalServerError)
			c.JSON(map[string]interface{}{
				"status": http.StatusText(http.StatusInternalServerError),
			})
		}

		fields.StatusCode = c.Fasthttp.Response.StatusCode()
		fields.Latency = time.Since(start).Seconds()

		switch {
		case rvr != nil:
			Log.Error().EmbedObject(fields).Msg("panic recover")
		case fields.StatusCode >= 500:
			Log.Error().EmbedObject(fields).Msg("server error")
		case fields.StatusCode >= 400:
			Log.Error().EmbedObject(fields).Msg("client error")
		case fields.StatusCode >= 300:
			Log.Warn().EmbedObject(fields).Msg("redirect")
		case fields.StatusCode >= 200:
			Log.Info().EmbedObject(fields).Msg("success")
		case fields.StatusCode >= 100:
			Log.Info().EmbedObject(fields).Msg("informative")
		default:
			Log.Warn().EmbedObject(fields).Msg("unknown status")
		}
	}()
	c.Next()
}
