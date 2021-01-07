package qmp

import (
	"context"
	"encoding/json"
	"net"
)

// Dial dials a QMP unix socket and performs the initial handshake.
func Dial(socket string) (net.Conn, error) {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, err
	}

	if err := json.NewEncoder(conn).Encode(map[string]interface{}{"execute": "qmp_capabilities"}); err != nil {
		conn.Close()
		return nil, err
	}

	obj := map[string]interface{}{}
	if err := json.NewDecoder(conn).Decode(&obj); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

// Shutdown issues a `system_powerdown` command to the connected QMP socket.
func Shutdown(ctx context.Context, conn net.Conn) error {
	if err := json.NewEncoder(conn).Encode(map[string]string{"execute": "system_powerdown"}); err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	for {
		obj := map[string]interface{}{}
		if err := json.NewDecoder(conn).Decode(&obj); err != nil {
			return err
		}
		if event, ok := obj["event"]; ok && event == "SHUTDOWN" {
			return nil
		}
	}
}
