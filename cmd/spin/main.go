package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/rakyll/statik/fs"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli/v2"

	"github.com/erikh/spin/clients/api"
	"github.com/erikh/spin/pkg/vm"

	_ "github.com/erikh/spin/statik"
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
			Name:      "serve",
			Usage:     "Serve consoles for your VMs",
			ArgsUsage: " ",
			Action:    serve,
		},
		{
			Name:      "view",
			Usage:     "View a VM's screen in your browser",
			ArgsUsage: "[id]",
			Action:    view,
		},
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
					Name:    "image",
					Aliases: []string{"i"},
					Usage:   "Operations on VM images",
					Subcommands: []*cli.Command{
						{
							Name:      "list",
							Usage:     "list images for the vm",
							ArgsUsage: "[id]",
							Action:    vmImageList,
						},
						{
							Name:      "detach",
							Usage:     "detach an image from a VM by VM ID and image index",
							ArgsUsage: "[id] [index]",
							Action:    vmImageDetach,
						},
					},
				},
				{
					Name:      "list",
					Usage:     "List all VMs by ID + Name",
					ArgsUsage: " ",
					Action:    list,
				},
				{
					Name:      "get",
					Usage:     "Retrieve a VM by ID",
					ArgsUsage: "[id]",
					Action:    get,
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
						&cli.StringSliceFlag{
							Name:    "port",
							Aliases: []string{"p"},
							Usage:   "Syntax: <guestport>:<hostaddr>:<hostport>",
						},
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
		Timeout: 60,
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

	v := &vm.Create{
		Core: vm.Core{
			Name:   name,
			CPUs:   ctx.Uint("cpus"),
			Memory: ctx.Uint("memory"),
			Ports:  map[uint]string{},
		},
		Storage: []vm.Storage{
			{
				Volume:    stringPtr(volume),
				Image:     name,
				ImageSize: uintPtr(ctx.Uint("image-size")),
			},
		},
	}

	if ctx.String("cdrom") != "" {
		v.Storage = append(v.Storage, vm.Storage{
			CDROM: true,
			Image: ctx.String("cdrom"),
		})
	}

	for _, port := range ctx.StringSlice("port") {
		parts := strings.SplitN(port, ":", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid port specification: %v", port)
		}

		guestPort, err := strconv.ParseUint(parts[0], 10, 16)
		if err != nil {
			return err
		}

		v.Ports[uint(guestPort)] = parts[1]
	}

	id, err := getClient(ctx).VMCreate(context.Background(), v)
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

func get(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	id, err := strconv.ParseUint(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}

	ret, err := getClient(ctx).VMGet(context.Background(), id)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(ret)
}

func vmImageList(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	id, err := strconv.ParseUint(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}

	ret, err := getClient(ctx).VMGet(context.Background(), id)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	fmt.Fprintf(w, "INDEX\tPATH\tCDROM?\n")

	for x, image := range ret.Images {
		fmt.Fprintf(w, "%d\t%s\t%v\n", x, image.Path, image.CDROM)
	}

	return w.Flush()
}

func vmImageDetach(ctx *cli.Context) error {
	if ctx.Args().Len() != 2 {
		return errors.New("invalid arguments; see --help")
	}

	id, err := strconv.ParseUint(ctx.Args().Get(0), 10, 64)
	if err != nil {
		return err
	}

	index, err := strconv.ParseUint(ctx.Args().Get(1), 10, 64)
	if err != nil {
		return err
	}

	ret, err := getClient(ctx).VMGet(context.Background(), id)
	if err != nil {
		return err
	}

	if uint64(len(ret.Images)) <= index {
		return errors.New("invalid index")
	}

	ret.Images = append(ret.Images[:index], ret.Images[index+1:]...)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	return getClient(ctx).VMUpdate(context.Background(), id, ret)
}

func serve(ctx *cli.Context) error {
	staticFilesystem, err := fs.New()
	if err != nil {
		return err
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.RequestURI)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		path := u.Path

		if path == "/" {
			path = "/index.html"
		}

		f, err := staticFilesystem.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, err.Error(), 404)
			} else {
				http.Error(w, err.Error(), 500)
			}

			return
		}

		w.Header().Add("content-type", mime.TypeByExtension(filepath.Ext(path)))

		if _, err := io.Copy(w, f); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	return http.ListenAndServe("localhost:8083", mux)
}

func view(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	id, err := strconv.ParseUint(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}

	return open.Start(fmt.Sprintf("http://localhost:8083?id=%d", id))
}
