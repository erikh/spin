package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/erikh/spin/pkg/qmp"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Usage = "Send QMP commands to unix socket monitors."

	app.Commands = []*cli.Command{
		{
			Name:        "raw",
			Usage:       "Accepts JSON commands over STDIN, yields response over STDOUT.",
			Description: "Accepts JSON commands over STDIN, yields response over STDOUT.",
			UsageText:   path.Base(os.Args[0]) + " [monitor socket]",
			Action:      raw,
		},
		{
			Name:        "shutdown",
			Usage:       "Sends a shutdown command to the guest, and waits for it to shutdown.",
			Description: "Sends a shutdown command to the guest, and waits for it to shutdown.",
			UsageText:   path.Base(os.Args[0]) + " [monitor socket]",
			Flags: []cli.Flag{
				&cli.DurationFlag{
					Name:    "wait",
					Aliases: []string{"w"},
					Usage:   "How long to wait before giving up",
					Value:   10 * time.Second,
				},
			},
			Action: shutdown,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func shutdown(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("please provide the monitor socket (see --help)")
	}

	conn, err := qmp.Dial(ctx.Args().First())
	if err != nil {
		return err
	}
	defer conn.Close()

	cCtx, cancel := context.WithTimeout(context.Background(), ctx.Duration("wait"))
	defer cancel()
	return qmp.Shutdown(cCtx, conn)
}

func raw(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("please provide the monitor socket (see --help)")
	}

	conn, err := qmp.Dial(ctx.Args().First())
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := io.Copy(conn, os.Stdin); err != nil {
		return err
	}

	_, err = io.Copy(os.Stdout, conn)
	return err
}
