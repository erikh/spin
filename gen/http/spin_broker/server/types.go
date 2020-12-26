// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker HTTP server types
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package server

import (
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "spin-broker" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Resource name
	Resource *string `form:"resource,omitempty" json:"resource,omitempty" xml:"resource,omitempty"`
	// Action name
	Action *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
	// Action parameters
	Parameters map[string]string `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// CompleteRequestBody is the type of the "spin-broker" service "complete"
// endpoint HTTP request body.
type CompleteRequestBody struct {
	// Command ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// status of work
	Status *bool `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// reason of success/failure
	StatusReason *string `form:"status_reason,omitempty" json:"status_reason,omitempty" xml:"status_reason,omitempty"`
}

// StatusResponseBody is the type of the "spin-broker" service "status"
// endpoint HTTP response body.
type StatusResponseBody struct {
	// Pass/Fail status
	Status bool `form:"status" json:"status" xml:"status"`
	// Failure reason (if any)
	Reason *string `form:"reason,omitempty" json:"reason,omitempty" xml:"reason,omitempty"`
}

// NextResponseBody is the type of the "spin-broker" service "next" endpoint
// HTTP response body.
type NextResponseBody struct {
	// Command ID
	UUID string `form:"uuid" json:"uuid" xml:"uuid"`
	// resource type
	Resource string `form:"resource" json:"resource" xml:"resource"`
	// action name
	Action string `form:"action" json:"action" xml:"action"`
	// parameters for action
	Parameters map[string]string `form:"parameters,omitempty" json:"parameters,omitempty" xml:"parameters,omitempty"`
}

// NewStatusResponseBody builds the HTTP response body from the result of the
// "status" endpoint of the "spin-broker" service.
func NewStatusResponseBody(res *spinbroker.StatusResult) *StatusResponseBody {
	body := &StatusResponseBody{
		Status: res.Status,
		Reason: res.Reason,
	}
	return body
}

// NewNextResponseBody builds the HTTP response body from the result of the
// "next" endpoint of the "spin-broker" service.
func NewNextResponseBody(res *spinbroker.NextResult) *NextResponseBody {
	body := &NextResponseBody{
		UUID:     res.UUID,
		Resource: res.Resource,
		Action:   res.Action,
	}
	if res.Parameters != nil {
		body.Parameters = make(map[string]string, len(res.Parameters))
		for key, val := range res.Parameters {
			tk := key
			tv := val
			body.Parameters[tk] = tv
		}
	}
	return body
}

// NewAddPayload builds a spin-broker service add endpoint payload.
func NewAddPayload(body *AddRequestBody, id string) *spinbroker.AddPayload {
	v := &spinbroker.AddPayload{
		Resource: *body.Resource,
		Action:   *body.Action,
	}
	if body.Parameters != nil {
		v.Parameters = make(map[string]string, len(body.Parameters))
		for key, val := range body.Parameters {
			tk := key
			tv := val
			v.Parameters[tk] = tv
		}
	}
	v.ID = id

	return v
}

// NewEnqueuePayload builds a spin-broker service enqueue endpoint payload.
func NewEnqueuePayload(id string) *spinbroker.EnqueuePayload {
	v := &spinbroker.EnqueuePayload{}
	v.ID = id

	return v
}

// NewStatusPayload builds a spin-broker service status endpoint payload.
func NewStatusPayload(id string) *spinbroker.StatusPayload {
	v := &spinbroker.StatusPayload{}
	v.ID = id

	return v
}

// NewNextPayload builds a spin-broker service next endpoint payload.
func NewNextPayload(resource string) *spinbroker.NextPayload {
	v := &spinbroker.NextPayload{}
	v.Resource = resource

	return v
}

// NewCompletePayload builds a spin-broker service complete endpoint payload.
func NewCompletePayload(body *CompleteRequestBody) *spinbroker.CompletePayload {
	v := &spinbroker.CompletePayload{
		ID:           *body.ID,
		Status:       *body.Status,
		StatusReason: body.StatusReason,
	}

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Resource == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("resource", "body"))
	}
	if body.Action == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action", "body"))
	}
	return
}

// ValidateCompleteRequestBody runs the validations defined on
// CompleteRequestBody
func ValidateCompleteRequestBody(body *CompleteRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	return
}
