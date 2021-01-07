// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver service
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package spinapiserver

import (
	"context"
)

// Bridge between the outer-facing UIs and the internals
type Service interface {
	// VMCreate implements vm_create.
	VMCreate(context.Context, *CreateVM) (res uint64, err error)
	// VMDelete implements vm_delete.
	VMDelete(context.Context, *VMDeletePayload) (err error)
	// VMList implements vm_list.
	VMList(context.Context) (res []uint64, err error)
	// VMGet implements vm_get.
	VMGet(context.Context, *VMGetPayload) (res *UpdatedVM, err error)
	// VMUpdate implements vm_update.
	VMUpdate(context.Context, *VMUpdatePayload) (err error)
	// ControlStart implements control_start.
	ControlStart(context.Context, *ControlStartPayload) (err error)
	// ControlStop implements control_stop.
	ControlStop(context.Context, *ControlStopPayload) (err error)
	// ControlShutdown implements control_shutdown.
	ControlShutdown(context.Context, *ControlShutdownPayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "spin-apiserver"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [8]string{"vm_create", "vm_delete", "vm_list", "vm_get", "vm_update", "control_start", "control_stop", "control_shutdown"}

// CreateVM is the payload type of the spin-apiserver service vm_create method.
type CreateVM struct {
	// Storage references
	Storage []*Storage
	// Name of VM; does not need to be unique
	Name string
	// CPU count
	Cpus uint
	// Memory (in megabytes)
	Memory uint
}

// VMDeletePayload is the payload type of the spin-apiserver service vm_delete
// method.
type VMDeletePayload struct {
	// ID of VM to delete
	ID uint64
}

// VMGetPayload is the payload type of the spin-apiserver service vm_get method.
type VMGetPayload struct {
	// ID of VM to retrieve
	ID uint64
}

// UpdatedVM is the result type of the spin-apiserver service vm_get method.
type UpdatedVM struct {
	// Image references
	Images []*Image
	// Name of VM; does not need to be unique
	Name string
	// CPU count
	Cpus uint
	// Memory (in megabytes)
	Memory uint
}

// VMUpdatePayload is the payload type of the spin-apiserver service vm_update
// method.
type VMUpdatePayload struct {
	// ID of VM to Update
	ID uint64
	// VM Manifest to Update
	VM *UpdatedVM
}

// ControlStartPayload is the payload type of the spin-apiserver service
// control_start method.
type ControlStartPayload struct {
	// ID of VM to start
	ID uint64
}

// ControlStopPayload is the payload type of the spin-apiserver service
// control_stop method.
type ControlStopPayload struct {
	// ID of VM to stop
	ID uint64
}

// ControlShutdownPayload is the payload type of the spin-apiserver service
// control_shutdown method.
type ControlShutdownPayload struct {
	// ID of VM to shutdown
	ID uint64
}

type Storage struct {
	// Volume name; required if image is not a cdrom
	Volume *string
	// Image filename, no `/` characters
	Image string
	// Image size (in gigabytes); required if image is not a cdrom
	ImageSize *uint
	// Is this image a cdrom?
	Cdrom bool
}

type Image struct {
	// Image path
	Path string
	// Is this a cdrom image?
	Cdrom bool
	// Volume name
	Volume *string
}
