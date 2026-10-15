package modules

import (
	"github.com/gofiber/fiber/v3"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

var _ types.ModuleFunc = ConfigureFiber

// ConfigureFiber Configures all services required by the Fiber app.
func ConfigureFiber(registry types.ServiceRegistry) error {
	if err := registration.RegisterInstance(registry, fiber.Config{
		AppName:   "parsley-app-recipe",
		Immutable: true,
	}); err != nil {
		return err
	}

	if err := registry.Register(newFiber, types.LifetimeSingleton); err != nil {
		return err
	}

	return registry.RegisterModule(RegisterRouteHandlers)
}

// newFiber Activator method for new Fiber instances.
func newFiber(config fiber.Config) *fiber.App {
	return fiber.New(config)
}
