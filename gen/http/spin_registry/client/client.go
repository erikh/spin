// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry client HTTP transport
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

// Client lists the spin-registry service endpoint HTTP clients.
type Client struct {
	// VMCreate Doer is the HTTP client used to make requests to the vm/create
	// endpoint.
	VMCreateDoer goahttp.Doer

	// VMUpdate Doer is the HTTP client used to make requests to the vm/update
	// endpoint.
	VMUpdateDoer goahttp.Doer

	// VMDelete Doer is the HTTP client used to make requests to the vm/delete
	// endpoint.
	VMDeleteDoer goahttp.Doer

	// VMGet Doer is the HTTP client used to make requests to the vm/get endpoint.
	VMGetDoer goahttp.Doer

	// VMList Doer is the HTTP client used to make requests to the vm/list endpoint.
	VMListDoer goahttp.Doer

	// StorageVolumesList Doer is the HTTP client used to make requests to the
	// storage/volumes/list endpoint.
	StorageVolumesListDoer goahttp.Doer

	// StorageVolumesCreate Doer is the HTTP client used to make requests to the
	// storage/volumes/create endpoint.
	StorageVolumesCreateDoer goahttp.Doer

	// StorageVolumesDelete Doer is the HTTP client used to make requests to the
	// storage/volumes/delete endpoint.
	StorageVolumesDeleteDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the spin-registry service
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
		VMCreateDoer:             doer,
		VMUpdateDoer:             doer,
		VMDeleteDoer:             doer,
		VMGetDoer:                doer,
		VMListDoer:               doer,
		StorageVolumesListDoer:   doer,
		StorageVolumesCreateDoer: doer,
		StorageVolumesDeleteDoer: doer,
		RestoreResponseBody:      restoreBody,
		scheme:                   scheme,
		host:                     host,
		decoder:                  dec,
		encoder:                  enc,
	}
}

// VMCreate returns an endpoint that makes HTTP requests to the spin-registry
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
			return nil, goahttp.ErrRequestError("spin-registry", "vm/create", err)
		}
		return decodeResponse(resp)
	}
}

// VMUpdate returns an endpoint that makes HTTP requests to the spin-registry
// service vm/update server.
func (c *Client) VMUpdate() goa.Endpoint {
	var (
		encodeRequest  = EncodeVMUpdateRequest(c.encoder)
		decodeResponse = DecodeVMUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildVMUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VMUpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-registry", "vm/update", err)
		}
		return decodeResponse(resp)
	}
}

// VMDelete returns an endpoint that makes HTTP requests to the spin-registry
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
			return nil, goahttp.ErrRequestError("spin-registry", "vm/delete", err)
		}
		return decodeResponse(resp)
	}
}

// VMGet returns an endpoint that makes HTTP requests to the spin-registry
// service vm/get server.
func (c *Client) VMGet() goa.Endpoint {
	var (
		decodeResponse = DecodeVMGetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildVMGetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VMGetDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-registry", "vm/get", err)
		}
		return decodeResponse(resp)
	}
}

// VMList returns an endpoint that makes HTTP requests to the spin-registry
// service vm/list server.
func (c *Client) VMList() goa.Endpoint {
	var (
		decodeResponse = DecodeVMListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildVMListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VMListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-registry", "vm/list", err)
		}
		return decodeResponse(resp)
	}
}

// StorageVolumesList returns an endpoint that makes HTTP requests to the
// spin-registry service storage/volumes/list server.
func (c *Client) StorageVolumesList() goa.Endpoint {
	var (
		decodeResponse = DecodeStorageVolumesListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStorageVolumesListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StorageVolumesListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-registry", "storage/volumes/list", err)
		}
		return decodeResponse(resp)
	}
}

// StorageVolumesCreate returns an endpoint that makes HTTP requests to the
// spin-registry service storage/volumes/create server.
func (c *Client) StorageVolumesCreate() goa.Endpoint {
	var (
		encodeRequest  = EncodeStorageVolumesCreateRequest(c.encoder)
		decodeResponse = DecodeStorageVolumesCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStorageVolumesCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StorageVolumesCreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-registry", "storage/volumes/create", err)
		}
		return decodeResponse(resp)
	}
}

// StorageVolumesDelete returns an endpoint that makes HTTP requests to the
// spin-registry service storage/volumes/delete server.
func (c *Client) StorageVolumesDelete() goa.Endpoint {
	var (
		encodeRequest  = EncodeStorageVolumesDeleteRequest(c.encoder)
		decodeResponse = DecodeStorageVolumesDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStorageVolumesDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StorageVolumesDeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("spin-registry", "storage/volumes/delete", err)
		}
		return decodeResponse(resp)
	}
}
