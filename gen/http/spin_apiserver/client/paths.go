// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the spin-apiserver service.
//
// Command:
// $ goa gen github.com/erikh/spin/design

package client

import (
	"fmt"
)

// VMCreateSpinApiserverPath returns the URL path to the spin-apiserver service vm_create HTTP endpoint.
func VMCreateSpinApiserverPath() string {
	return "/vm/create"
}

// VMDeleteSpinApiserverPath returns the URL path to the spin-apiserver service vm_delete HTTP endpoint.
func VMDeleteSpinApiserverPath(id uint64) string {
	return fmt.Sprintf("/vm/delete/%v", id)
}

// VMListSpinApiserverPath returns the URL path to the spin-apiserver service vm_list HTTP endpoint.
func VMListSpinApiserverPath() string {
	return "/vm/list"
}

// VMGetSpinApiserverPath returns the URL path to the spin-apiserver service vm_get HTTP endpoint.
func VMGetSpinApiserverPath(id uint64) string {
	return fmt.Sprintf("/vm/get/%v", id)
}

// VMUpdateSpinApiserverPath returns the URL path to the spin-apiserver service vm_update HTTP endpoint.
func VMUpdateSpinApiserverPath(id uint64) string {
	return fmt.Sprintf("/vm/update/%v", id)
}

// ControlStartSpinApiserverPath returns the URL path to the spin-apiserver service control_start HTTP endpoint.
func ControlStartSpinApiserverPath(id uint64) string {
	return fmt.Sprintf("/control/start/%v", id)
}

// ControlStopSpinApiserverPath returns the URL path to the spin-apiserver service control_stop HTTP endpoint.
func ControlStopSpinApiserverPath(id uint64) string {
	return fmt.Sprintf("/control/stop/%v", id)
}

// ControlShutdownSpinApiserverPath returns the URL path to the spin-apiserver service control_shutdown HTTP endpoint.
func ControlShutdownSpinApiserverPath(id uint64) string {
	return fmt.Sprintf("/control/shutdown/%v", id)
}
