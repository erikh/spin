// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker service
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package spinbroker

import (
	"context"
	"encoding/json"

	goa "goa.design/goa/v3/pkg"
)

// The message broker for the other services
type Service interface {
	// Create a new package; a collection of items to join into the queue
	// simultaneously
	New(context.Context) (res string, err error)
	// Add a command to the package
	Add(context.Context, *AddPayload) (res string, err error)
	// Enqueue the package into the various resource queues
	Enqueue(context.Context, *EnqueuePayload) (res []string, err error)
	// Get the status for a package
	Status(context.Context, *StatusPayload) (res *StatusResult, err error)
	// Get the next command for a given resource
	Next(context.Context, *NextPayload) (res *NextResult, err error)
	// Mark a command as completed with a result status
	Complete(context.Context, *CompletePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "spin-broker"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"new", "add", "enqueue", "status", "next", "complete"}

// AddPayload is the payload type of the spin-broker service add method.
type AddPayload struct {
	// Package ID
	ID string
	// Resource name
	Resource string
	// Action name
	Action string
	// Action parameters
	Parameters map[string]interface{}
}

// EnqueuePayload is the payload type of the spin-broker service enqueue method.
type EnqueuePayload struct {
	// Package ID
	ID string
}

// StatusPayload is the payload type of the spin-broker service status method.
type StatusPayload struct {
	// Package ID
	ID string
}

// StatusResult is the result type of the spin-broker service status method.
type StatusResult struct {
	// Pass/Fail status
	Status bool
	// Failure reason (if any)
	Reason *string
	// Failure causer as UUID (if any)
	Causer *string
}

// NextPayload is the payload type of the spin-broker service next method.
type NextPayload struct {
	// resource type
	Resource string
}

// NextResult is the result type of the spin-broker service next method.
type NextResult struct {
	// Command ID
	UUID string
	// resource type
	Resource string
	// action name
	Action string
	// Action parameters
	Parameters map[string]json.RawMessage
}

// CompletePayload is the payload type of the spin-broker service complete
// method.
type CompletePayload struct {
	// Command ID
	ID string
	// status of work
	Status bool
	// reason of success/failure
	StatusReason *string
}

// MakeRecordNotFound builds a goa.ServiceError from an error.
func MakeRecordNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "record_not_found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
