package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	registryclient "github.com/erikh/spin/clients/registry"
	"github.com/erikh/spin/pkg/services"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Description = "Control and start the spin-registry"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "dbpath",
			Aliases: []string{"db"},
			Usage:   "Path to database file",
			Value:   "spin-registry.db",
		},
		&cli.StringFlag{
			Name:    "host",
			Aliases: []string{"t"},
			Usage:   "host:port to listen on",
			Value:   "localhost:8081",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:        "start",
			ArgsUsage:   " ",
			Usage:       "Start the registry",
			Description: "Start the registry",
			Action:      start,
		},
		{
			Name:        "message",
			Usage:       "Message the registry",
			Description: "Message the registry",
			Subcommands: []*cli.Command{
				{
					Name:        "storage",
					Usage:       "Send a message to the storage subsystem",
					Description: "Send a message to the storage subsystem",
					Subcommands: []*cli.Command{
						{
							Name:        "image",
							Usage:       "Work with storage images",
							Description: "Work with storage images",
							Subcommands: []*cli.Command{
								{
									Name:        "create",
									Usage:       "Create a new storage image inside a volume. Takes JSON over STDIN.",
									Description: "Create a new storage image inside a volume. Takes JSON over STDIN.",
									Action:      messageStorageImageCreate,
								},
								{
									Name:        "delete",
									Usage:       "Delete an existing image",
									Description: "Delete an existing image",
									ArgsUsage:   "[volume] [name]",
									Action:      messageStorageImageDelete,
								},
								{
									Name:        "list",
									Usage:       "List images for a volume",
									Description: "List images for a volume",
									ArgsUsage:   " ",
									Action:      messageStorageImageList,
								},
								{
									Name:        "get",
									Usage:       "Get an image for a volume",
									Description: "Get an image for a volume",
									ArgsUsage:   "[volume] [name]",
									Action:      messageStorageImageGet,
								},
							},
						},
						{
							Name:        "volume",
							Usage:       "Work with storage volumes (volumes hold images)",
							Description: "Work with storage volumes (volumes hold images)",
							Subcommands: []*cli.Command{
								{
									Name:        "create",
									Usage:       "Create a new storage volume",
									Description: "Create a new storage volume",
									ArgsUsage:   "[name]",
									Action:      messageStorageVolumeCreate,
								},
								{
									Name:        "delete",
									Usage:       "Delete an existing storage volume",
									Description: "Delete an existing storage volume",
									ArgsUsage:   "[name]",
									Action:      messageStorageVolumeDelete,
								},
								{
									Name:        "list",
									Usage:       "List volumes",
									Description: "List volumes",
									ArgsUsage:   " ",
									Action:      messageStorageVolumeList,
								},
							},
						},
					},
				},
				{
					Name:        "vm",
					Usage:       "Send a message to the VM subsystem",
					Description: "Send a message to the VM subsystem",
					Subcommands: []*cli.Command{
						{
							Name:        "create",
							Usage:       "Create a new VM, accepting JSON over STDIN. Returns a uint64 ID.",
							Description: "Create a new VM, accepting JSON over STDIN. Returns a uint64 ID.",
							ArgsUsage:   " ",
							Action:      messageVMCreate,
						},
						{
							Name:        "update",
							Usage:       "Update a VM, accepting JSON over STDIN and the ID to update as a parameter. Returns a uint64 ID.",
							Description: "Update a VM, accepting JSON over STDIN and the ID to update as a parameter. Returns a uint64 ID.",
							ArgsUsage:   "[id]",
							Action:      messageVMUpdate,
						},
						{
							Name:        "delete",
							Usage:       "Delete a VM, accepting an ID as a parameter.",
							Description: "Delete a VM, accepting an ID as a parameter.",
							ArgsUsage:   "[id]",
							Action:      messageVMDelete,
						},
						{
							Name:        "get",
							Usage:       "Get a VM, accepting an ID as a parameter. Returns a JSON document describing the VM.",
							Description: "Get a VM, accepting an ID as a parameter. Returns a JSON document describing the VM.",
							ArgsUsage:   "[id]",
							Action:      messageVMGet,
						},
						{
							Name:        "list",
							Usage:       "List the IDs of all VMs.",
							Description: "List the IDs of all VMs.",
							ArgsUsage:   " ",
							Action:      messageVMList,
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

func start(ctx *cli.Context) error {
	handler, err := services.Registry(ctx.String("dbpath"), true)
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
		log.Printf("Registry listening on %q", host)
		errc <- srv.ListenAndServe()
	}()

	log.Printf("Shutting down: %v", <-errc)

	// Shutdown gracefully with a 30s timeout.
	cCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return srv.Shutdown(cCtx)
}

func getClient(ctx *cli.Context) *registryclient.Client {
	cc := registryclient.Config{
		Host:    ctx.String("host"),
		Timeout: 1,
	}

	return registryclient.New(cc)
}

func getID(ctx *cli.Context) (uint64, error) {
	if ctx.Args().Len() != 1 {
		return 0, errors.New("invalid arguments")
	}

	id, err := strconv.ParseUint(ctx.Args().Get(0), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("While parsing id: %v", err)
	}

	return id, nil
}
