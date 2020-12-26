package broker

import "errors"

// Dispatcher encapsulates a dispatching system that consists of actions
// (strings) that correspond to processed/validated properties and a processing
// function.
type Dispatcher map[string]Action

// Action is the definition of the protocol action item.
type Action struct {
	RequiredParameters []string
	OptionalParameters []string
	Dispatch           func(Command) error
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
func (d Dispatcher) Dispatch(c Command) error {
	action, ok := d[c.Action]
	if !ok {
		return ErrActionNotFound
	}

	for _, param := range action.RequiredParameters {
		res, ok := c.Parameters[param]
		if !ok || res == "" {
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
