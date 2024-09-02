package modules

import (
	"parsley-app/internal/services"

	"github.com/matzefriedrich/parsley/pkg/types"
)

// ConfigureGreeter Configures the services.Greeter service dependencies.
func ConfigureGreeter(registry types.ServiceRegistry) error {
	registry.Register(services.NewGreeterFactory("Hi"), types.LifetimeTransient)
	return nil
}
