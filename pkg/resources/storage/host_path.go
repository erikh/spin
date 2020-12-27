package storage

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	"code.hollensbe.org/erikh/spin/pkg/agent"
	"code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"
)

func hostPathDispatcher(basePath string) dispatcher.Dispatcher {
	bp := func(strs ...string) (string, error) {
		if len(strs) == 0 {
			return "", errors.New("empty path")
		}

		strs = append([]string{basePath}, strs...)

		full := filepath.Join(strs...)
		rel, err := filepath.Rel(basePath, full)
		if err != nil {
			return "", err
		}

		if filepath.Clean(rel) == filepath.Clean(basePath) || strings.Contains(rel, "..") {
			return "", errors.New("path falls below root")
		}

		return full, nil
	}

	return dispatcher.Dispatcher{
		"add_volume": {
			RequiredParameters: []string{"path"},
			Dispatch: func(c dispatcher.Command) error {
				path, err := bp(c.Parameters["path"])
				if err != nil {
					return err
				}

				if fi, err := os.Stat(path); os.IsNotExist(err) {
					return os.MkdirAll(path, 0700)
				} else if err != nil {
					return err
				} else if !fi.IsDir() {
					return errors.New("path is not a directory")
				}

				return errors.New("already exists")
			},
		},
		"remove_volume": {
			RequiredParameters: []string{"path"},
			Dispatch: func(c dispatcher.Command) error {
				path, err := bp(c.Parameters["path"])
				if err != nil {
					return err
				}

				if fi, err := os.Stat(path); err != nil {
					return err
				} else if !fi.IsDir() {
					return errors.New("path is not a directory")
				}

				return os.RemoveAll(path)
			},
		},
		"create_image": {
			RequiredParameters: []string{"volume_path", "image_name", "image_size"},
			Dispatch: func(c dispatcher.Command) error {
				path, err := bp(c.Parameters["volume_path"], c.Parameters["image_name"])
				if err != nil {
					return err
				}

				return exec.Command("qemu-img", "create", "-f", "raw", path, fmt.Sprintf("%sG", c.Parameters["image_size"])).Run()
			},
		},
		"delete_image": {
			RequiredParameters: []string{"volume_path", "image_name"},
			Dispatch: func(c dispatcher.Command) error {
				path, err := bp(c.Parameters["volume_path"], c.Parameters["image_name"])
				if err != nil {
					return err
				}

				return os.Remove(path)
			},
		},
		"resize_image": {
			RequiredParameters: []string{"volume_path", "image_name", "image_size"},
			Dispatch: func(c dispatcher.Command) error {
				return nil
			},
		},
		"move_image": {
			RequiredParameters: []string{"image_name", "volume", "target_volume"},
			Dispatch: func(c dispatcher.Command) error {
				return nil
			},
		},
	}
}

func NewHostPathAgent(cc brokerclient.Config) *agent.Agent {
	return agent.New(cc, ResourceType, hostPathDispatcher("/tmp/host-path-test"))
}
