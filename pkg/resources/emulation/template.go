package emulation

import (
	"bytes"
	"text/template"
)

const systemdUnit = ` 
[Unit]
Description=Virtual Machine #{{ .ID }}: {{ .Name }}

[Service]
Type=simple
ExecStart={{ .Command }} {{ range $value := .Args }}{{ $value }} {{ end }}
TimeoutStopSec=30
KillSignal=SIGCONT
FinalKillSignal=SIGKILL

[Install]
WantedBy=default.target
`

type templateConfig struct {
	ID      uint64
	Name    string
	Command string
	Args    []string
}

func runTemplate(tc templateConfig) (string, error) { // nolint
	t, err := template.New("systemd-unit").Parse(systemdUnit)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)

	if err := t.Execute(buf, tc); err != nil {
		return "", err
	}

	return buf.String(), nil
}
