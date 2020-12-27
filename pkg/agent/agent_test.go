package agent

import (
	"context"
	"fmt"
	"testing"

	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	"code.hollensbe.org/erikh/spin/pkg/broker"
	"code.hollensbe.org/erikh/spin/pkg/services"
)

func TestBasicDispatch(t *testing.T) {
	host := services.SetupTestBroker(t)

	dispatcher := broker.Dispatcher{
		"an_action": {
			Dispatch: func(c broker.Command) error {
				fmt.Println(c)
				return nil
			},
		},
	}
	cc := ClientConfig{Proto: "http", Host: host, Timeout: 1}
	a := New(cc, "resource", dispatcher)
	client := cc.MakeClient()

	pkg, err := client.New()(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Add()(context.Background(), &spinbroker.AddPayload{
		ID:       pkg.(string),
		Resource: "resource",
		Action:   "an_action",
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.Enqueue()(context.Background(), &spinbroker.EnqueuePayload{ID: pkg.(string)}); err != nil {
		t.Fatal(err)
	}

	if err := a.Tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	status, err := client.Status()(context.Background(), &spinbroker.StatusPayload{ID: pkg.(string)})
	if err != nil {
		t.Fatal(err)
	}

	result := status.(*spinbroker.StatusResult)
	if !result.Status {
		t.Fatalf("status was unexpectedly invalid. Reason: %v", result.Reason)
	}
}
