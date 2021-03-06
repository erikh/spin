package supervisor

import (
	"errors"
	"fmt"

	"github.com/coreos/go-systemd/v22/dbus"
)

type systemd struct {
	conn *dbus.Conn
}

// New creates a new systemd supervisor
func New() (Interface, error) {
	conn, err := dbus.NewUserConnection()
	if err != nil {
		return nil, err
	}

	return &systemd{conn: conn}, nil
}

// Review all configuration and start services as necessary
func (s *systemd) Reload() error {
	return s.conn.Reload()
}

// Start a service
func (s *systemd) Start(svc string) error {
	_, _, err := s.conn.EnableUnitFiles([]string{svc}, false, true)
	if err != nil {
		return err
	}

	ch := make(chan string, 1)
	if _, err := s.conn.StartUnit(svc, "replace", ch); err != nil {
		return err
	}

	switch res := <-ch; res {
	case "done":
		return nil
	default:
		return fmt.Errorf("systemd stop call returned a %q state", res)
	}
}

// Forcefully stop a service
func (s *systemd) Stop(svc string) error {
	ch := make(chan string, 1)
	_, err := s.conn.StopUnit(svc, "replace", ch)
	if err != nil {
		return err
	}

	switch res := <-ch; res {
	case "done":
		return nil
	default:
		return fmt.Errorf("systemd stop call returned a %q state", res)
	}
}

func (s *systemd) Running(svc string) (bool, error) {
	prop, err := s.conn.GetUnitProperty(svc, "ActiveState")
	if err != nil {
		return false, err
	}

	res, ok := prop.Value.Value().(string)
	if !ok {
		return false, errors.New("invalid result from systemd call")
	}

	return res == "active", nil
}
