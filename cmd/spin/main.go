package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"code.hollensbe.org/erikh/spin/clients/api"
	spinapiserver "code.hollensbe.org/erikh/spin/gen/spin_apiserver"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Spin is a tool for managing VMs"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "host",
			Aliases: []string{"t"},
			Usage:   "host:port to communicate with",
			Value:   "localhost:8082",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:      "start",
			Usage:     "Start a VM by ID",
			ArgsUsage: "[id]",
			Action:    start,
		},
		{
			Name:      "stop",
			Usage:     "Stop a VM by ID; this will attempt a graceful, then force kill after 30s",
			ArgsUsage: "[id]",
			Action:    stop,
		},
		{
			Name:      "shutdown",
			Usage:     "Gracefully shutdown a VM by ID. Will not force",
			ArgsUsage: "[id]",
			Action:    shutdown,
		},
		{
			Name:  "vm",
			Usage: "Manipulate VMs",
			Subcommands: []*cli.Command{
				{
					Name:      "list",
					Usage:     "List all VMs by ID + Name",
					ArgsUsage: " ",
					Action:    list,
				},
				{
					Name:      "delete",
					Usage:     "Delete a VM by ID",
					ArgsUsage: "[id]",
					Action:    delete,
				},
				{
					Name:      "create",
					Usage:     "Create a VM",
					ArgsUsage: "[name] [volume]",
					Action:    create,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "cdrom",
							Aliases: []string{"d"},
						},
						&cli.UintFlag{
							Name:    "cpus",
							Aliases: []string{"c"},
							Value:   2,
						},
						&cli.UintFlag{
							Name:    "memory",
							Usage:   "Memory, in megabytes",
							Aliases: []string{"m"},
							Value:   1024,
						},
						&cli.UintFlag{
							Name:    "image-size",
							Usage:   "Image size, in gigabytes",
							Aliases: []string{"s"},
							Value:   10,
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getClient(ctx *cli.Context) *api.Client {
	return api.New(api.Config{
		Host:    ctx.String("host"),
		Timeout: 1,
	})
}

func stringPtr(s string) *string {
	return &s
}

func uintPtr(u uint) *uint {
	return &u
}

func create(ctx *cli.Context) error {
	if ctx.Args().Len() != 2 {
		return errors.New("invalid arguments; see --help")
	}

	name := ctx.Args().Get(0)
	volume := ctx.Args().Get(1)

	vm := &spinapiserver.CreateVM{
		Name:   name,
		Cpus:   ctx.Uint("cpus"),
		Memory: ctx.Uint("memory"),
		Storage: []*spinapiserver.Storage{
			{
				Volume:    stringPtr(volume),
				Image:     name,
				ImageSize: uintPtr(ctx.Uint("image-size")),
			},
		},
	}

	if ctx.String("cdrom") != "" {
		vm.Storage = append(vm.Storage, &spinapiserver.Storage{
			Cdrom: true,
			Image: ctx.String("cdrom"),
		})
	}

	id, err := getClient(ctx).VMCreate(context.Background(), vm)
	if err != nil {
		return err
	}

	fmt.Println(id)
	return nil
}

func delete(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	id, err := strconv.ParseUint(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}

	return getClient(ctx).VMDelete(context.Background(), id)
}

func start(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	id, err := strconv.ParseUint(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}

	return getClient(ctx).ControlStart(context.Background(), id)
}

func stop(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	id, err := strconv.ParseUint(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}

	return getClient(ctx).ControlStop(context.Background(), id)
}

func shutdown(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	id, err := strconv.ParseUint(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}

	return getClient(ctx).ControlShutdown(context.Background(), id)
}

func list(ctx *cli.Context) error {
	ids, err := getClient(ctx).VMList(context.Background())
	if err != nil {
		return err
	}

	for _, id := range ids {
		fmt.Println(id)
	}

	return nil
}
