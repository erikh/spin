package storage

import (
	"context"
	"errors"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"testing"
	"time"

	brokerclient "github.com/erikh/spin/clients/broker"
	spinbroker "github.com/erikh/spin/gen/spin_broker"
	"github.com/erikh/spin/pkg/services"
	goa "goa.design/goa/v3/pkg"
)

type testArgs struct {
	client *brokerclient.Client
	dir    string
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

var testTable = map[string]test{
	"garbage": {
		commands: []command{{Action: "garbage"}},
		pass:     false,
	},
	"add_volume green": {
		commands: []command{{Action: "add_volume", Parameters: map[string]interface{}{
			"path": "test",
		}}},
		pass: true,
		validate: func(ta testArgs) error {
			_, err := os.Stat(filepath.Join(ta.dir, "test"))
			return err
		},
	},
	"add_volume red 1": {
		commands: []command{{Action: "add_volume", Parameters: map[string]interface{}{
			"path": "test/one",
		}}},
		pass: false,
	},
	"add_volume red 2": {
		commands: []command{{Action: "add_volume", Parameters: map[string]interface{}{
			"path": nil,
		}}},
		pass: false,
	},
	"add_volume red 3": {
		commands: []command{{Action: "add_volume", Parameters: map[string]interface{}{
			"path": "test/../",
		}}},
		pass: false,
	},
	"add_volume red 4": {
		commands: []command{{Action: "add_volume", Parameters: map[string]interface{}{
			"path": ".",
		}}},
		pass: false,
	},
	"add_volume red 5": {
		commands: []command{{Action: "add_volume", Parameters: map[string]interface{}{
			"path": "",
		}}},
		pass: false,
	},
	"remove_volume green": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "remove_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
		},
		pass: true,
		validate: func(ta testArgs) error {
			_, err := os.Stat(filepath.Join(ta.dir, "test"))
			if os.IsNotExist(err) {
				return nil
			}

			if err == nil {
				return errors.New("still exists")
			}

			return err
		},
	},
	"remove_volume red 1": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "remove_volume",
				Parameters: map[string]interface{}{
					"path": ".",
				},
			},
		},
		pass: false,
	},
	"remove_volume red 2": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "remove_volume",
				Parameters: map[string]interface{}{
					"path": "test/one",
				},
			},
		},
		pass: false,
	},
	"remove_volume red 3": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "remove_volume",
				Parameters: map[string]interface{}{
					"path": "..",
				},
			},
		},
		pass: false,
	},
	"remove_volume red 4": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "remove_volume",
				Parameters: map[string]interface{}{
					"path": nil,
				},
			},
		},
		pass: false,
	},
	"remove_volume red 5": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "remove_volume",
				Parameters: map[string]interface{}{
					"path": "",
				},
			},
		},
		pass: false,
	},
	"create_image green": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
		},
		pass: true,
		validate: func(ta testArgs) error {
			fi, err := os.Stat(filepath.Join(ta.dir, "test", "test.raw"))
			if err != nil {
				return err
			}

			if fi.Size() != int64(5*math.Pow(1024, 3)) {
				return errors.New("image_size is not valid")
			}
			return nil
		},
	},
	"create_image red 1": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 0,
				},
			},
		},
		pass: false,
	},
	"create_image red 2": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": nil,
				},
			},
		},
		pass: false,
	},
	"create_image red 3": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "",
					"image_size": 5,
				},
			},
		},
		pass: false,
	},
	"create_image red 4": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      nil,
					"image_size": 5,
				},
			},
		},
		pass: false,
	},
	"create_image red 5": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
		},
		pass: false,
	},
	"create_image red 6": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     nil,
					"image":      "test.raw",
					"image_size": 5,
				},
			},
		},
		pass: false,
	},
	"create_image red 7": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     ".",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
		},
		pass: false,
	},
	"create_image red 8": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "/",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
		},
		pass: false,
	},
	"create_image red 9": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      ".",
					"image_size": 5,
				},
			},
		},
		pass: false,
	},
	"create_image red 10": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "/",
					"image_size": 5,
				},
			},
		},
		pass: false,
	},
	"delete_image green": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": "test",
					"image":  "test.raw",
				},
			},
		},
		pass: true,
		validate: func(ta testArgs) error {
			_, err := os.Stat(filepath.Join(ta.dir, "test", "test.raw"))
			if os.IsNotExist(err) {
				return nil
			}

			if err == nil {
				return errors.New("image still exists")
			}

			return err
		},
	},
	"delete_image red 1": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": "",
					"image":  "test.raw",
				},
			},
		},
		pass: false,
	},
	"delete_image red 2": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": nil,
					"image":  "test.raw",
				},
			},
		},
		pass: false,
	},
	"delete_image red 3": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": ".",
					"image":  "test.raw",
				},
			},
		},
		pass: false,
	},
	"delete_image red 4": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": "test/one",
					"image":  "test.raw",
				},
			},
		},
		pass: false,
	},
	"delete_image red 5": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": "test",
					"image":  nil,
				},
			},
		},
		pass: false,
	},
	"delete_image red 6": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": "test",
					"image":  "",
				},
			},
		},
		pass: false,
	},
	"delete_image red 7": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": "test",
					"image":  "..",
				},
			},
		},
		pass: false,
	},
	"delete_image red 8": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": "test",
					"image":  "test/",
				},
			},
		},
		pass: false,
	},
	"delete_image red 9": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"volume": "test",
				},
			},
		},
		pass: false,
	},
	"delete_image red 10": {
		commands: []command{
			{
				Action: "add_volume",
				Parameters: map[string]interface{}{
					"path": "test",
				},
			},
			{
				Action: "create_image",
				Parameters: map[string]interface{}{
					"volume":     "test",
					"image":      "test.raw",
					"image_size": 5,
				},
			},
			{
				Action: "delete_image",
				Parameters: map[string]interface{}{
					"image": "test/",
				},
			},
		},
		pass: false,
	},
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
			Parameters: command.Parameters,
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

func TestHostPathAgent(t *testing.T) {
	tempdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		os.RemoveAll(tempdir)
	})

	brokerHost := services.SetupTestBroker(t)
	clientConfig := brokerclient.Config{
		Host:    brokerHost,
		Timeout: 1,
	}

	agent := NewHostPathAgent(tempdir, clientConfig)
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
		if err := os.RemoveAll(tempdir); err != nil {
			t.Fatal(err)
		}

		if err := os.MkdirAll(tempdir, 0700); err != nil {
			t.Fatal(err)
		}

		pkg := sendMessages(ctx, t, client, test.commands)

		for {
			result, err := client.Status(ctx, pkg)
			if err != nil {
				if e, ok := err.(*goa.ServiceError); ok && e.ErrorName() == "record_not_found" {
					time.Sleep(100 * time.Millisecond)
					continue
				} else if ok {
					t.Fatalf("Test %q failed with: %v", name, e.ErrorName())
				} else {
					t.Fatalf("Test %q failed with: %v (%T)", name, err, err.(*goa.ServiceError).ErrorName())
				}
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
				client: client,
				dir:    tempdir,
			}

			if err := test.validate(ta); err != nil && test.pass {
				t.Fatalf("Error occurred during validation for %q: %v", name, err)
			} else if err == nil && !test.pass {
				t.Fatalf("No error occurred during validation for %q", name)
			}
		}
	}
}
