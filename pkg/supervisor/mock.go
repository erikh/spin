package supervisor

// Mock mocks the supervisor
type Mock struct {
	Reloads   uint
	Starts    map[string]uint
	Stops     map[string]uint
	errorNext error
}

// NewMock instantiates the mock properly.
func NewMock() *Mock {
	return &Mock{
		Starts: map[string]uint{},
		Stops:  map[string]uint{},
	}
}

// ErrorNext sets the error for the next call to return. The error is cleared
// after returning the error the first time.
func (m *Mock) ErrorNext(e error) {
	m.errorNext = e
}

func (m *Mock) returnErr() error {
	if m.errorNext != nil {
		err := m.errorNext
		m.errorNext = nil
		return err
	}

	return nil
}

// Reload reviews all configuration and starts services as necessary
func (m *Mock) Reload() error {
	if err := m.returnErr(); err != nil {
		return err
	}

	m.Reloads++
	return nil
}

// Start a service
func (m *Mock) Start(svc string) error {
	if err := m.returnErr(); err != nil {
		return err
	}

	m.Starts[svc]++
	return nil
}

// Stop a service
func (m *Mock) Stop(svc string) error {
	if err := m.returnErr(); err != nil {
		return err
	}

	m.Stops[svc]++
	return nil
}
