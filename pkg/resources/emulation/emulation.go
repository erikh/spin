package emulation

import (
	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	"code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"
)

// ResourceType encapuslates our standard resource type, for all storage
// agents.
const ResourceType = "emulation"

// DispatcherConfig configures the emulation Dispatcher() call. It is important
// for those that edit this struct, that the call must be edited too, for
// uniformity.
type DispatcherConfig struct {
	WriteConfig  dispatcher.Func
	RemoveConfig dispatcher.Func
	Start        dispatcher.Func
	Stop         dispatcher.Func
	Shutdown     dispatcher.Func
}

// Dispatcher customizes a dispatcher for the purposes of making a emulation
// dispatcher with a pre-defined emulation-compatible API. See DispatcherConfig for more
// information.
func Dispatcher(dc DispatcherConfig) dispatcher.Table {
	return dispatcher.Table{
		"write_config": {
			RequiredParameters: dispatcher.ParameterTable{
				"vm": func() interface{} { return &spinregistry.VM{} },
				"id": dispatcher.TypeUint64,
			},
			Dispatch: dc.WriteConfig,
		},
		"remove_config": {
			RequiredParameters: dispatcher.ParameterTable{"id": dispatcher.TypeUint64},
			Dispatch:           dc.RemoveConfig,
		},
		"start": {
			RequiredParameters: dispatcher.ParameterTable{"id": dispatcher.TypeUint64},
			Dispatch:           dc.Start,
		},
		"stop": {
			RequiredParameters: dispatcher.ParameterTable{"id": dispatcher.TypeUint64},
			Dispatch:           dc.Stop,
		},
		"shutdown": {
			RequiredParameters: dispatcher.ParameterTable{"id": dispatcher.TypeUint64},
			Dispatch:           dc.Shutdown,
		},
	}
}
