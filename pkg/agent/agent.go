package agent

import (
	"context"
	"log"
	"time"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	"code.hollensbe.org/erikh/spin/pkg/broker"
)

type Agent struct {
	resource   string
	client     *brokerclient.Client
	dispatcher broker.Dispatcher
}

func New(cc brokerclient.Config, resource string, dispatcher broker.Dispatcher) *Agent {
	return &Agent{
		resource:   resource,
		dispatcher: dispatcher,
		client:     brokerclient.New(cc),
	}
}

func (a *Agent) Tick(ctx context.Context) error {
	nr, err := a.client.Next(ctx, &spinbroker.NextPayload{Resource: a.resource})
	if err != nil {
		return err
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

	err = a.client.Complete(ctx, &spinbroker.CompletePayload{
		ID:           nr.UUID,
		Status:       err == nil,
		StatusReason: sr,
	})

	return err
}

func (a *Agent) Loop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := a.Tick(ctx); err != nil {
			// FIXME will this actually work
			if err != broker.ErrRecordNotFound {
				log.Println(err)
				time.Sleep(time.Second)
			}
		}
	}
}
