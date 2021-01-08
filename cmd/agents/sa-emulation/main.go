package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	brokerclient "github.com/erikh/spin/clients/broker"
	"github.com/erikh/spin/pkg/resources/emulation"
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
		&cli.StringFlag{
			Name:    "monitor-directory",
			Aliases: []string{"md"},
			Usage:   "The directory where the QMP monitor sockets will be placed",
		},
	}

	app.Action = start

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func start(ctx *cli.Context) error {
	agent, err := emulation.NewAgent(emulation.AgentConfig{
		ClientConfig: brokerclient.Config{
			Host:    "localhost:8080",
			Timeout: 1,
		},
		SystemDir:  ctx.String("service-directory"),
		MonitorDir: ctx.String("monitor-directory"),
	})
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
