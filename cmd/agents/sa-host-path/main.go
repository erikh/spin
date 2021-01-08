package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	brokerclient "github.com/erikh/spin/clients/broker"
	"github.com/erikh/spin/pkg/resources/storage"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"golang.org/x/sys/unix"
)

func main() {
	app := cli.NewApp()
	app.Description = "Host-path agent for Spin"

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "root-path",
			Aliases: []string{"p"},
			Value:   filepath.Join(home, ".config/spin/images"),
		},
	}

	app.Action = start

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func start(ctx *cli.Context) error {
	agent := storage.NewHostPathAgent(ctx.String("root-path"), brokerclient.Config{
		Host:    "localhost:8080",
		Timeout: 1,
	})

	cCtx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, unix.SIGTERM, unix.SIGINT)
		<-sigChan
		cancel()
	}()

	return agent.Loop(cCtx)
}
