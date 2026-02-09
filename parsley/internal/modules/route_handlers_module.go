package modules

import (
	"context"

	"github.com/gofiber/recipes/parsley-app/internal/route_handlers"
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// RegisterRouteHandlers Registers all route handlers for the GoFiber app using the Parsley dependency injection framework.
// It sets up the necessary features and registers individual route handlers.
//
// Parameters:
// - registry: The ServiceRegistry instance that keeps track of service registrations.
//
// Returns:
// - error: Any error that occurred during the registration process.
func RegisterRouteHandlers(registry types.ServiceRegistry) error {
	// The Parsley app requires a list of RouteHandler instances, so we must enable it.
	if err := features.RegisterList[route_handlers.RouteHandler](context.Background(), registry); err != nil {
		return err
	}

	// RouteHandler implementations are going to be registered here (add more as route handler registrations as needed)
	registration.RegisterTransient(registry, route_handlers.NewGreeterRouteHandler)

	return nil
}
