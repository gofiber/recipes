package main

import (
	"context"

	"parsley-app/internal"
	"parsley-app/internal/modules"

	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

func main() {
	context := context.Background()

	// Runs a Fiber instance as a Parsley-enabled app
	bootstrap.RunParsleyApplication(context, internal.NewApp,
		modules.ConfigureFiber,
		modules.ConfigureGreeter)
}
