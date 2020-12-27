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

	"code.hollensbe.org/erikh/spin/pkg/services"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Authors = []*cli.Author{
		{
			Email: "github@hollensbe.org",
			Name:  "Erik Hollensbe",
		},
	}

	app.Usage = "Start the spin-broker, a message bus for Spin."
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "dbpath",
			Aliases: []string{"db"},
			Usage:   "Path to database file",
			Value:   "spin-broker.db",
		},
		&cli.StringFlag{
			Name:    "host",
			Aliases: []string{"t"},
			Usage:   "host:port to listen on",
			Value:   "localhost:8080",
		},
	}

	app.Action = start

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func start(ctx *cli.Context) error {
	handler, err := services.Broker(ctx.String("dbpath"))
	if err != nil {
		return err
	}

	host := ctx.String("host")

	srv := &http.Server{Addr: host, Handler: handler}
	errc := make(chan error, 2)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Start HTTP server in a separate goroutine.
	go func() {
		log.Printf("HTTP server listening on %q", host)
		errc <- srv.ListenAndServe()
	}()

	log.Printf("Shutting down: %v", <-errc)

	// Shutdown gracefully with a 30s timeout.
	cCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return srv.Shutdown(cCtx)
}
