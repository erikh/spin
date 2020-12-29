package dispatcher

import "errors"

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
	Parameters   map[string]interface{}
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

// Action is the definition of the protocol action item.
type Action struct {
	RequiredParameters []string
	OptionalParameters []string
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

// Dispatch dispatches the Command, validating the parameters beforehand.
func (t Table) Dispatch(c Command) error {
	action, ok := t[c.Action]
	if !ok {
		return ErrActionNotFound
	}

	for _, param := range action.RequiredParameters {
		res, ok := c.Parameters[param]
		if !ok || res == nil {
			return ErrMissingRequiredParameter
		}
	}

	for key := range c.Parameters {
		var found bool

		for _, param := range action.RequiredParameters {
			if key == param {
				found = true
				break
			}
		}

		if !found {
			for _, param := range action.OptionalParameters {
				if key == param {
					found = true
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
