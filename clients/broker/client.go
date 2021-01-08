package brokerclient

import (
	"context"
	"net/http"
	"time"

	"github.com/erikh/spin/gen/http/spin_broker/client"
	spinbroker "github.com/erikh/spin/gen/spin_broker"
	goahttp "goa.design/goa/v3/http"
)

const brokerProto = "http"

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
		brokerProto,
		cc.Host,
		&http.Client{Timeout: time.Duration(cc.Timeout) * time.Second},
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		false,
	)}
}

// New creates a new package.
func (c *Client) New(ctx context.Context) (string, error) {
	pkg, err := c.client.New()(ctx, nil)
	if err != nil {
		return "", err
	}

	return pkg.(string), nil
}

// Add adds a command to the package. The ID here is the package UUID, not the
// one of the command. The UUID of the newly added command will be returned as
// a part of this call.
func (c *Client) Add(ctx context.Context, p *spinbroker.AddPayload) (string, error) {
	uuid, err := c.client.Add()(ctx, p)
	if err != nil {
		return "", err
	}

	return uuid.(string), nil
}

// Enqueue enqueues the package in the resource queues.
func (c *Client) Enqueue(ctx context.Context, uuid string) ([]string, error) {
	res, err := c.client.Enqueue()(ctx, &spinbroker.EnqueuePayload{ID: uuid})
	if err != nil {
		return nil, err
	}

	return res.([]string), nil
}

// Status checks the status of a package by UUID. It returns a StatusResult on
// a successful fetch, which may contain a failed value. If the status is
// failed, the StatusReason may have further information.
func (c *Client) Status(ctx context.Context, uuid string) (*spinbroker.StatusResult, error) {
	res, err := c.client.Status()(ctx, &spinbroker.StatusPayload{ID: uuid})
	if err != nil {
		return nil, err
	}

	return res.(*spinbroker.StatusResult), nil
}

// Complete completes a command, with the UUID being that of the command, the
// status being a pass/fail and the reason provided if the status failed.
func (c *Client) Complete(ctx context.Context, uuid string, status bool, reason *string) error {
	_, err := c.client.Complete()(ctx, &spinbroker.CompletePayload{
		ID:           uuid,
		Status:       status,
		StatusReason: reason,
	})
	return err
}

// Next returns the next item in the queue by resource. If there are no items,
// broker.ErrRecordNotFound will be returned.
func (c *Client) Next(ctx context.Context, resource string) (*spinbroker.NextResult, error) {
	res, err := c.client.Next()(ctx, &spinbroker.NextPayload{Resource: resource})
	if err != nil {
		return nil, err
	}

	return res.(*spinbroker.NextResult), nil
}
