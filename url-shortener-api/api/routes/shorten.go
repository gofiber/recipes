package routes

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/amalshaji/stoyle/database"
	"github.com/amalshaji/stoyle/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

// ShortenURL validates a URL, applies rate limiting, and stores a short alias in Redis.
func ShortenURL(c fiber.Ctx) error {
	body := new(request)
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	ctx := c.Context()

	// Rate limiting: each IP gets API_QUOTA calls per 30 minutes.
	_, err := database.DB1.Get(ctx, c.IP()).Result()
	if err == redis.Nil {
		// First request from this IP — initialise quota.
		quota := os.Getenv("API_QUOTA")
		if setErr := database.DB1.Set(ctx, c.IP(), quota, 30*60*time.Second).Err(); setErr != nil {
			log.Printf("warn: failed to set rate limit key for %s: %v", c.IP(), setErr)
		}
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	} else {
		val, getErr := database.DB1.Get(ctx, c.IP()).Result()
		if getErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "cannot read rate limit from DB",
			})
		}
		valInt, convErr := strconv.Atoi(val)
		if convErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "invalid rate limit value in DB",
			})
		}
		if valInt <= 0 {
			limit, ttlErr := database.DB1.TTL(ctx, c.IP()).Result()
			if ttlErr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "cannot read rate limit TTL from DB",
				})
			}
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":            "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "haha... nice try",
		})
	}

	body.URL = helpers.EnforceHTTP(body.URL)

	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	existing, getErr := database.DB0.Get(ctx, id).Result()
	if getErr != nil && getErr != redis.Nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}
	if existing != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL short already in use",
		})
	}

	if body.Expiry == 0 {
		body.Expiry = 24 // default 24 hours
	}
	if err = database.DB0.Set(ctx, id, body.URL, body.Expiry*3600*time.Second).Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}

	resp := response{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  10,
		XRateLimitReset: 30,
	}

	if err = database.DB1.Decr(ctx, c.IP()).Err(); err != nil {
		log.Printf("warn: failed to decrement rate limit for %s: %v", c.IP(), err)
	}

	val, getErr := database.DB1.Get(ctx, c.IP()).Result()
	if getErr == nil {
		if n, convErr := strconv.Atoi(val); convErr == nil {
			resp.XRateRemaining = n
		} else {
			log.Printf("warn: unexpected rate limit value %q for %s: %v", val, c.IP(), convErr)
		}
	}

	ttl, ttlErr := database.DB1.TTL(ctx, c.IP()).Result()
	if ttlErr == nil {
		resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute
	}

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id
	return c.Status(fiber.StatusOK).JSON(resp)
}
