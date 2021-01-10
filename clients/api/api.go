package api

import (
	"context"
	"net/http"
	"time"

	"github.com/erikh/spin/gen/http/spin_apiserver/client"
	spinapiserver "github.com/erikh/spin/gen/spin_apiserver"
	"github.com/erikh/spin/pkg/vm"
	goahttp "goa.design/goa/v3/http"
)

const apiProto = "http"

// Config is the configuration struct for the HTTP client. Timeout is in
// seconds.  Host should contain a port if not port 80.
type Config struct {
	Host    string
	Timeout int
}

// Client is the outer handle for the broker HTTP client.
type Client struct {
	client *client.Client
}

// New constructs a new http client from the configuration provided.
func New(cc Config) *Client {
	return &Client{client.NewClient(
		apiProto,
		cc.Host,
		&http.Client{Timeout: time.Duration(cc.Timeout) * time.Second},
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		false,
	)}
}

// VMCreate creates a new VM based on the properties provided.
func (c *Client) VMCreate(ctx context.Context, vm *vm.Create) (uint64, error) {
	if err := vm.Validate(); err != nil {
		return 0, err
	}

	res, err := c.client.VMCreate()(ctx, vm)
	if err != nil {
		return 0, err
	}

	return res.(uint64), nil
}

// VMDelete deletes a vm by ID.
func (c *Client) VMDelete(ctx context.Context, id uint64) error {
	_, err := c.client.VMDelete()(ctx, &spinapiserver.VMDeletePayload{ID: id})
	return err
}

// VMList lists the ids of all vms
func (c *Client) VMList(ctx context.Context) ([]uint64, error) {
	res, err := c.client.VMList()(ctx, nil)
	if err != nil {
		return nil, err
	}

	return res.([]uint64), nil
}

// VMUpdate updates a single VM for an ID.
func (c *Client) VMUpdate(ctx context.Context, id uint64, vm *vm.Transient) error {
	if err := vm.Validate(); err != nil {
		return err
	}

	_, err := c.client.VMUpdate()(ctx, &spinapiserver.VMUpdatePayload{ID: id, VM: vm})
	return err
}

// VMGet gets a VM by ID
func (c *Client) VMGet(ctx context.Context, id uint64) (*vm.Transient, error) {
	res, err := c.client.VMGet()(ctx, &spinapiserver.VMGetPayload{ID: id})
	if err != nil {
		return nil, err
	}

	return res.(*vm.Transient), nil
}

// ControlStart starts a VM by id.
func (c *Client) ControlStart(ctx context.Context, id uint64) error {
	_, err := c.client.ControlStart()(ctx, &spinapiserver.ControlStartPayload{ID: id})
	return err
}

// ControlStop attempts to gracefully stop a VM by id. If it fails, it will
// terminate the vm forcefully after 30 seconds.
func (c *Client) ControlStop(ctx context.Context, id uint64) error {
	_, err := c.client.ControlStop()(ctx, &spinapiserver.ControlStopPayload{ID: id})
	return err
}

// ControlShutdown attempts to gracefully shutdown the VM by ID. It will take
// no further actions if it does not succeed.
func (c *Client) ControlShutdown(ctx context.Context, id uint64) error {
	_, err := c.client.ControlShutdown()(ctx, &spinapiserver.ControlShutdownPayload{ID: id})
	return err
}
