package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	"code.hollensbe.org/erikh/spin/pkg/resources/emulation"
	"github.com/urfave/cli/v2"
	"golang.org/x/sys/unix"
)

func main() {
	app := cli.NewApp()
	app.Description = "Emulation agent for Spin"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "service-directory",
			Aliases: []string{"d"},
			Usage:   "The directory where the .service files will be placed",
		},
	}

	app.Action = start

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func start(ctx *cli.Context) error {
	agent, err := emulation.NewAgent(brokerclient.Config{
		Host:    "localhost:8080",
		Timeout: 1,
	}, ctx.String("service-directory"))
	if err != nil {
		return err
	}

	cCtx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, unix.SIGTERM, unix.SIGINT)
		<-sigChan
		cancel()
	}()

	return agent.Loop(cCtx)
}
