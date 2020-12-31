package emulation

import (
	"strings"
	"testing"

	"github.com/andreyvit/diff"
)

func TestTemplate(t *testing.T) {
	result, err := runTemplate(templateConfig{
		ID:      1,
		Name:    "one",
		Command: "/bin/qemu-system-x86_64",
		Args:    []string{"-cpu", "host"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(result) != templates["one"] {
		t.Fatal("'one' content was not equal", diff.LineDiff(result, templates["one"]))
	}
}

var templates = map[string]string{
	// note that after the execstart there is one significant piece of
	// whitespace. that's just how the template works out right now.
	"one": `
[Unit]
Description=Virtual Machine #1: one

[Service]
Type=simple
ExecStart=/bin/qemu-system-x86_64 -cpu host 
TimeoutStopSec=30
KillSignal=SIGCONT
FinalKillSignal=SIGKILL

[Install]
WantedBy=default.target`,
}

func init() {
	for key := range templates {
		templates[key] = strings.TrimSpace(templates[key])
	}
}
