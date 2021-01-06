package emulation

import (
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
	// FIXME change this
	spinQMPBin = "/home/erikh/bin/spin-qmp"
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
			id := c.Parameter("id").(*uint64)
			vm := c.Parameter("vm").(*spinregistry.UpdatedVM)

			tc, err := vmToTemplateConfig(ac, *id, vm)
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

			fn := configFilename(*id)

			os.Remove(fn)

			f, err := os.Create(fn)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := io.WriteString(f, tpl); err != nil {
				return err
			}

			return ac.Supervisor.Reload()
		},
		RemoveConfig: func(c dispatcher.Command) error {
			id := c.Parameter("id").(*uint64)
			if err := os.Remove(configFilename(*id)); err != nil {
				return err
			}

			return ac.Supervisor.Reload()
		},
		Start: func(c dispatcher.Command) error {
			id := c.Parameter("id").(*uint64)
			return ac.Supervisor.Start(serviceName(*id))
		},
		Stop: func(c dispatcher.Command) error {
			id := c.Parameter("id").(*uint64)
			return ac.Supervisor.Stop(serviceName(*id))
		},
		Shutdown: func(c dispatcher.Command) error {
			return nil
		},
	}
}

// MonitorDir is the directory where the qemu control monitors are kept
var MonitorDir = filepath.Join(spin.ConfigDir(), "monitors")

// AgentConfig is the configuration struct for the constructor.
type AgentConfig struct {
	SystemDir    string
	MonitorDir   string
	ClientConfig brokerclient.Config
	Supervisor   supervisor.Interface
}

// Validate the configuration. Yields error on any, also sets default values if
// required.
func (ac *AgentConfig) Validate() error {
	if ac.SystemDir == "" {
		ac.SystemDir = systemdDir()
	}

	if ac.MonitorDir == "" {
		ac.MonitorDir = MonitorDir
	}

	if ac.Supervisor == nil {
		var err error
		ac.Supervisor, err = supervisor.New()
		if err != nil {
			return err
		}
	}

	return nil
}

func (ac *AgentConfig) monitorPath(id uint64) string {
	return filepath.Join(ac.MonitorDir, fmt.Sprintf("%d", id))
}

// NewAgent creates an agent; the configuration is validated and errors are
// returned if they occur.
func NewAgent(ac AgentConfig) (*agent.Agent, error) {
	if err := ac.Validate(); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(ac.MonitorDir, 0700); err != nil {
		return nil, err
	}

	return agent.New(ac.ClientConfig, ResourceType, Dispatcher(emulationAgent(ac))), nil
}
