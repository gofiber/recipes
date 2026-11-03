// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
)

// DynamicRoute represents a runtime-managed route configuration.
type DynamicRoute struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Response string `json:"response"`
}

// RouteManager manages runtime dynamic routes thread-safely.
type RouteManager struct {
	mu     sync.RWMutex
	routes map[string]DynamicRoute
}

func newRouteManager() *RouteManager {
	return &RouteManager{
		routes: make(map[string]DynamicRoute),
	}
}

// add reports false when the path is already registered.
func (rm *RouteManager) add(r DynamicRoute) bool {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if _, exists := rm.routes[r.Path]; exists {
		return false
	}
	rm.routes[r.Path] = r
	return true
}

func (rm *RouteManager) remove(path string) bool {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if _, exists := rm.routes[path]; exists {
		delete(rm.routes, path)
		return true
	}
	return false
}

func (rm *RouteManager) get(path string) (DynamicRoute, bool) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	r, exists := rm.routes[path]
	return r, exists
}

func (rm *RouteManager) list() []DynamicRoute {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	res := make([]DynamicRoute, 0, len(rm.routes))
	for _, r := range rm.routes {
		res = append(res, r)
	}
	return res
}

func setupApp(rm *RouteManager, adminKey string) *fiber.App {
	app := fiber.New()

	// Static core routes
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Dynamic Route Management API",
			"docs":    "/admin/routes to inspect routes",
		})
	})

	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Admin Route Management APIs, protected by a bearer API key
	adminKeyHash := sha256.Sum256([]byte(adminKey))
	admin := app.Group("/admin/routes", keyauth.New(keyauth.Config{
		Validator: func(_ fiber.Ctx, key string) (bool, error) {
			// Comparing hashes keeps the check constant-time even when the lengths differ.
			keyHash := sha256.Sum256([]byte(key))
			if subtle.ConstantTimeCompare(adminKeyHash[:], keyHash[:]) == 1 {
				return true, nil
			}
			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	}))

	// List all routes (both registered dynamic routes and core stack summary)
	admin.Get("/", func(c fiber.Ctx) error {
		stackRoutes := make([]string, 0)
		for _, methodSlice := range app.Stack() {
			for _, r := range methodSlice {
				stackRoutes = append(stackRoutes, r.Method+" "+r.Path)
			}
		}

		return c.JSON(fiber.Map{
			"dynamic_routes": rm.list(),
			"static_routes":  stackRoutes,
		})
	})

	// Register a new dynamic route at runtime
	admin.Post("/", func(c fiber.Ctx) error {
		var req DynamicRoute
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
		}
		// The dispatcher looks up "/" + the wildcard, so a path without the slash could never match.
		if !strings.HasPrefix(req.Path, "/") || req.Response == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "path starting with / and response are required"})
		}
		if !rm.add(req) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Route already registered"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Route registered successfully",
			"route":   req,
		})
	})

	// Remove/unregister a dynamic route at runtime
	// A wildcard, not :name, so nested paths such as /team/features can be removed too.
	admin.Delete("/*", func(c fiber.Ctx) error {
		path := "/" + c.Params("*")
		if rm.remove(path) {
			return c.JSON(fiber.Map{"message": "Route removed", "path": path})
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Route not found"})
	})

	// Dynamic route dispatcher for user endpoints
	app.Get("/d/*", func(c fiber.Ctx) error {
		subPath := "/" + c.Params("*")
		route, exists := rm.get(subPath)
		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Dynamic route not found"})
		}
		return c.JSON(fiber.Map{
			"route":    route.Name,
			"path":     route.Path,
			"response": route.Response,
		})
	})

	return app
}

func main() {
	rm := newRouteManager()

	// Seed with an initial sample route
	rm.add(DynamicRoute{
		Name:     "welcome-promo",
		Path:     "/promo",
		Response: "Welcome to Fiber v3! Use coupon FIBER3 for 20% off.",
	})

	adminKey := os.Getenv("ADMIN_API_KEY")
	if adminKey == "" {
		log.Fatal("ADMIN_API_KEY must be set, it protects the /admin/routes API")
	}

	app := setupApp(rm, adminKey)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server listening on http://127.0.0.1:%s", port)
	log.Fatal(app.Listen(":" + port))
}
