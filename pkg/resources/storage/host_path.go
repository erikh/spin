package storage

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	brokerclient "github.com/erikh/spin/clients/broker"
	"github.com/erikh/spin/pkg/agent"
	"github.com/erikh/spin/pkg/agent/dispatcher"
	"golang.org/x/sys/unix"
)

func validateVolumePath(p string) error {
	if strings.Contains(p, "/") || strings.Contains(p, "..") {
		return errors.New("volumes may not contain path components")
	}

	if p == "" || p == "." {
		return errors.New("volumes may not refer to the root directory")
	}

	return nil
}

func hostPathDispatcher(basePath string) DispatcherConfig {
	bp := func(strs ...interface{}) (string, error) {
		if len(strs) == 0 {
			return "", errors.New("empty path")
		}

		res := []string{basePath}

		for _, str := range strs {
			if str, ok := str.(*string); ok && str != nil {
				if err := validateVolumePath(*str); err != nil {
					return "", err
				}

				res = append(res, *str)
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
			path, err := bp(c.Parameter("path").(*string))
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
			path, err := bp(c.Parameter("path").(*string))
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
			path, err := bp(c.Parameter("volume").(*string), c.Parameter("image").(*string))
			if err != nil {
				return err
			}

			size := c.Parameter("image_size").(*uint64)
			if size == nil || *size == 0 {
				return errors.New("image size must not be empty")
			}

			// FIXME add a debug trap for these shell commands later
			cmd := exec.Command("qemu-img", "create", "-f", "raw", path, fmt.Sprintf("%dG", *size))
			out, err := cmd.CombinedOutput()
			if err != nil {
				return errors.New(string(out))
			}

			return nil
		},
		DeleteImage: func(c dispatcher.Command) error {
			path, err := bp(c.Parameter("volume").(*string), c.Parameter("image").(*string))
			if err != nil {
				return err
			}

			return os.Remove(path)
		},
		CopyImage: func(c dispatcher.Command) error {
			fromPath, err := bp(c.Parameter("from_volume").(*string), c.Parameter("from_image").(*string))
			if err != nil {
				return err
			}

			toPath, err := bp(c.Parameter("to_volume").(*string), c.Parameter("to_image").(*string))
			if err != nil {
				return err
			}
			fromStat, err := os.Stat(fromPath)
			if err != nil {
				return fmt.Errorf("while locating original image: %v", err)
			}

			if _, err := os.Stat(toPath); err == nil {
				return errors.New("target file exists")
			}

			to, err := os.Create(toPath)
			if err != nil {
				return fmt.Errorf("while opening target file: %v", err)
			}
			defer to.Close()

			from, err := os.Open(fromPath)
			if err != nil {
				return fmt.Errorf("while reading from file: %v", err)
			}
			defer from.Close()

			for toWrite := fromStat.Size(); toWrite > 0; {
				written, err := unix.Sendfile(int(to.Fd()), int(from.Fd()), nil, int(toWrite))
				if err != nil {
					if err == io.EOF {
						return nil
					}
					return err
				}

				toWrite -= int64(written)
			}

			return nil
		},
		ResizeImage: func(c dispatcher.Command) error {
			return nil
		},
		MoveImage: func(c dispatcher.Command) error {
			return nil
		},
	}
}

// NewHostPathAgent creates a new host-path agent. rootPath is the root of all
// volumes; and the client configuration is used to talk to the broker.
func NewHostPathAgent(rootPath string, cc brokerclient.Config) *agent.Agent {
	return agent.New(cc, ResourceType, Dispatcher(hostPathDispatcher(rootPath)))
}
