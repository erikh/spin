package dispatcher

import (
	"encoding/json"
	"errors"
	"fmt"
)

// TypeUint64 is a callback that returns a *uint64 for marshalling.
func TypeUint64() interface{} { var i uint64; return &i }

// TypeString is a callback that returns a *string for marshalling.
func TypeString() interface{} { var i string; return &i }

// Command is a unit of instruction; it contains a UUID, the unique identifier
// of the commmand, a Resource, the type of command to execute, an action, the
// name of the command to execute, and parameters, a collection of items that
// relate to the action for the purposes of execution.
//
// Commands are typically fed to Packages, then the Package is Enqueued, Next()
// calls are made to yield the commands for the resource, the command is
// processed, FinishCommand is called to finish the command, then statuses are
// polled and eventually yielded.
type Command struct {
	UUID         string
	Resource     string
	Action       string
	Parameters   map[string]json.RawMessage
	parameters   map[string]interface{}
	Dependencies []string
}

// Table encapsulates a dispatching system that consists of actions
// (strings) that correspond to processed/validated properties and a processing
// function.
type Table map[string]Action

// Func is the dispatch action function. It accepts a command and the status is
// published back to the broker automatically as a part of the Tick() and
// Loop() calls.
type Func func(Command) error

// ParameterTable is a map of parameter name -> type creation function.  The
// type creation function is expected to return a type that is compatible with
// the JSON marshalling of the API. You must return a pointer from this call or
// encoding/json will vomit.
type ParameterTable map[string]func() interface{}

// Action is the definition of the protocol action item.
type Action struct {
	RequiredParameters ParameterTable
	OptionalParameters ParameterTable
	Dispatch           Func
}

var (
	// ErrActionNotFound is returned when an action cannot be dispatched.
	ErrActionNotFound = errors.New("Action not found")
	// ErrMissingRequiredParameter is for when required parameters are missing.
	ErrMissingRequiredParameter = errors.New("Required parameters missing")
	// ErrInvalidParameter is for when parameters are supplied that are not allowed.
	ErrInvalidParameter = errors.New("Invalid parameters")
)

// Parameter returns a parsed parameter or nil if no parameter exists. It is
// your responsibility to know the type.
func (c Command) Parameter(key string) interface{} {
	return c.parameters[key]
}

// Dispatch dispatches the Command, validating the parameters beforehand.
func (t Table) Dispatch(c Command) error {
	action, ok := t[c.Action]
	if !ok {
		return ErrActionNotFound
	}

	for param := range action.RequiredParameters {
		res, ok := c.Parameters[param]
		if !ok || res == nil {
			return ErrMissingRequiredParameter
		}
	}

	c.parameters = map[string]interface{}{}

	for key, initialValue := range c.Parameters {
		var found bool

		for param, coerceFunc := range action.RequiredParameters {
			if key == param {
				found = true

				if coerceFunc != nil {
					coerce := coerceFunc()
					if err := json.Unmarshal(initialValue, coerce); err != nil {
						return err
					}

					c.parameters[key] = coerce
				} else {
					return fmt.Errorf("please set a validation function for %q", key)
				}

				break
			}
		}

		if !found {
			for param, coerceFunc := range action.OptionalParameters {
				if key == param {
					found = true

					if coerceFunc != nil {
						coerce := coerceFunc()
						if err := json.Unmarshal(initialValue, coerce); err != nil {
							return err
						}

						c.parameters[param] = coerce
					} else {
						return fmt.Errorf("please set a validation function for %q", key)
					}

					break
				}
			}
		}

		if !found {
			return ErrInvalidParameter
		}
	}

	return action.Dispatch(c)
}
