package broker

import (
	"errors"
	"testing"
)

func TestDispatcher(t *testing.T) {
	var action1 int
	var action2 int

	dispatcher := Dispatcher{
		"action1": Action{
			OptionalParameters: []string{"baz"},
			RequiredParameters: []string{"foo", "bar"},
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
	}

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
				Parameters: map[string]string{
					"foo": "quux",
					"bar": "quux2",
				},
			},
			pass: true,
		},
		"action1-green2": {
			command: Command{
				Action: "action1",
				Parameters: map[string]string{
					"foo": "quux",
					"bar": "quux2",
					"baz": "quux3",
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
				Parameters: map[string]string{
					"bar": "quux2",
					"baz": "quux3",
				},
			},
			pass:  false,
			error: ErrMissingRequiredParameter,
		},
		"action1-red3": {
			command: Command{
				Action: "action1",
				Parameters: map[string]string{
					"foo": "quux2",
					"baz": "quux3",
				},
			},
			pass:  false,
			error: ErrMissingRequiredParameter,
		},
		"action1-red4": {
			command: Command{
				Action: "action1",
				Parameters: map[string]string{
					"baz": "quux3",
				},
			},
			pass:  false,
			error: ErrMissingRequiredParameter,
		},
		"action1-red5": {
			command: Command{
				Action: "action1",
				Parameters: map[string]string{
					"foo":  "quux",
					"bar":  "quux2",
					"quux": "quux3",
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
				Parameters: map[string]string{
					"foo": "quux",
				},
			},
			pass:  false,
			error: ErrInvalidParameter,
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
