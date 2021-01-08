package agent

import (
	"context"
	"log"
	"time"

	brokerclient "github.com/erikh/spin/clients/broker"
	"github.com/erikh/spin/pkg/agent/dispatcher"
	goa "goa.design/goa/v3/pkg"
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

	// try to ensure the message gets delivered by not using the context passed;
	// which may have been cancelled while this was running.
	return a.client.Complete(context.Background(), nr.UUID, err == nil, sr)
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
			if e, ok := err.(*goa.ServiceError); ok && e.ErrorName() == "record_not_found" {
				time.Sleep(100 * time.Millisecond)
			} else if err != nil {
				log.Println("ERROR:", err)
				time.Sleep(time.Second)
			}
		}
	}
}
