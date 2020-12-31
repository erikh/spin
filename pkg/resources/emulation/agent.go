package emulation

import (
	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	"code.hollensbe.org/erikh/spin/pkg/agent"
)

// const systemdUserDir = ".config/systemd/user"
//
// func systemdDir() string {
// 	dir, err := homedir.Dir()
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return filepath.Join(dir, systemdUserDir)
// }

func emulationAgent() DispatcherConfig {
	return DispatcherConfig{}
}

func NewEmulationAgent(cc brokerclient.Config) *agent.Agent {
	return agent.New(cc, ResourceType, Dispatcher(emulationAgent()))
}
