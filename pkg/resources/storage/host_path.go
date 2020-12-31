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

func hostPathDispatcher(basePath string) DispatcherConfig {
	bp := func(strs ...interface{}) (string, error) {
		if len(strs) == 0 {
			return "", errors.New("empty path")
		}

		res := []string{basePath}

		for _, str := range strs {
			if str, ok := str.(string); ok {
				res = append(res, str)
			} else {
				return "", errors.New("path must be a string")
			}
		}

		full := filepath.Join(res...)
		rel, err := filepath.Rel(basePath, full)
		if err != nil {
			return "", err
		}

		if filepath.Clean(rel) == filepath.Clean(basePath) || strings.Contains(rel, "..") {
			return "", errors.New("path falls below root")
		}

		return full, nil
	}

	return DispatcherConfig{
		AddVolume: func(c dispatcher.Command) error {
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
		RemoveVolume: func(c dispatcher.Command) error {
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
		CreateImage: func(c dispatcher.Command) error {
			path, err := bp(c.Parameters["volume_path"], c.Parameters["image_name"])
			if err != nil {
				return err
			}

			// FIXME add a debug trap for these shell commands later
			cmd := exec.Command("qemu-img", "create", "-f", "raw", path, fmt.Sprintf("%.0fG", c.Parameters["image_size"]))
			return cmd.Run()
		},
		DeleteImage: func(c dispatcher.Command) error {
			path, err := bp(c.Parameters["volume_path"], c.Parameters["image_name"])
			if err != nil {
				return err
			}

			return os.Remove(path)
		},
		ResizeImage: func(c dispatcher.Command) error {
			return nil
		},
		MoveImage: func(c dispatcher.Command) error {
			return nil
		},
	}
}

func NewHostPathAgent(rootPath string, cc brokerclient.Config) *agent.Agent {
	return agent.New(cc, ResourceType, Dispatcher(hostPathDispatcher(rootPath)))
}
