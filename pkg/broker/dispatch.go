package broker

import "errors"

type Dispatcher map[string]Action

type Action struct {
	RequiredParameters []string
	OptionalParameters []string
	Dispatch           func(Command) error
}

var (
	ErrActionNotFound           = errors.New("Action not found")
	ErrMissingRequiredParameter = errors.New("Required parameters missing")
	ErrInvalidParameter         = errors.New("Invalid parameters")
)

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
