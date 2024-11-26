package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

var _ types.ModuleFunc = ConfigureFiber

// ConfigureFiber Configures all services required by the Fiber app.
func ConfigureFiber(registry types.ServiceRegistry) error {
	registration.RegisterInstance(registry, fiber.Config{
		AppName:   "parsley-app-recipe",
		Immutable: true,
	})

	registry.Register(newFiber, types.LifetimeSingleton)
	registry.RegisterModule(RegisterRouteHandlers)

	return nil
}

// newFiber Activator method for new Fiber instances.
func newFiber(config fiber.Config) *fiber.App {
	return fiber.New(config)
}
