package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	registryclient "code.hollensbe.org/erikh/spin/clients/registry"
	"code.hollensbe.org/erikh/spin/pkg/services"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Manage the spin-broker, a message bus for Spin."

	app.Commands = []*cli.Command{
		{
			Name:        "start",
			Usage:       "start the service",
			Description: "start the service",
			Action:      start,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "registry-host",
					Aliases: []string{"r"},
					Value:   "localhost:8081",
				},
				&cli.StringFlag{
					Name:    "broker-host",
					Aliases: []string{"b"},
					Value:   "localhost:8080",
				},
				&cli.StringFlag{
					Name:    "listen",
					Aliases: []string{"l"},
					Value:   "localhost:8082",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func start(ctx *cli.Context) error {
	broker := brokerclient.Config{Host: ctx.String("broker-host"), Timeout: 1}
	registry := registryclient.Config{Host: ctx.String("registry-host"), Timeout: 1}
	handler, err := services.APIServer(true, broker, registry)
	if err != nil {
		return err
	}

	listen := ctx.String("listen")

	srv := &http.Server{Addr: listen, Handler: handler}
	errc := make(chan error, 2)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Start HTTP server in a separate goroutine.
	go func() {
		log.Printf("APIServer listening on %q", listen)
		errc <- srv.ListenAndServe()
	}()

	log.Printf("Shutting down: %v", <-errc)

	// Shutdown gracefully with a 30s timeout.
	cCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return srv.Shutdown(cCtx)
}
