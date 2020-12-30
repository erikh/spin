package agent

import (
	"context"
	"time"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	"code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"
)

// Agent is a struct encapuslating a runner agent that executes tasks for a
// specific resource. Upon dispatch, it will automatically complete the
// resource in the queue. It is assumed that the consumer will be familiar with
// the semantics of the broker and dispatcher before implementing an agent.
type Agent struct {
	resource   string
	client     *brokerclient.Client
	dispatcher dispatcher.Table
}

// New constructs a new agent. Typically used inside other constructors, this
// sews together the dispatcher, resource and a client for future "ticking" or
// looping.
func New(cc brokerclient.Config, resource string, dispatcher dispatcher.Table) *Agent {
	return &Agent{
		resource:   resource,
		dispatcher: dispatcher,
		client:     brokerclient.New(cc),
	}
}

// Tick runs the loop iteration one time and returns any error. If there is
// nothing in the queue or calling complete yields an error, that will be
// returned. Otherwise, a status will be set based on the result of the
// dispatch as a part of the Complete call.
func (a *Agent) Tick(ctx context.Context) error {
	nr, err := a.client.Next(ctx, a.resource)
	if err != nil {
		return err
	}

	err = a.dispatcher.Dispatch(dispatcher.Command{
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

	return a.client.Complete(ctx, nr.UUID, err == nil, sr)
}

// Loop runs the full loop which will wait at appropriate times to avoid
// bombing out the service. Cancel the context to terminate it.
func (a *Agent) Loop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := a.Tick(ctx); err != nil {
			// FIXME handle these errors in some more intelligent fashion
			if err != nil {
				time.Sleep(time.Second)
			}
		}
	}
}
