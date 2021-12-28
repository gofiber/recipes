package handler

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// GEO will fetch IP GEO data from ipapi.co
// Original Code at https://github.com/gofiber/recipes/blob/master/geoip
func GEO() fiber.Handler {
	// Create fasthttp client
	client := fasthttp.Client{}

	// Return handler
	return func(c *fiber.Ctx) error {
		// Get domain from param else default to remote IP
		ip := c.Params("ip", c.IP())

		// Get request/response from pool
		req := fasthttp.AcquireRequest()
		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(res)

		// Set request URL
		req.SetRequestURI("https://ipapi.co/" + ip + "/json")

		// Make request
		if err := client.DoTimeout(req, res, 30*time.Second); err != nil {
			return err
		}

		// Check statuscode
		if res.StatusCode() != fiber.StatusOK {
			return errors.New("invalid statuscode")
		}

		// Set correct content-type
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		// Send body
		return c.Send(res.Body())
	}
}
