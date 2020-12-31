package emulation

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	"code.hollensbe.org/erikh/spin/pkg/agent"
	"code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"
	"github.com/mitchellh/go-homedir"
)

const (
	systemdUserDir = ".config/systemd/user"
	qemuPath       = "/bin/qemu-system-x86_64"
)

func systemdDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(dir, systemdUserDir)
}

func emulationAgent(dir string) DispatcherConfig {
	if dir == "" {
		dir = systemdDir()
	}

	return DispatcherConfig{
		WriteConfig: func(c dispatcher.Command) error {
			vm, err := commandToVM(c.Parameters["vm"].(map[string]interface{}))
			if err != nil {
				return err
			}

			id := uint64(c.Parameters["id"].(float64))

			tc, err := vmToTemplateConfig(id, vm)
			if err != nil {
				return err
			}

			tpl, err := runTemplate(tc)
			if err != nil {
				return err
			}

			name := fmt.Sprintf("spin-%d.service", id)

			if err := os.MkdirAll(dir, 0700); err != nil {
				return err
			}

			fn := filepath.Join(dir, name)
			os.Remove(fn)

			f, err := os.Create(fn)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := io.WriteString(f, tpl); err != nil {
				return err
			}

			// FIXME reload systemd
			return nil
		},
		RemoveConfig: func(c dispatcher.Command) error {
			return nil
		},
		Start: func(c dispatcher.Command) error {
			return nil
		},
		Stop: func(c dispatcher.Command) error {
			return nil
		},
		Shutdown: func(c dispatcher.Command) error {
			return nil
		},
	}
}

func commandToVM(vm map[string]interface{}) (*spinregistry.VM, error) {
	content, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}

	var ret spinregistry.VM

	if err := json.Unmarshal(content, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func NewAgent(bc brokerclient.Config, dir string) *agent.Agent {
	return agent.New(bc, ResourceType, Dispatcher(emulationAgent(dir)))
}
