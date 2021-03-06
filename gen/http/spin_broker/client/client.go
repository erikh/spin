// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker client HTTP transport
//
// Command:
// $ goa gen github.com/erikh/spin/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the spin-broker service endpoint HTTP clients.
type Client struct {
	// New Doer is the HTTP client used to make requests to the new endpoint.
	NewDoer goahttp.Doer

	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Enqueue Doer is the HTTP client used to make requests to the enqueue
	// endpoint.
	EnqueueDoer goahttp.Doer

	// Status Doer is the HTTP client used to make requests to the status endpoint.
	StatusDoer goahttp.Doer

	// Next Doer is the HTTP client used to make requests to the next endpoint.
	NextDoer goahttp.Doer

	// Complete Doer is the HTTP client used to make requests to the complete
	// endpoint.
	CompleteDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the spin-broker service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		NewDoer:             doer,
		AddDoer:             doer,
		EnqueueDoer:         doer,
		StatusDoer:          doer,
		NextDoer:            doer,
		CompleteDoer:        doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// New returns an endpoint that makes HTTP requests to the spin-broker service
// new server.
func (c *Client) New() goa.Endpoint {
	var (
		decodeResponse = DecodeNewResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildNewRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.NewDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-broker", "new", err)
		}
		return decodeResponse(resp)
	}
}

// Add returns an endpoint that makes HTTP requests to the spin-broker service
// add server.
func (c *Client) Add() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddRequest(c.encoder)
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-broker", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Enqueue returns an endpoint that makes HTTP requests to the spin-broker
// service enqueue server.
func (c *Client) Enqueue() goa.Endpoint {
	var (
		decodeResponse = DecodeEnqueueResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildEnqueueRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.EnqueueDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-broker", "enqueue", err)
		}
		return decodeResponse(resp)
	}
}

// Status returns an endpoint that makes HTTP requests to the spin-broker
// service status server.
func (c *Client) Status() goa.Endpoint {
	var (
		decodeResponse = DecodeStatusResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStatusRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StatusDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-broker", "status", err)
		}
		return decodeResponse(resp)
	}
}

// Next returns an endpoint that makes HTTP requests to the spin-broker service
// next server.
func (c *Client) Next() goa.Endpoint {
	var (
		decodeResponse = DecodeNextResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildNextRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.NextDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-broker", "next", err)
		}
		return decodeResponse(resp)
	}
}

// Complete returns an endpoint that makes HTTP requests to the spin-broker
// service complete server.
func (c *Client) Complete() goa.Endpoint {
	var (
		encodeRequest  = EncodeCompleteRequest(c.encoder)
		decodeResponse = DecodeCompleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCompleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CompleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-broker", "complete", err)
		}
		return decodeResponse(resp)
	}
}
