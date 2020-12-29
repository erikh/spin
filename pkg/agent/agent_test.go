package agent

import (
	"context"
	"errors"
	"testing"
	"time"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	"code.hollensbe.org/erikh/spin/pkg/agent/dispatcher"
	"code.hollensbe.org/erikh/spin/pkg/services"
	goa "goa.design/goa/v3/pkg"
)

func TestBasicDispatch(t *testing.T) {
	host := services.SetupTestBroker(t)

	dispatcher := dispatcher.Table{
		"an_action": {
			Dispatch: func(c dispatcher.Command) error {
				return nil
			},
		},
		"error_action": {
			Dispatch: func(c dispatcher.Command) error {
				return errors.New("this is an error")
			},
		},
	}
	cc := brokerclient.Config{Host: host, Timeout: 1}
	a := New(cc, "resource", dispatcher)
	client := brokerclient.New(cc)

	pkg, err := client.New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Add(context.Background(), &spinbroker.AddPayload{
		ID:       pkg,
		Resource: "resource",
		Action:   "an_action",
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.Enqueue(context.Background(), pkg); err != nil {
		t.Fatal(err)
	}

	if err := a.Tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	status, err := client.Status(context.Background(), pkg)
	if err != nil {
		t.Fatal(err)
	}

	if !status.Status {
		t.Fatalf("status was unexpectedly invalid. Reason: %v", status.Reason)
	}

	pkg, err = client.New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Add(context.Background(), &spinbroker.AddPayload{
		ID:       pkg,
		Resource: "resource",
		Action:   "an_action",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Add(context.Background(), &spinbroker.AddPayload{
		ID:       pkg,
		Resource: "resource",
		Action:   "error_action",
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.Enqueue(context.Background(), pkg); err != nil {
		t.Fatal(err)
	}

	if err := a.Tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	status, err = client.Status(context.Background(), pkg)
	if err == nil {
		t.Fatal("No error with unfinished package")
	}

	if err := a.Tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	status, err = client.Status(context.Background(), pkg)
	if err != nil {
		t.Fatal(err)
	}

	if status.Status {
		t.Fatalf("status was unexpectedly valid.")
	}

	if *status.Reason != "this is an error" {
		t.Fatalf("Unexpected reason: %v", *status.Reason)
	}
}

func TestAgentLoop(t *testing.T) {
	count := 100 // x2 (positive+negative tests)

	host := services.SetupTestBroker(t)

	dispatcher := dispatcher.Table{
		"an_action": {
			Dispatch: func(c dispatcher.Command) error {
				return nil
			},
		},
		"error_action": {
			Dispatch: func(c dispatcher.Command) error {
				return errors.New("this is an error")
			},
		},
	}

	packages := map[string]bool{}

	cc := brokerclient.Config{Host: host, Timeout: 1}
	a := New(cc, "resource", dispatcher)
	client := brokerclient.New(cc)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	go a.Loop(ctx)

	for i := 0; i < count; i++ {
		for res := true; res; res = !res {
			pkg, err := client.New(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			action := "an_action"
			if !res {
				action = "error_action"
			}

			_, err = client.Add(context.Background(), &spinbroker.AddPayload{
				ID:       pkg,
				Resource: "resource",
				Action:   action,
			})
			if err != nil {
				t.Fatal(err)
			}

			if _, err := client.Enqueue(context.Background(), pkg); err != nil {
				t.Fatal(err)
			}

			packages[pkg] = res
		}
	}

	for pkg, res := range packages {
	top:
		stat, err := client.Status(context.Background(), pkg)
		if err != nil {
			switch err := err.(type) {
			case *goa.ServiceError:
				switch err.Name {
				case "record_not_found":
					time.Sleep(100 * time.Millisecond)
					goto top
				}
			}

			t.Fatal(err)
		}

		if stat.Status != res {
			t.Fatal("result wasn't expected")
		}
	}
}
