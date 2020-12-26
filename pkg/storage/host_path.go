package storage

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"code.hollensbe.org/erikh/spin/gen/http/spin_broker/client"
	"code.hollensbe.org/erikh/spin/pkg/agent"
	"code.hollensbe.org/erikh/spin/pkg/broker"
)

func hostPathDispatcher(client *client.Client) broker.Dispatcher {
	return broker.Dispatcher{
		"add_volume": {
			RequiredParameters: []string{"path"},
			Dispatch: func(c broker.Command) error {
				if fi, err := os.Stat(c.Parameters["path"]); err != nil {
					return err
				} else if !fi.IsDir() {
					return errors.New("path is not a directory")
				}

				return nil
			},
		},
		"remove_volume": {
			RequiredParameters: []string{"path"},
			Dispatch: func(c broker.Command) error {
				if fi, err := os.Stat(c.Parameters["path"]); err != nil {
					return err
				} else if !fi.IsDir() {
					return errors.New("path is not a directory")
				}

				return os.RemoveAll(c.Parameters["path"])
			},
		},
		"create_image": {
			RequiredParameters: []string{"volume_path", "image_name", "image_size"},
			Dispatch: func(c broker.Command) error {
				return exec.Command("truncate", "-s", fmt.Sprintf("%sG", c.Parameters["image_size"]), filepath.Join(c.Parameters["volume_path"], c.Parameters["image_name"])).Run()
			},
		},
		"delete_image": {
			RequiredParameters: []string{"volume_path", "image_name"},
			Dispatch: func(c broker.Command) error {
				return os.Remove(filepath.Join(c.Parameters["volume_path"], c.Parameters["image_name"]))
			},
		},
		"resize_image": {
			RequiredParameters: []string{"volume_path", "image_name", "image_size"},
			Dispatch: func(c broker.Command) error {
				return nil
			},
		},
		"move_image": {
			RequiredParameters: []string{"image_name", "volume", "target_volume"},
			Dispatch: func(c broker.Command) error {
				return nil
			},
		},
	}
}

func NewHostPathAgent(cc agent.ClientConfig) *agent.Agent {
	return agent.New(cc, ResourceType, hostPathDispatcher(cc.MakeClient()))
}
