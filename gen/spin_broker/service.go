// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker service
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package spinbroker

import (
	"context"
)

// The message broker for the other services
type Service interface {
	// New implements new.
	New(context.Context) (res string, err error)
	// Add implements add.
	Add(context.Context, *AddPayload) (res string, err error)
	// Enqueue implements enqueue.
	Enqueue(context.Context, *EnqueuePayload) (res []string, err error)
	// Status implements status.
	Status(context.Context, *StatusPayload) (res *StatusResult, err error)
	// Next implements next.
	Next(context.Context, *NextPayload) (res *NextResult, err error)
	// Complete implements complete.
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
	// queue ID
	ID string
	// Resource name
	Resource string
	// Action name
	Action string
	// Action parameters
	Parameters []string
}

// EnqueuePayload is the payload type of the spin-broker service enqueue method.
type EnqueuePayload struct {
	// queue ID
	ID string
}

// StatusPayload is the payload type of the spin-broker service status method.
type StatusPayload struct {
	// queue ID
	ID string
}

// StatusResult is the result type of the spin-broker service status method.
type StatusResult struct {
	// Pass/Fail status
	Status bool
}

// NextPayload is the payload type of the spin-broker service next method.
type NextPayload struct {
	// resource type
	Resource string
}

// NextResult is the result type of the spin-broker service next method.
type NextResult struct {
	// resource type
	Resource string
	// action name
	Action string
	// parameters for action
	Parameters []string
}

// CompletePayload is the payload type of the spin-broker service complete
// method.
type CompletePayload struct {
	// queue ID
	ID string
	// status of work
	Status bool
	// reason of success/failure
	StatusReason *string
}
