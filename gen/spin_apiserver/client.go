// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver client
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package spinapiserver

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "spin-apiserver" service client.
type Client struct {
	VMCreateEndpoint        goa.Endpoint
	VMDeleteEndpoint        goa.Endpoint
	VMListEndpoint          goa.Endpoint
	ControlStartEndpoint    goa.Endpoint
	ControlStopEndpoint     goa.Endpoint
	ControlShutdownEndpoint goa.Endpoint
}

// NewClient initializes a "spin-apiserver" service client given the endpoints.
func NewClient(vMCreate, vMDelete, vMList, controlStart, controlStop, controlShutdown goa.Endpoint) *Client {
	return &Client{
		VMCreateEndpoint:        vMCreate,
		VMDeleteEndpoint:        vMDelete,
		VMListEndpoint:          vMList,
		ControlStartEndpoint:    controlStart,
		ControlStopEndpoint:     controlStop,
		ControlShutdownEndpoint: controlShutdown,
	}
}

// VMCreate calls the "vm_create" endpoint of the "spin-apiserver" service.
func (c *Client) VMCreate(ctx context.Context, p *CreateVM) (res uint64, err error) {
	var ires interface{}
	ires, err = c.VMCreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(uint64), nil
}

// VMDelete calls the "vm_delete" endpoint of the "spin-apiserver" service.
func (c *Client) VMDelete(ctx context.Context, p *VMDeletePayload) (err error) {
	_, err = c.VMDeleteEndpoint(ctx, p)
	return
}

// VMList calls the "vm_list" endpoint of the "spin-apiserver" service.
func (c *Client) VMList(ctx context.Context) (res []uint64, err error) {
	var ires interface{}
	ires, err = c.VMListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]uint64), nil
}

// ControlStart calls the "control_start" endpoint of the "spin-apiserver"
// service.
func (c *Client) ControlStart(ctx context.Context, p *ControlStartPayload) (err error) {
	_, err = c.ControlStartEndpoint(ctx, p)
	return
}

// ControlStop calls the "control_stop" endpoint of the "spin-apiserver"
// service.
func (c *Client) ControlStop(ctx context.Context, p *ControlStopPayload) (err error) {
	_, err = c.ControlStopEndpoint(ctx, p)
	return
}

// ControlShutdown calls the "control_shutdown" endpoint of the
// "spin-apiserver" service.
func (c *Client) ControlShutdown(ctx context.Context, p *ControlShutdownPayload) (err error) {
	_, err = c.ControlShutdownEndpoint(ctx, p)
	return
}
