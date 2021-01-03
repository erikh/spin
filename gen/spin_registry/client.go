// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry client
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package spinregistry

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "spin-registry" service client.
type Client struct {
	VMCreateEndpoint goa.Endpoint
	VMUpdateEndpoint goa.Endpoint
	VMDeleteEndpoint goa.Endpoint
	VMGetEndpoint    goa.Endpoint
	VMListEndpoint   goa.Endpoint
}

// NewClient initializes a "spin-registry" service client given the endpoints.
func NewClient(vMCreate, vMUpdate, vMDelete, vMGet, vMList goa.Endpoint) *Client {
	return &Client{
		VMCreateEndpoint: vMCreate,
		VMUpdateEndpoint: vMUpdate,
		VMDeleteEndpoint: vMDelete,
		VMGetEndpoint:    vMGet,
		VMListEndpoint:   vMList,
	}
}

// VMCreate calls the "vm/create" endpoint of the "spin-registry" service.
func (c *Client) VMCreate(ctx context.Context, p *VM) (res uint64, err error) {
	var ires interface{}
	ires, err = c.VMCreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(uint64), nil
}

// VMUpdate calls the "vm/update" endpoint of the "spin-registry" service.
func (c *Client) VMUpdate(ctx context.Context, p *UpdateVM) (err error) {
	_, err = c.VMUpdateEndpoint(ctx, p)
	return
}

// VMDelete calls the "vm/delete" endpoint of the "spin-registry" service.
func (c *Client) VMDelete(ctx context.Context, p *VMDeletePayload) (err error) {
	_, err = c.VMDeleteEndpoint(ctx, p)
	return
}

// VMGet calls the "vm/get" endpoint of the "spin-registry" service.
func (c *Client) VMGet(ctx context.Context, p *VMGetPayload) (res *VM, err error) {
	var ires interface{}
	ires, err = c.VMGetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*VM), nil
}

// VMList calls the "vm/list" endpoint of the "spin-registry" service.
func (c *Client) VMList(ctx context.Context) (res []uint64, err error) {
	var ires interface{}
	ires, err = c.VMListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]uint64), nil
}
