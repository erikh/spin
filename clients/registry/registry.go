package registryclient

import (
	"context"
	"net/http"
	"time"

	"code.hollensbe.org/erikh/spin/gen/http/spin_registry/client"
	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	goahttp "goa.design/goa/v3/http"
)

const registryProto = "http"

// Config is the configuration struct for the HTTP client. Timeout is in
// seconds.  Host should contain a port if not port 80.
type Config struct {
	Host    string
	Timeout int
}

// Client is the outer handle for the registry HTTP client.
type Client struct {
	client *client.Client
}

// New constructs a new http client from the configuration provided.
func New(cc Config) *Client {
	return &Client{client.NewClient(
		registryProto,
		cc.Host,
		&http.Client{Timeout: time.Duration(cc.Timeout) * time.Second},
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		false,
	)}
}

// Create creates a new vm.
func (c *Client) VMCreate(ctx context.Context, vm *spinregistry.VM) (uint64, error) {
	pkg, err := c.client.VMCreate()(ctx, vm)
	if err != nil {
		return 0, err
	}

	return pkg.(uint64), nil
}

// Update updates a vm by id.
func (c *Client) VMUpdate(ctx context.Context, id uint64, vm *spinregistry.VM) error {
	_, err := c.client.VMUpdate()(ctx, &spinregistry.UpdateVM{ID: id, VM: vm})
	return err
}

// Delete deletes a vm by id.
func (c *Client) VMDelete(ctx context.Context, id uint64) error {
	_, err := c.client.VMDelete()(ctx, &spinregistry.VMDeletePayload{ID: id})
	return err
}

// Get retrieves a vm by id.
func (c *Client) Get(ctx context.Context, id uint64) (*spinregistry.VM, error) {
	vm, err := c.client.VMGet()(ctx, &spinregistry.VMGetPayload{ID: id})
	if err != nil {
		return nil, err
	}

	return vm.(*spinregistry.VM), nil
}

// List retrieves all IDs of all VMs.
func (c *Client) VMList(ctx context.Context) ([]uint64, error) {
	res, err := c.client.VMList()(ctx, nil)
	if err != nil {
		return nil, err
	}

	return res.([]uint64), nil
}
