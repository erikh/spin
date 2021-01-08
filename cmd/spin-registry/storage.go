package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	spinregistry "github.com/erikh/spin/gen/spin_registry"
	"github.com/urfave/cli/v2"
)

func messageStorageVolumeCreate(ctx *cli.Context) error {
	if ctx.Args().Len() != 2 {
		return errors.New("invalid arguments; see --help")
	}

	return getClient(ctx).StorageVolumeCreate(context.Background(), ctx.Args().Get(0), ctx.Args().Get(1))
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

func messageStorageImageList(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return errors.New("invalid arguments; see --help")
	}

	list, err := getClient(ctx).StorageImageList(context.Background(), ctx.Args().First())
	if err != nil {
		return err
	}

	for _, item := range list {
		fmt.Println(item)
	}

	return nil
}

func messageStorageImageGet(ctx *cli.Context) error {
	if ctx.Args().Len() != 2 {
		return errors.New("invalid arguments; see --help")
	}

	image, err := getClient(ctx).StorageImageGet(context.Background(), ctx.Args().Get(0), ctx.Args().Get(1))
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(image)
}

func messageStorageImageCreate(ctx *cli.Context) error {
	s := &spinregistry.Storage{}

	if err := json.NewDecoder(os.Stdin).Decode(s); err != nil {
		return err
	}

	img, err := getClient(ctx).StorageImageCreate(context.Background(), s)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(img)
}

func messageStorageImageDelete(ctx *cli.Context) error {
	if ctx.Args().Len() != 2 {
		return errors.New("invalid arguments; see --help")
	}

	return getClient(ctx).StorageImageDelete(context.Background(), ctx.Args().Get(0), ctx.Args().Get(1))
}
