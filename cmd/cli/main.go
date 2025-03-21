package main

import (
	"context"

	"github.com/skolldire/go-batch-app/cmd/cli/processors/hello_world"
)

func main() {
	ctx := context.Background()
	app := hello_world.NewService()
	app.Run(ctx)
}
