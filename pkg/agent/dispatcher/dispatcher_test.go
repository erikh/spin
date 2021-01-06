package dispatcher

import (
	"encoding/json"
	"errors"
	"testing"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
)

func TestDispatcher(t *testing.T) {
	var action1 int
	var action2 int

	dispatcher := Table{
		"action1": Action{
			OptionalParameters: ParameterTable{"baz": TypeString},
			RequiredParameters: ParameterTable{"foo": TypeString, "bar": TypeString},
			Dispatch: func(c Command) error {
				action1++
				return nil
			},
		},
		"action2": Action{
			Dispatch: func(c Command) error {
				action2++
				return nil
			},
		},
		"error": Action{
			Dispatch: func(c Command) error {
				return errors.New("this is an error")
			},
		},
		"with_typed_parameters": Action{
			OptionalParameters: ParameterTable{"vm": func() interface{} { return &spinregistry.UpdatedVM{} }},
			RequiredParameters: ParameterTable{"string": TypeString, "uint": TypeUint64},
			Dispatch: func(c Command) error {
				switch c.Parameter("string").(type) {
				case *string:
				default:
					return errors.New("invalid type for string")
				}

				switch c.Parameter("uint").(type) {
				case *uint64:
				default:
					return errors.New("invalid type for string")
				}

				switch c.Parameter("vm").(type) {
				case *spinregistry.UpdatedVM:
				case nil:
				default:
					return errors.New("invalid type for string")
				}

				return nil
			},
		},
	}

	vm, _ := json.Marshal(&spinregistry.UpdatedVM{
		Name:   "foo",
		Cpus:   1,
		Memory: 1024,
		Images: []*spinregistry.Image{
			{
				Path: "test.raw",
			},
		},
	})

	table := map[string]struct {
		command Command
		pass    bool
		error   error
	}{
		"error": {
			command: Command{
				Action: "error",
			},
			pass:  false,
			error: errors.New("this is an error"),
		},
		"action1-green1": {
			command: Command{
				Action: "action1",
				Parameters: map[string]json.RawMessage{
					"foo": []byte(`"quux"`),
					"bar": []byte(`"quux2"`),
				},
			},
			pass: true,
		},
		"action1-green2": {
			command: Command{
				Action: "action1",
				Parameters: map[string]json.RawMessage{
					"foo": []byte(`"quux"`),
					"bar": []byte(`"quux2"`),
					"baz": []byte(`"quux3"`),
				},
			},
			pass: true,
		},
		"action1-red1": {
			command: Command{
				Action: "action1",
			},
			pass:  false,
			error: ErrMissingRequiredParameter,
		},
		"action1-red2": {
			command: Command{
				Action: "action1",
				Parameters: map[string]json.RawMessage{
					"bar": []byte(`"quux2"`),
					"baz": []byte(`"quux3"`),
				},
			},
			pass:  false,
			error: ErrMissingRequiredParameter,
		},
		"action1-red3": {
			command: Command{
				Action: "action1",
				Parameters: map[string]json.RawMessage{
					"foo": []byte(`"quux2"`),
					"baz": []byte(`"quux3"`),
				},
			},
			pass:  false,
			error: ErrMissingRequiredParameter,
		},
		"action1-red4": {
			command: Command{
				Action: "action1",
				Parameters: map[string]json.RawMessage{
					"baz": []byte(`"quux3"`),
				},
			},
			pass:  false,
			error: ErrMissingRequiredParameter,
		},
		"action1-red5": {
			command: Command{
				Action: "action1",
				Parameters: map[string]json.RawMessage{
					"foo":  []byte(`"quux"`),
					"bar":  []byte(`"quux2"`),
					"quux": []byte(`"quux3"`),
				},
			},
			pass:  false,
			error: ErrInvalidParameter,
		},
		"action2-green1": {
			command: Command{
				Action: "action2",
			},
			pass: true,
		},
		"action2-red1": {
			command: Command{
				Action: "action2",
				Parameters: map[string]json.RawMessage{
					"foo": []byte(`"quux"`),
				},
			},
			pass:  false,
			error: ErrInvalidParameter,
		},
		"with_parameters-green1": {
			command: Command{
				Action: "with_typed_parameters",
				Parameters: map[string]json.RawMessage{
					"string": []byte(`"string"`),
					"uint":   []byte(`1`),
				},
			},
			pass: true,
		},
		"with_parameters-green2": {
			command: Command{
				Action: "with_typed_parameters",
				Parameters: map[string]json.RawMessage{
					"string": []byte(`"string"`),
					"uint":   []byte(`1`),
					"vm":     vm,
				},
			},
			pass: true,
		},
		"with_parameters-red1": {
			command: Command{
				Action: "with_typed_parameters",
				Parameters: map[string]json.RawMessage{
					"string": []byte(`1`),
					"uint":   []byte(`1`),
					"vm":     vm,
				},
			},
			pass: false,
		},
		"with_parameters-red2": {
			command: Command{
				Action: "with_typed_parameters",
				Parameters: map[string]json.RawMessage{
					"string": []byte(`"string"`),
					"uint":   []byte(`"1"`),
					"vm":     vm,
				},
			},
			pass: false,
		},
		"with_parameters-red3": {
			command: Command{
				Action: "with_typed_parameters",
				Parameters: map[string]json.RawMessage{
					"string": []byte(`"string"`),
					"uint":   []byte(`1`),
					"vm":     []byte(`1`),
				},
			},
			pass: false,
		},
	}

	for name, test := range table {
		err := dispatcher.Dispatch(test.command)

		if err != nil {
			if test.pass {
				t.Fatalf("Test %q did not pass: %v", name, err)
			} else if test.error != nil && test.error.Error() != err.Error() {
				t.Fatalf("Test %q did not fail with the right error: %v (expected %v)", name, err, test.error)
			}
		} else if !test.pass {
			t.Fatalf("Test %q should not have passed and did", name)
		}
	}
}
