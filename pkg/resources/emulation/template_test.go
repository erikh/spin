package emulation

import (
	"strings"
	"testing"

	"github.com/andreyvit/diff"
)

func TestTemplate(t *testing.T) {
	for key, tc := range templateInputs {
		result, err := runTemplate(tc)
		if err != nil {
			t.Fatal(err)
		}

		if strings.TrimSpace(result) != strings.TrimSpace(templateResults[key]) {
			t.Fatalf("%q content was not equal:\n%s", key, diff.LineDiff(result, templateResults[key]))
		}
	}
}

var templateInputs = map[string]templateConfig{
	"one": {
		ID:      1,
		Name:    "one",
		Command: "/bin/qemu-system-x86_64",
		SpinQMP: "/bin/spin-qmp",
		Args:    []string{"-cpu", "host"},
		Home:    "/home/erikh",
	},
}

var templateResults = map[string]string{
	// note that after the execstart there is one significant piece of
	// whitespace. that's just how the template works out right now.
	"one": `
[Unit]
Description=Virtual Machine #1: one

[Service]
Type=simple
ExecStart=/bin/qemu-system-x86_64 -cpu host 
ExecStop=/bin/spin-qmp shutdown /home/erikh/.config/spin/monitors/1
TimeoutStopSec=30
KillSignal=SIGCONT
FinalKillSignal=SIGKILL

[Install]
WantedBy=default.target`,
}
