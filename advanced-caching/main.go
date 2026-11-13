package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
)

// cacheGenerationHeader carries the tag generation into the cache key; any client value is overwritten.
const cacheGenerationHeader = "X-Cache-Generation"

// InvalidationRegistry hands out a cache generation per tag.
type InvalidationRegistry struct {
	mu          sync.RWMutex
	generations map[string]uint64
}

// NewInvalidationRegistry initializes a thread-safe registry.
func NewInvalidationRegistry() *InvalidationRegistry {
	return &InvalidationRegistry{
		generations: make(map[string]uint64),
	}
}

// InvalidateTag starts a new generation for tag, so its next requests miss the old entries.
func (r *InvalidationRegistry) InvalidateTag(tag string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.generations[tag]++
}

// Generation returns the current generation of tag.
func (r *InvalidationRegistry) Generation(tag string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return strconv.FormatUint(r.generations[tag], 10)
}

// SetupApp builds the Fiber app with advanced cache patterns.
func SetupApp(adminKey string) (*fiber.App, *InvalidationRegistry) {
	app := fiber.New()
	registry := NewInvalidationRegistry()

	// 1. Advanced Caching Middleware with KeyHeaders.
	// HTTP method restrictions (GET/HEAD only), Cache-Control directives (no-cache, no-store),
	// canonical query sorting (?a=1&b=2 == ?b=2&a=1), and header segment escaping
	// are handled automatically by middleware/cache out of the box.
	cacheMiddleware := cache.New(cache.Config{
		Expiration:  30 * time.Second,
		CacheHeader: "X-Cache-Status",
		KeyHeaders:  []string{"Accept", "Accept-Encoding", "Accept-Language", "X-Tenant-ID", "X-User-Role", cacheGenerationHeader},
	})

	// CacheInvalidator only sees the request, not the stored entry, so it cannot tell
	// stale entries from fresh ones. A generation in the key can: old entries just expire.
	tagGeneration := func(c fiber.Ctx) error {
		// The cache key sorts repeated values, so ?category=a&category=b would share an entry with its reverse.
		if len(c.Request().URI().QueryArgs().PeekMulti("category")) > 1 {
			return fiber.NewError(fiber.StatusBadRequest, "category must be given at most once")
		}
		c.Request().Header.Set(cacheGenerationHeader, registry.Generation(c.Query("category", "general")))
		return c.Next()
	}

	api := app.Group("/api", tagGeneration, cacheMiddleware)

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

	// 2. Cache Invalidation Endpoint, protected by a bearer API key
	adminKeyHash := sha256.Sum256([]byte(adminKey))
	requireAdminKey := keyauth.New(keyauth.Config{
		Validator: func(_ fiber.Ctx, key string) (bool, error) {
			// Comparing hashes keeps the check constant-time even when the lengths differ.
			keyHash := sha256.Sum256([]byte(key))
			if subtle.ConstantTimeCompare(adminKeyHash[:], keyHash[:]) == 1 {
				return true, nil
			}
			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	})

	app.Post("/api/items/invalidate", requireAdminKey, func(c fiber.Ctx) error {
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
			"status":          "invalidated",
		})
	})

	// Health check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	return app, registry
}

func main() {
	adminKey := os.Getenv("ADMIN_API_KEY")
	if adminKey == "" {
		log.Fatal("ADMIN_API_KEY must be set, it protects POST /api/items/invalidate")
	}

	app, _ := SetupApp(adminKey)
	log.Fatal(app.Listen(":3000"))
}
