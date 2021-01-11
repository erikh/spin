package storage

import "github.com/erikh/spin/pkg/agent/dispatcher"

// ResourceType encapuslates our standard resource type, for all storage
// agents.
const ResourceType = "storage"

// DispatcherConfig configures the storage Dispatcher() call. It is important
// for those that edit this struct, that the call must be edited too, for
// uniformity.
type DispatcherConfig struct {
	AddVolume    dispatcher.Func
	RemoveVolume dispatcher.Func
	CreateImage  dispatcher.Func
	DeleteImage  dispatcher.Func
	ResizeImage  dispatcher.Func
	MoveImage    dispatcher.Func
	CopyImage    dispatcher.Func
}

// Dispatcher customizes a dispatcher for the purposes of making a storage
// dispatcher with a pre-defined storage-compatible API. See DispatcherConfig for more
// information.
func Dispatcher(dc DispatcherConfig) dispatcher.Table {
	return dispatcher.Table{
		"add_volume": {
			RequiredParameters: dispatcher.ParameterTable{"path": dispatcher.TypeString},
			Dispatch:           dc.AddVolume,
		},
		"remove_volume": {
			RequiredParameters: dispatcher.ParameterTable{"path": dispatcher.TypeString},
			Dispatch:           dc.RemoveVolume,
		},
		"create_image": {
			RequiredParameters: dispatcher.ParameterTable{
				"volume":     dispatcher.TypeString,
				"image":      dispatcher.TypeString,
				"image_size": dispatcher.TypeUint64,
			},
			Dispatch: dc.CreateImage,
		},
		"delete_image": {
			RequiredParameters: dispatcher.ParameterTable{
				"volume": dispatcher.TypeString,
				"image":  dispatcher.TypeString,
			},
			Dispatch: dc.DeleteImage,
		},
		"copy_image": {
			RequiredParameters: dispatcher.ParameterTable{
				"from_volume": dispatcher.TypeString,
				"to_volume":   dispatcher.TypeString,
				"from_image":  dispatcher.TypeString,
				"to_image":    dispatcher.TypeString,
			},
			Dispatch: dc.CopyImage,
		},
		"resize_image": {
			RequiredParameters: dispatcher.ParameterTable{
				"volume":     dispatcher.TypeString,
				"image":      dispatcher.TypeString,
				"image_size": dispatcher.TypeUint64,
			},
			Dispatch: dc.ResizeImage,
		},
		"move_image": {
			RequiredParameters: dispatcher.ParameterTable{
				"image":         dispatcher.TypeString,
				"volume":        dispatcher.TypeString,
				"target_volume": dispatcher.TypeString,
			},
			Dispatch: dc.MoveImage,
		},
	}
}
