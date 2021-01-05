// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver client HTTP transport
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the spin-apiserver service endpoint HTTP clients.
type Client struct {
	// VMCreate Doer is the HTTP client used to make requests to the vm/create
	// endpoint.
	VMCreateDoer goahttp.Doer

	// VMDelete Doer is the HTTP client used to make requests to the vm/delete
	// endpoint.
	VMDeleteDoer goahttp.Doer

	// ControlStart Doer is the HTTP client used to make requests to the
	// control/start endpoint.
	ControlStartDoer goahttp.Doer

	// ControlStop Doer is the HTTP client used to make requests to the
	// control/stop endpoint.
	ControlStopDoer goahttp.Doer

	// ControlShutdown Doer is the HTTP client used to make requests to the
	// control/shutdown endpoint.
	ControlShutdownDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the spin-apiserver service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		VMCreateDoer:        doer,
		VMDeleteDoer:        doer,
		ControlStartDoer:    doer,
		ControlStopDoer:     doer,
		ControlShutdownDoer: doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// VMCreate returns an endpoint that makes HTTP requests to the spin-apiserver
// service vm/create server.
func (c *Client) VMCreate() goa.Endpoint {
	var (
		encodeRequest  = EncodeVMCreateRequest(c.encoder)
		decodeResponse = DecodeVMCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildVMCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VMCreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-apiserver", "vm/create", err)
		}
		return decodeResponse(resp)
	}
}

// VMDelete returns an endpoint that makes HTTP requests to the spin-apiserver
// service vm/delete server.
func (c *Client) VMDelete() goa.Endpoint {
	var (
		decodeResponse = DecodeVMDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildVMDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VMDeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-apiserver", "vm/delete", err)
		}
		return decodeResponse(resp)
	}
}

// ControlStart returns an endpoint that makes HTTP requests to the
// spin-apiserver service control/start server.
func (c *Client) ControlStart() goa.Endpoint {
	var (
		decodeResponse = DecodeControlStartResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildControlStartRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ControlStartDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-apiserver", "control/start", err)
		}
		return decodeResponse(resp)
	}
}

// ControlStop returns an endpoint that makes HTTP requests to the
// spin-apiserver service control/stop server.
func (c *Client) ControlStop() goa.Endpoint {
	var (
		decodeResponse = DecodeControlStopResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildControlStopRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ControlStopDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-apiserver", "control/stop", err)
		}
		return decodeResponse(resp)
	}
}

// ControlShutdown returns an endpoint that makes HTTP requests to the
// spin-apiserver service control/shutdown server.
func (c *Client) ControlShutdown() goa.Endpoint {
	var (
		decodeResponse = DecodeControlShutdownResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildControlShutdownRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ControlShutdownDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-apiserver", "control/shutdown", err)
		}
		return decodeResponse(resp)
	}
}
