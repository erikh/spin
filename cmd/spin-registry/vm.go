package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	spinregistry "github.com/erikh/spin/gen/spin_registry"
	"github.com/urfave/cli/v2"
)

func messageVMCreate(ctx *cli.Context) error {
	var vm spinregistry.UpdatedVM

	if err := json.NewDecoder(os.Stdin).Decode(&vm); err != nil {
		return fmt.Errorf("Error decoding JSON document: %v", err)
	}

	id, err := getClient(ctx).VMCreate(context.Background(), &vm)
	if err != nil {
		return err
	}

	fmt.Println(id)
	return nil
}

func messageVMUpdate(ctx *cli.Context) error {
	id, err := getID(ctx)
	if err != nil {
		return err
	}

	var vm spinregistry.UpdatedVM

	if err := json.NewDecoder(os.Stdin).Decode(&vm); err != nil {
		return fmt.Errorf("Error decoding JSON document: %v", err)
	}

	return getClient(ctx).VMUpdate(context.Background(), id, &vm)
}

func messageVMDelete(ctx *cli.Context) error {
	id, err := getID(ctx)
	if err != nil {
		return err
	}

	return getClient(ctx).VMDelete(context.Background(), id)
}

func messageVMGet(ctx *cli.Context) error {
	id, err := getID(ctx)
	if err != nil {
		return err
	}

	vm, err := getClient(ctx).VMGet(context.Background(), id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(vm)
}

func messageVMList(ctx *cli.Context) error {
	ids, err := getClient(ctx).VMList(context.Background())
	if err != nil {
		return err
	}

	for _, id := range ids {
		fmt.Println(id)
	}

	return nil
}
