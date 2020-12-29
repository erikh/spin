package storage

import "code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"

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
}

// Dispatcher customizes a dispatcher for the purposes of making a storage
// dispatcher with a pre-defined storage-compatible API. See DispatcherConfig for more
// information.
func Dispatcher(dc DispatcherConfig) dispatcher.Table {
	return dispatcher.Table{
		"add_volume": {
			RequiredParameters: []string{"path"},
			Dispatch:           dc.AddVolume,
		},
		"remove_volume": {
			RequiredParameters: []string{"path"},
			Dispatch:           dc.RemoveVolume,
		},
		"create_image": {
			RequiredParameters: []string{"volume_path", "image_name", "image_size"},
			Dispatch:           dc.CreateImage,
		},
		"delete_image": {
			RequiredParameters: []string{"volume_path", "image_name"},
			Dispatch:           dc.DeleteImage,
		},
		"resize_image": {
			RequiredParameters: []string{"volume_path", "image_name", "image_size"},
			Dispatch:           dc.ResizeImage,
		},
		"move_image": {
			RequiredParameters: []string{"image_name", "volume", "target_volume"},
			Dispatch:           dc.MoveImage,
		},
	}
}
