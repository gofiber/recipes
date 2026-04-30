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
		Expiration:          exp,
		DisableCacheControl: false,
	})
}

// GeoLocation fetches the details of the IP from a public http API
func GeoLocation(c fiber.Ctx) error {
	ip := c.Params("ip")
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to fetch geo data",
		})
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to read response body",
		})
	}

	var resp response
	if err := json.Unmarshal(body, &resp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to parse geo data",
		})
	}
	if resp.Status == "fail" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "enter an ip",
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
