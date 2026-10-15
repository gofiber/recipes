package main

import (
	"context"
	"log"

	"github.com/gofiber/recipes/parsley-app/internal"
	"github.com/gofiber/recipes/parsley-app/internal/modules"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

func main() {

	ctx := context.Background()

	// Runs a Fiber instance as a Parsley-enabled app
	if err := bootstrap.RunParsleyApplication(ctx, internal.NewApp,
		modules.ConfigureFiber,
		modules.ConfigureGreeter); err != nil {
		log.Fatal(err)
	}
}
