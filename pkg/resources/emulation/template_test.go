package emulation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	spinregistry "github.com/erikh/spin/gen/spin_registry"
)

const testPath = "template-tests"

func TestRunTemplate(t *testing.T) {
	ac := AgentConfig{
		SystemDir:  "<systemdir>/",
		MonitorDir: "<monitordir>/",
	}

	infos, err := ioutil.ReadDir(testPath)
	if err != nil {
		t.Fatal(err)
	}

	for _, info := range infos {
		if filepath.Ext(info.Name()) == ".json" {
			content, err := ioutil.ReadFile(filepath.Join(testPath, info.Name()))
			if err != nil {
				t.Fatalf("Error reading %q: %v", info.Name(), err)
			}

			var vm spinregistry.UpdatedVM
			if err := json.Unmarshal(content, &vm); err != nil {
				t.Fatalf("Error parsing %q: %v", info.Name(), err)
			}

			tc, err := vmToTemplateConfig(ac, 1, &vm)
			if err != nil {
				t.Fatal(err)
			}

			tc.SpinQMP = "<spin-qmp>"

			result, err := runTemplate(tc)
			if err != nil {
				t.Fatal(err)
			}

			fn := fmt.Sprintf("%s.service", filepath.Join(testPath, strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))))
			if fi, err := os.Stat(fn); err != nil && os.IsNotExist(err) {
				if err := ioutil.WriteFile(fn, []byte(result), 0600); err != nil {
					t.Fatal(err)
				}

				t.Log("template generated, skipped this iteration")
				continue
			} else if err != nil {
				t.Fatal(err)
			} else if fi.Mode()&os.ModeType != 0 {
				t.Fatalf("%q result file was not a file", fi.Name())
			}

			content, err = ioutil.ReadFile(fn)
			if err != nil {
				t.Fatal(err)
			}

			if result != string(content) {
				t.Fatalf("%q mismatched on template generation:\n%s", info.Name(), diff.LineDiff(result, string(content)))
			}
		}
	}
}

func TestTemplateConfig(t *testing.T) {
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
