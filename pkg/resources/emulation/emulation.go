package emulation

import "code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"

// ResourceType encapuslates our standard resource type, for all storage
// agents.
const ResourceType = "emulation"

// DispatcherConfig configures the storage Dispatcher() call. It is important
// for those that edit this struct, that the call must be edited too, for
// uniformity.
type DispatcherConfig struct {
	WriteConfig  dispatcher.Func
	RemoveConfig dispatcher.Func
	Start        dispatcher.Func
	Stop         dispatcher.Func
	Shutdown     dispatcher.Func
}

// Dispatcher customizes a dispatcher for the purposes of making a storage
// dispatcher with a pre-defined storage-compatible API. See DispatcherConfig for more
// information.
func Dispatcher(dc DispatcherConfig) dispatcher.Table {
	return dispatcher.Table{
		"write_config": {
			RequiredParameters: []string{"vm", "id"},
			Dispatch:           dc.WriteConfig,
		},
		"remove_config": {
			RequiredParameters: []string{"id"},
			Dispatch:           dc.RemoveConfig,
		},
		"start": {
			RequiredParameters: []string{"id"},
			Dispatch:           dc.Start,
		},
		"stop": {
			RequiredParameters: []string{"id"},
			Dispatch:           dc.Stop,
		},
		"shutdown": {
			RequiredParameters: []string{"id"},
			Dispatch:           dc.Shutdown,
		},
	}
}
