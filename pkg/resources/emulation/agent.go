package emulation

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"code.hollensbe.org/erikh/spin"
	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	"code.hollensbe.org/erikh/spin/pkg/agent"
	"code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"
	"code.hollensbe.org/erikh/spin/pkg/supervisor"
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

func emulationAgent(ac AgentConfig) DispatcherConfig {
	serviceName := func(id uint64) string {
		return fmt.Sprintf("spin-%d.service", id)
	}

	configFilename := func(id uint64) string {
		return filepath.Join(ac.SystemDir, serviceName(id))
	}

	return DispatcherConfig{
		WriteConfig: func(c dispatcher.Command) error {
			vm, err := commandToVM(c.Parameters["vm"].(map[string]interface{}))
			if err != nil {
				return err
			}

			id := uint64(c.Parameters["id"].(float64))

			tc, err := vmToTemplateConfig(ac, id, vm)
			if err != nil {
				return err
			}

			tpl, err := runTemplate(tc)
			if err != nil {
				return err
			}

			if err := os.MkdirAll(ac.SystemDir, 0700); err != nil && !os.IsExist(err) {
				return err
			}

			fn := configFilename(id)

			os.Remove(fn)

			f, err := os.Create(fn)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := io.WriteString(f, tpl); err != nil {
				return err
			}

			s, err := supervisor.New()
			if err != nil {
				return err
			}

			return s.Reload()
		},
		RemoveConfig: func(c dispatcher.Command) error {
			id := uint64(c.Parameters["id"].(float64))
			if err := os.Remove(configFilename(id)); err != nil {
				return err
			}

			s, err := supervisor.New()
			if err != nil {
				return err
			}

			return s.Reload()
		},
		Start: func(c dispatcher.Command) error {
			s, err := supervisor.New()
			if err != nil {
				return err
			}

			id := uint64(c.Parameters["id"].(float64))
			return s.Start(serviceName(id))
		},
		Stop: func(c dispatcher.Command) error {
			s, err := supervisor.New()
			if err != nil {
				return err
			}
			id := uint64(c.Parameters["id"].(float64))
			return s.Stop(serviceName(id))
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

// MonitorDir is the directory where the qemu control monitors are kept
var MonitorDir = filepath.Join(spin.ConfigDir(), "monitors")

type AgentConfig struct {
	SystemDir    string
	MonitorDir   string
	ClientConfig brokerclient.Config
}

func (ac *AgentConfig) Validate() error {
	if ac.SystemDir == "" {
		ac.SystemDir = systemdDir()
	}

	if ac.MonitorDir == "" {
		ac.MonitorDir = MonitorDir
	}

	return nil
}

func (ac *AgentConfig) monitorPath(id uint64) string {
	return filepath.Join(ac.MonitorDir, fmt.Sprintf("%d", id))
}

func NewAgent(ac AgentConfig) (*agent.Agent, error) {
	if err := ac.Validate(); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(ac.MonitorDir, 0700); err != nil {
		return nil, err
	}

	return agent.New(ac.ClientConfig, ResourceType, Dispatcher(emulationAgent(ac))), nil
}
