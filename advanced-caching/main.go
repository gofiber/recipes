package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
)

// InvalidationRegistry tracks cache tags and determines invalidation rules.
type InvalidationRegistry struct {
	mu            sync.RWMutex
	invalidatedAt map[string]time.Time
}

// NewInvalidationRegistry initializes a thread-safe registry.
func NewInvalidationRegistry() *InvalidationRegistry {
	return &InvalidationRegistry{
		invalidatedAt: make(map[string]time.Time),
	}
}

// InvalidateTag marks all cache entries under a given tag invalid from this moment.
func (r *InvalidationRegistry) InvalidateTag(tag string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.invalidatedAt[tag] = time.Now()
}

// IsInvalidated checks if an entry tagged with tag was invalidated after it was generated.
func (r *InvalidationRegistry) IsInvalidated(tag string, entryTimestamp time.Time) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	invTime, exists := r.invalidatedAt[tag]
	if !exists {
		return false
	}
	return !entryTimestamp.After(invTime)
}

// generateDynamicKey computes a deterministic cache key that segments by role,
// authenticated tenant/user, and query parameters.
func generateDynamicKey(c fiber.Ctx) string {
	role := c.Get("X-User-Role", "anonymous")
	tenant := c.Get("X-Tenant-ID", "global")
	path := c.Path()
	rawQuery := string(c.Request().URI().QueryString())

	rawKey := fmt.Sprintf("%s:%s:%s?%s", role, tenant, path, rawQuery)
	hash := sha256.Sum256([]byte(rawKey))
	return hex.EncodeToString(hash[:])
}

// AppConfig configures the Fiber application and cache behavior.
type AppConfig struct {
	Storage fiber.Storage
}

// SetupApp builds the Fiber app with advanced cache patterns.
func SetupApp(cfg ...AppConfig) (*fiber.App, *InvalidationRegistry) {
	app := fiber.New()
	registry := NewInvalidationRegistry()

	var storage fiber.Storage
	if len(cfg) > 0 && cfg[0].Storage != nil {
		storage = cfg[0].Storage
	}

	// 1. Advanced Caching Middleware with KeyGenerator, CacheInvalidator & Conditional Caching
	cacheMiddleware := cache.New(cache.Config{
		Expiration:   30 * time.Second,
		CacheHeader:  "X-Cache-Status",
		Storage:      storage,
		KeyGenerator: generateDynamicKey,
		Next: func(c fiber.Ctx) bool {
			// Skip caching for mutations or when client explicitly bypasses cache
			if c.Method() != fiber.MethodGet && c.Method() != fiber.MethodHead {
				return true
			}
			cacheControl := c.Get("Cache-Control")
			return strings.Contains(cacheControl, "no-cache") || strings.Contains(cacheControl, "no-store")
		},
		CacheInvalidator: func(c fiber.Ctx) bool {
			tag := c.Query("category")
			if tag == "" {
				if id := c.Params("id"); id != "" {
					tag = "item:" + id
				}
			}
			if tag != "" {
				// If tag was marked invalidated, invalidate the cached entry
				return registry.IsInvalidated(tag, time.Time{})
			}
			return false
		},
	})

	api := app.Group("/api", cacheMiddleware)

	// Sample cached endpoint: items listing segmented by tenant & query
	api.Get("/items", func(c fiber.Ctx) error {
		tag := c.Query("category", "general")

		return c.JSON(fiber.Map{
			"timestamp": time.Now().UnixNano(),
			"category":  tag,
			"items": []string{
				"item-alpha",
				"item-beta",
			},
		})
	})

	// Sample cached endpoint: specific item details
	api.Get("/items/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		return c.JSON(fiber.Map{
			"id":        id,
			"timestamp": time.Now().UnixNano(),
			"status":    "active",
		})
	})

	// 2. Cache Invalidation Endpoint
	app.Post("/api/items/invalidate", func(c fiber.Ctx) error {
		var req struct {
			Tag string `json:"tag"`
		}
		if err := c.Bind().Body(&req); err != nil || req.Tag == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "tag is required",
			})
		}

		registry.InvalidateTag(req.Tag)

		return c.JSON(fiber.Map{
			"invalidated_tag": req.Tag,
			"status":          "purged",
		})
	})

	// Health check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	return app, registry
}

func main() {
	app, _ := SetupApp()
	log.Fatal(app.Listen(":3000"))
}
