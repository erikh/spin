package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

func messageStorageVolumeCreate(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	return getClient(ctx).StorageVolumeCreate(context.Background(), ctx.Args().First())
}

func messageStorageVolumeDelete(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	return getClient(ctx).StorageVolumeDelete(context.Background(), ctx.Args().First())
}

func messageStorageVolumeList(ctx *cli.Context) error {
	list, err := getClient(ctx).StorageVolumeList(context.Background())
	if err != nil {
		return err
	}

	for _, item := range list {
		fmt.Println(item)
	}

	return nil
}
