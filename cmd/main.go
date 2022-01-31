package main

import (
	"context"
	"github.com/BGrewell/boop/internal"
	"log"
)

func processFlags() {

}

func main() {

	processFlags()
	ctx := context.Background()
	controller, err := internal.NewProxyController(ctx)
	if err != nil {
		log.Fatalf("error creating controller: %v\n", err)
	}

	err = controller.Start(ctx)
	if err != nil {
		log.Fatalf("error starting: %v\n", err)
	}

}
