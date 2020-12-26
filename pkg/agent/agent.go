package agent

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	goahttp "goa.design/goa/v3/http"

	"code.hollensbe.org/erikh/spin/gen/http/spin_broker/client"
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	"code.hollensbe.org/erikh/spin/pkg/broker"
)

type ClientConfig struct {
	Proto   string
	Host    string
	Timeout int
}

type Agent struct {
	resource   string
	client     *client.Client
	dispatcher broker.Dispatcher
}

func (cc ClientConfig) MakeClient() *client.Client {
	return client.NewClient(
		cc.Proto,
		cc.Host,
		&http.Client{Timeout: time.Duration(cc.Timeout) * time.Second},
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		false,
	)
}

func New(cc ClientConfig, resource string, dispatcher broker.Dispatcher) *Agent {
	return &Agent{
		resource:   resource,
		dispatcher: dispatcher,
		client:     cc.MakeClient(),
	}
}

func (a *Agent) Tick() error {
	p, err := a.client.Next()(context.Background(), &spinbroker.NextPayload{Resource: a.resource})
	if err != nil {
		return err
	}

	nr, ok := p.(*spinbroker.NextResult)
	if !ok {
		return errors.New("invalid result")
	}

	err = a.dispatcher.Dispatch(broker.Command{
		UUID:       nr.UUID,
		Resource:   nr.Resource,
		Action:     nr.Action,
		Parameters: nr.Parameters,
	})

	var sr *string
	if err != nil {
		s := err.Error()
		sr = &s
	}

	_, err = a.client.Complete()(context.Background(), &spinbroker.CompletePayload{
		ID:           nr.UUID,
		Status:       err == nil,
		StatusReason: sr,
	})

	return err
}

func (a *Agent) Loop() error {
	for {
		if err := a.Tick(); err != nil {
			// FIXME will this actually work
			if err != broker.ErrRecordNotFound {
				log.Println(err)
				time.Sleep(time.Second)
			}
		}
	}
}
