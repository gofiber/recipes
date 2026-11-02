// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	"os"
	"sync"

	"github.com/gofiber/fiber/v3"
)

// DynamicRoute represents a runtime-managed route configuration.
type DynamicRoute struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Response string `json:"response"`
	Enabled  bool   `json:"enabled"`
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

func (rm *RouteManager) add(r DynamicRoute) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.routes[r.Path] = r
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

func setupApp(rm *RouteManager) *fiber.App {
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

	// Admin Route Management APIs
	admin := app.Group("/admin/routes")

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
		if req.Path == "" || req.Response == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "path and response are required"})
		}

		req.Enabled = true
		rm.add(req)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Route registered successfully",
			"route":   req,
		})
	})

	// Remove/unregister a dynamic route at runtime
	admin.Delete("/:name", func(c fiber.Ctx) error {
		name := c.Params("name")
		path := "/" + name
		if rm.remove(path) {
			return c.JSON(fiber.Map{"message": "Route removed", "path": path})
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Route not found"})
	})

	// Dynamic route dispatcher for user endpoints
	app.Get("/d/*", func(c fiber.Ctx) error {
		subPath := "/" + c.Params("*")
		route, exists := rm.get(subPath)
		if !exists || !route.Enabled {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Dynamic route not found or disabled"})
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
		Enabled:  true,
	})

	app := setupApp(rm)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server listening on http://127.0.0.1:%s", port)
	log.Fatal(app.Listen(":" + port))
}
