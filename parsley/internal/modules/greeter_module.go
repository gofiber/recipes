package modules

import (
	"github.com/gofiber/recipes/parsley-app/internal/services"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// ConfigureGreeter Configures the services.Greeter service dependencies.
func ConfigureGreeter(registry types.ServiceRegistry) error {
	return registry.Register(services.NewGreeterFactory, types.LifetimeTransient)
}
