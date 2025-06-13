package main

import (
	"context"

	"github.com/gofiber/recipes/parsley-app/internal"
	"github.com/gofiber/recipes/parsley-app/internal/modules"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

func main() {

	ctx := context.Background()

	// Runs a Fiber instance as a Parsley-enabled app
	bootstrap.RunParsleyApplication(ctx, internal.NewApp,
		modules.ConfigureFiber,
		modules.ConfigureGreeter)
}
