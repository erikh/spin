package emulation

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	"code.hollensbe.org/erikh/spin/pkg/services"
	"code.hollensbe.org/erikh/spin/pkg/supervisor"
	goa "goa.design/goa/v3/pkg"
)

var testTable = map[string]test{
	"garbage": {
		commands: []command{{Action: "garbage"}},
		pass:     false,
	},
	"write_config_no_vm": {
		commands: []command{{
			Action: "write_config",
			Parameters: map[string]interface{}{
				"id": 1,
			}}},
		pass: false,
	},
	"write_config_no_id": {
		commands: []command{{
			Action: "write_config",
			Parameters: map[string]interface{}{
				"vm": &spinregistry.VM{
					Name:   "foo",
					Cpus:   1,
					Memory: 1024,
					Storage: []*spinregistry.Storage{
						{
							Image:     "test.raw",
							Volume:    "test",
							ImageSize: 50,
						},
					},
				},
			}}},
		pass: false,
	},
	"write_config_basic": {
		commands: []command{{
			Action: "write_config",
			Parameters: map[string]interface{}{
				"id": 1,
				"vm": &spinregistry.VM{
					Name:   "foo",
					Cpus:   1,
					Memory: 1024,
					Storage: []*spinregistry.Storage{
						{
							Image:     "test.raw",
							Volume:    "test",
							ImageSize: 50,
						},
					},
				},
			}}},
		pass: true,
		validate: func(ta testArgs) error {
			if ta.agentConfig.Supervisor.(*supervisor.Mock).Reloads != 1 {
				return errors.New("agent reloads were not equal to 1")
			}

			p := filepath.Join(ta.agentConfig.SystemDir, "spin-1.service")

			if _, err := os.Stat(p); err != nil {
				return err
			}

			f, err := os.Open(p)
			if err != nil {
				return err
			}
			defer f.Close()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}

			if !strings.Contains(string(content), "-m 1024M") {
				return errors.New("memory arg was invalid/not present")
			}

			return nil
		},
	},
	"remove_config_no_id": {
		commands: []command{{
			Action: "remove_config",
		}},
		pass: false,
	},
	"remove_config_non_existent_id": {
		commands: []command{{
			Action: "remove_config",
			Parameters: map[string]interface{}{
				"id": 1,
			},
		}},
		pass: false,
	},
	"remove_config_basic": {
		commands: []command{
			{
				Action: "write_config",
				Parameters: map[string]interface{}{
					"id": 1,
					"vm": &spinregistry.VM{
						Name:   "foo",
						Cpus:   1,
						Memory: 1024,
						Storage: []*spinregistry.Storage{
							{
								Image:     "test.raw",
								Volume:    "test",
								ImageSize: 50,
							},
						},
					},
				},
			},
			{
				Action: "remove_config",
				Parameters: map[string]interface{}{
					"id": 1,
				},
			},
		},
		pass: true,
		validate: func(ta testArgs) error {
			p := filepath.Join(ta.agentConfig.SystemDir, "spin-1.service")

			_, err := os.Stat(p)
			if os.IsNotExist(err) {
				return nil
			}

			return err
		},
	},
}

type testArgs struct {
	client      *brokerclient.Client
	agentConfig AgentConfig
}

type test struct {
	commands []command
	validate func(testArgs) error
	pass     bool
}

type command struct {
	Action     string
	Parameters map[string]interface{}
}

func (c command) params() map[string]json.RawMessage {
	params := map[string]json.RawMessage{}
	for key, value := range c.Parameters {
		content, _ := json.Marshal(value)
		params[key] = content
	}

	return params
}

func sendMessages(ctx context.Context, t *testing.T, client *brokerclient.Client, commands []command) string {
	pkg, err := client.New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, command := range commands {
		_, err := client.Add(ctx, &spinbroker.AddPayload{
			ID:         pkg,
			Resource:   ResourceType,
			Action:     command.Action,
			Parameters: command.params(),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	if _, err := client.Enqueue(ctx, pkg); err != nil {
		t.Fatal(err)
	}

	return pkg
}

func TestAgent(t *testing.T) {
	tempdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.RemoveAll(tempdir)
	})

	sv := supervisor.NewMock()

	brokerHost := services.SetupTestBroker(t)
	clientConfig := brokerclient.Config{
		Host:    brokerHost,
		Timeout: 1,
	}

	systemDir := filepath.Join(tempdir, "system")
	monitorDir := filepath.Join(tempdir, "monitors")

	ac := AgentConfig{
		SystemDir:    systemDir,
		MonitorDir:   monitorDir,
		ClientConfig: clientConfig,
		Supervisor:   sv,
	}

	agent, err := NewAgent(ac)
	if err != nil {
		t.Fatal(err)
	}

	client := brokerclient.New(clientConfig)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	errChan := make(chan error, 1)

	go func() {
		errChan <- agent.Loop(ctx)
	}()

	select {
	case err := <-errChan:
		t.Fatal(err)
	case <-time.After(100 * time.Millisecond):
	}

	for name, test := range testTable {
		sv.Reset()
		os.RemoveAll(systemDir)
		os.RemoveAll(monitorDir)

		pkg := sendMessages(ctx, t, client, test.commands)

		for {
			result, err := client.Status(ctx, pkg)
			if err != nil {
				if e, ok := err.(*goa.ServiceError); ok && e.ErrorName() == "record_not_found" {
					time.Sleep(100 * time.Millisecond)
					continue
				}

				t.Fatalf("Test %q failed with: %v", name, err)
			}

			if result == nil {
				t.Fatalf("result was nil for test %q", name)
			}

			if !result.Status && test.pass {
				t.Fatalf("Status was not success for test %q", name)
			} else if result.Status && !test.pass {
				t.Fatalf("Status was true and test %q was not supposed to pass", name)
			}

			break
		}

		if test.validate != nil {
			ta := testArgs{
				client:      client,
				agentConfig: ac,
			}

			if err := test.validate(ta); err != nil && test.pass {
				t.Fatalf("Error occurred during validation for %q: %v", name, err)
			} else if err == nil && !test.pass {
				t.Fatalf("No error occurred during validation for %q", name)
			}
		}
	}

}
