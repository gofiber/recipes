package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
)

type response struct {
	Status    string  `json:"status"`
	Country   string  `json:"country"`
	Region    string  `json:"regionName"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
	ISP       string  `json:"isp"`
}

// CacheRequest caches the request for subsequent use
func CacheRequest(exp time.Duration) fiber.Handler {
	return cache.New(cache.Config{
		IdleTimeout:  exp,
		CacheControl: true,
	})
}

// GeoLocation fetches the details of the IP from a public http API
func GeoLocation(c fiber.Ctx) error {
	ip := c.Params("ip")
	res, _ := http.Get("http://ip-api.com/json/" + ip)
	body, _ := io.ReadAll(res.Body)

	var resp response
	json.Unmarshal(body, &resp)
	if resp.Status == "fail" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "enter an ip",
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
