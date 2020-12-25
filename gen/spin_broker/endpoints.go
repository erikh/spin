// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker endpoints
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package spinbroker

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "spin-broker" service endpoints.
type Endpoints struct {
	New      goa.Endpoint
	Add      goa.Endpoint
	Enqueue  goa.Endpoint
	Enqueued goa.Endpoint
	Status   goa.Endpoint
	Next     goa.Endpoint
	Complete goa.Endpoint
}

// NewEndpoints wraps the methods of the "spin-broker" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		New:      NewNewEndpoint(s),
		Add:      NewAddEndpoint(s),
		Enqueue:  NewEnqueueEndpoint(s),
		Enqueued: NewEnqueuedEndpoint(s),
		Status:   NewStatusEndpoint(s),
		Next:     NewNextEndpoint(s),
		Complete: NewCompleteEndpoint(s),
	}
}

// Use applies the given middleware to all the "spin-broker" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.New = m(e.New)
	e.Add = m(e.Add)
	e.Enqueue = m(e.Enqueue)
	e.Enqueued = m(e.Enqueued)
	e.Status = m(e.Status)
	e.Next = m(e.Next)
	e.Complete = m(e.Complete)
}

// NewNewEndpoint returns an endpoint function that calls the method "new" of
// service "spin-broker".
func NewNewEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.New(ctx)
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "spin-broker".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
		return s.Add(ctx, p)
	}
}

// NewEnqueueEndpoint returns an endpoint function that calls the method
// "enqueue" of service "spin-broker".
func NewEnqueueEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*EnqueuePayload)
		return s.Enqueue(ctx, p)
	}
}

// NewEnqueuedEndpoint returns an endpoint function that calls the method
// "enqueued" of service "spin-broker".
func NewEnqueuedEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*EnqueuedPayload)
		return s.Enqueued(ctx, p)
	}
}

// NewStatusEndpoint returns an endpoint function that calls the method
// "status" of service "spin-broker".
func NewStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StatusPayload)
		return s.Status(ctx, p)
	}
}

// NewNextEndpoint returns an endpoint function that calls the method "next" of
// service "spin-broker".
func NewNextEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*NextPayload)
		return s.Next(ctx, p)
	}
}

// NewCompleteEndpoint returns an endpoint function that calls the method
// "complete" of service "spin-broker".
func NewCompleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CompletePayload)
		return nil, s.Complete(ctx, p)
	}
}
