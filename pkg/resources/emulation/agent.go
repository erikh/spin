package emulation

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/erikh/spin"
	brokerclient "github.com/erikh/spin/clients/broker"
	"github.com/erikh/spin/pkg/agent"
	"github.com/erikh/spin/pkg/agent/dispatcher"
	"github.com/erikh/spin/pkg/qmp"
	"github.com/erikh/spin/pkg/supervisor"
	"github.com/erikh/spin/pkg/vm"
	"github.com/mitchellh/go-homedir"
)

const (
	systemdUserDir = ".config/systemd/user"
	qemuPath       = "/bin/qemu-system-x86_64"
)

var spinQMPBin = filepath.Join(baseDir(), "spin-qmp")

func baseDir() string {
	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	return filepath.Dir(path)
}

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

			vm := c.Parameter("vm").(*vm.Transient)
			if err := vm.Validate(); err != nil {
				return err
			}

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
			id := c.Parameter("id").(*uint64)
			conn, err := qmp.Dial(ac.monitorPath(*id))
			if err != nil {
				return err
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			return qmp.Shutdown(ctx, conn)
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
