package emulation

import (
	"bytes"
	"fmt"
	"text/template"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
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

func vmToTemplateConfig(ac AgentConfig, id uint64, vm *spinregistry.VM) (templateConfig, error) {
	args := []string{
		"-nodefaults",
		"-chardev",
		fmt.Sprintf("socket,server,nowait,id=char0,path=%s", ac.monitorPath(id)),
		"-mon",
		"chardev=char0,mode=control,pretty=on",
		"-machine",
		"accel=kvm",
		"-vga",
		"virtio",
		"-m",
		fmt.Sprintf("%dM", vm.Memory),
		"-cpu",
		"kvm64",
		"-smp",
		fmt.Sprintf("cpus=1,cores=%d,maxcpus=%d", vm.Cpus, vm.Cpus),
		"-nic",
		"user",
	}

	for i, storage := range vm.Storage {
		if storage.Cdrom != nil && *storage.Cdrom {
			args = append(args,
				"-drive",
				fmt.Sprintf("file=%s,media=cdrom,index=%d", storage.Image, i),
			)
		} else {
			args = append(args,
				"-drive",
				fmt.Sprintf(
					"driver=raw,if=virtio,file=%s,cache=none,media=disk,index=%d",
					storage.Image, i,
				))
		}
	}

	tc := templateConfig{
		ID:      id,
		Name:    vm.Name,
		Command: qemuPath,
		Args:    args,
	}

	return tc, nil
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
