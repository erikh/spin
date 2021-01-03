package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"time"

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
			Name: "shutdown",
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

func getConn(socket string) (net.Conn, error) {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, err
	}

	if err := json.NewEncoder(conn).Encode(map[string]interface{}{"execute": "qmp_capabilities"}); err != nil {
		conn.Close()
		return nil, err
	}

	obj := map[string]interface{}{}
	if err := json.NewDecoder(conn).Decode(&obj); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func shutdown(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("please provide the monitor socket (see --help)")
	}

	conn, err := getConn(ctx.Args().First())
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := json.NewEncoder(conn).Encode(map[string]string{"execute": "system_powerdown"}); err != nil {
		return err
	}

	go func() {
		time.Sleep(ctx.Duration("wait"))
		conn.Close()
	}()

	for {
		obj := map[string]interface{}{}
		if err := json.NewDecoder(conn).Decode(&obj); err != nil {
			return err
		}
		fmt.Println(obj)
		if event, ok := obj["event"]; ok && event == "SHUTDOWN" {
			return nil
		}
	}
}

func raw(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("please provide the monitor socket (see --help)")
	}

	conn, err := getConn(ctx.Args().First())
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
