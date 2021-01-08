package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	brokerclient "github.com/erikh/spin/clients/broker"
	spinbroker "github.com/erikh/spin/gen/spin_broker"
	"github.com/erikh/spin/pkg/services"
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

	app.Usage = "Manage the spin-broker, a message bus for Spin."
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

	app.Commands = []*cli.Command{
		{
			Name:        "start",
			Usage:       "Start the service",
			Description: "Start the service",
			Action:      start,
		},
		{
			Name:        "message",
			Usage:       "Message the broker",
			Description: "Message the broker",
			Subcommands: []*cli.Command{
				{
					Name:        "new",
					Usage:       "Create a new package. Returns a UUID.",
					Description: "Create a new package. Returns a UUID.",
					ArgsUsage:   " ",
					Action:      messageNew,
				},
				{
					Name:        "add",
					Usage:       "Add a command to a package. Returns the UUID of the added command.",
					Description: "Add a command to a package. Returns the UUID of the added command.",
					ArgsUsage:   "[pkg uuid] [resource] [action] [key=value parameters...]",
					Action:      messageAdd,
				},
				{
					Name:        "enqueue",
					Usage:       "Enqueue a package's items",
					Description: "Enqueue a package's items",
					ArgsUsage:   "[pkg uuid]",
					Action:      messageEnqueue,
				},
				{
					Name:        "status",
					Usage:       "Obtain status of a package",
					Description: "Obtain status of a package",
					ArgsUsage:   "[pkg uuid]",
					Action:      messageStatus,
				},
				{
					Name:        "complete",
					Usage:       "Set a status for a specific command, by UUID",
					Description: "Set a status for a specific command, by UUID",
					ArgsUsage:   "[command uuid] [true/false] [reason (optional)]",
					Action:      messageComplete,
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
	handler, err := services.Broker(ctx.String("dbpath"), false)
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
		log.Printf("Broker listening on %q", host)
		errc <- srv.ListenAndServe()
	}()

	log.Printf("Shutting down: %v", <-errc)

	// Shutdown gracefully with a 30s timeout.
	cCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return srv.Shutdown(cCtx)
}

func messageNew(ctx *cli.Context) error {
	cc := brokerclient.Config{
		Host:    ctx.String("host"),
		Timeout: 1,
	}

	client := brokerclient.New(cc)
	pkg, err := client.New(context.Background())
	if err != nil {
		return err
	}

	fmt.Println(pkg)
	return nil
}

func messageAdd(ctx *cli.Context) error {
	if ctx.Args().Len() < 3 {
		return errors.New("invalid arguments. try --help")
	}

	pkg := ctx.Args().Get(0)
	resource := ctx.Args().Get(1)
	action := ctx.Args().Get(2)
	parameters := map[string]interface{}{}

	if pkg == "" || resource == "" || action == "" {
		return errors.New("invalid parameters. try --help")
	}

	for i := 3; ctx.Args().Get(i) != ""; i++ {
		param := strings.SplitN(ctx.Args().Get(i), "=", 2)
		if len(param) != 2 {
			return errors.New("invalid key=value parameters")
		}

		var value interface{}

		if err := json.Unmarshal([]byte(param[1]), &value); err != nil {
			parameters[param[0]] = param[1]
		} else {
			parameters[param[0]] = value
		}
	}

	cc := brokerclient.Config{
		Host:    ctx.String("host"),
		Timeout: 1,
	}

	client := brokerclient.New(cc)
	uuid, err := client.Add(context.Background(), &spinbroker.AddPayload{
		ID:         pkg,
		Resource:   resource,
		Action:     action,
		Parameters: parameters,
	})
	if err != nil {
		return err
	}

	fmt.Println(uuid)
	return nil
}

func messageEnqueue(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments. try --help")
	}

	pkg := ctx.Args().First()
	if pkg == "" {
		return errors.New("invalid arguments. try --help")
	}

	cc := brokerclient.Config{
		Host:    ctx.String("host"),
		Timeout: 1,
	}

	client := brokerclient.New(cc)
	res, err := client.Enqueue(context.Background(), pkg)
	if err != nil {
		return err
	}

	for _, r := range res {
		fmt.Println(r)
	}

	return nil
}

func messageStatus(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments. try --help")
	}

	pkg := ctx.Args().First()
	if pkg == "" {
		return errors.New("invalid arguments. try --help")
	}

	cc := brokerclient.Config{
		Host:    ctx.String("host"),
		Timeout: 1,
	}

	client := brokerclient.New(cc)
	status, err := client.Status(context.Background(), pkg)
	if err != nil {
		return err
	}

	if !status.Status {
		fmt.Printf("Error during processing: %v (causer: %v)\n", *status.Reason, *status.Causer)
		os.Exit(1)
	}

	return nil
}

func messageComplete(ctx *cli.Context) error {
	if ctx.Args().Len() < 2 {
		return errors.New("invalid arguments. try --help")
	}

	command := ctx.Args().Get(0)
	result, err := strconv.ParseBool(ctx.Args().Get(1))
	if err != nil {
		return err
	}

	var sr *string

	if !result {
		s := ctx.Args().Get(2)
		sr = &s
	}

	cc := brokerclient.Config{
		Host:    ctx.String("host"),
		Timeout: 1,
	}
	client := brokerclient.New(cc)
	return client.Complete(context.Background(), command, result, sr)
}
