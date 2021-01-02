package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Usage = "Send QMP commands to unix socket monitors."
	app.Description = "Accepts JSON commands over STDIN, yields response over STDOUT."
	app.UsageText = path.Base(os.Args[0]) + " [monitor socket]"

	app.Action = start

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func start(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("please provide the monitor socket (see --help)")
	}

	socket := ctx.Args().First()

	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	defer conn.Close()

	obj := map[string]interface{}{}

	if err := json.NewEncoder(conn).Encode(map[string]interface{}{"execute": "qmp_capabilities"}); err != nil {
		return err
	}

	if err := json.NewDecoder(conn).Decode(&obj); err != nil {
		return err
	}

	if _, err := io.Copy(conn, os.Stdin); err != nil {
		return err
	}

	conn.Write([]byte("\n"))
	_, err = io.Copy(os.Stdout, conn)
	return err
}
