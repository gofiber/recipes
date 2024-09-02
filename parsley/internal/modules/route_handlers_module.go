package modules

import (
	"parsley-app/internal/route_handlers"

	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// RegisterRouteHandlers Registers all route handlers of the GoFiber app.
func RegisterRouteHandlers(registry types.ServiceRegistry) error {

	// The Parsley app requires a list of RouteHandler instances, so we must enable it.
	features.RegisterList[route_handlers.RouteHandler](registry)

	// RouteHandler implementations are going to be registered here.
	registration.RegisterTransient(registry, route_handlers.NewGreeterRouteHandler)

	return nil
}
