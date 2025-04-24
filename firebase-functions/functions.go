package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	adaptor "github.com/gofiber/fiber/v2/middleware/adaptor"
)

// CloudFunctionRouteToFiber route cloud function http.Handler to *fiber.App
// Internally, google calls the function with the /execute base URL
func CloudFunctionRouteToFiber(fiberApp *fiber.App, w http.ResponseWriter, r *http.Request) {
	adaptor.FiberApp(fiberApp)(w, r)
}
