package main

import (
	"context"
	"flag"
	"github.com/BGrewell/boop/internal"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func processFlags() {

	flag.String("ip", "", "ip address to listen on")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if viper.GetString("ip") == "" {
		log.Fatal("ip address is required")
	}
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

	<-ctx.Done()

}
