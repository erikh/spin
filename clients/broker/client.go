package brokerclient

import (
	"context"
	"net/http"
	"time"

	"code.hollensbe.org/erikh/spin/gen/http/spin_broker/client"
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	goahttp "goa.design/goa/v3/http"
)

type Config struct {
	Proto   string
	Host    string
	Timeout int
}

type Client struct {
	client *client.Client
}

func New(cc Config) *Client {
	return &Client{client.NewClient(
		cc.Proto,
		cc.Host,
		&http.Client{Timeout: time.Duration(cc.Timeout) * time.Second},
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		false,
	)}
}

func (c *Client) New(ctx context.Context) (string, error) {
	pkg, err := c.client.New()(ctx, nil)
	if err != nil {
		return "", err
	}

	return pkg.(string), nil
}

func (c *Client) Add(ctx context.Context, p *spinbroker.AddPayload) (string, error) {
	uuid, err := c.client.Add()(ctx, p)
	if err != nil {
		return "", err
	}

	return uuid.(string), nil
}

func (c *Client) Enqueue(ctx context.Context, uuid string) ([]string, error) {
	res, err := c.client.Enqueue()(ctx, &spinbroker.EnqueuePayload{ID: uuid})
	if err != nil {
		return nil, err
	}

	return res.([]string), nil
}

func (c *Client) Status(ctx context.Context, uuid string) (*spinbroker.StatusResult, error) {
	res, err := c.client.Status()(ctx, &spinbroker.StatusPayload{ID: uuid})
	if err != nil {
		return nil, err
	}

	return res.(*spinbroker.StatusResult), nil
}

func (c *Client) Complete(ctx context.Context, uuid string, status bool, reason *string) error {
	_, err := c.client.Complete()(ctx, &spinbroker.CompletePayload{
		ID:           uuid,
		Status:       status,
		StatusReason: reason,
	})
	return err
}

func (c *Client) Next(ctx context.Context, resource string) (*spinbroker.NextResult, error) {
	res, err := c.client.Next()(ctx, &spinbroker.NextPayload{Resource: resource})
	if err != nil {
		return nil, err
	}

	return res.(*spinbroker.NextResult), nil
}
