package emulation

import (
	"bytes"
	"fmt"
	"text/template"

	spinregistry "github.com/erikh/spin/gen/spin_registry"
	"github.com/mitchellh/go-homedir"
)

const systemdUnit = ` 
[Unit]
Description=Virtual Machine #{{ .ID }}: {{ .Name }}

[Service]
Type=simple
ExecStart={{ .Command }} {{ range $value := .Args }}{{ $value }} {{ end }}
ExecStop={{ .SpinQMP }} shutdown {{ .Home }}/.config/spin/monitors/{{ .ID }}
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
	SpinQMP string
	Home    string
}

func vmToTemplateConfig(ac AgentConfig, id uint64, vm *spinregistry.UpdatedVM) (templateConfig, error) {
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
		"-vnc",
		fmt.Sprintf("127.0.0.1:0,websocket=127.0.0.1:60%02d", id%100),
		"-m",
		fmt.Sprintf("%dM", vm.Memory),
		"-cpu",
		"kvm64",
		"-smp",
		fmt.Sprintf("cpus=1,cores=%d,maxcpus=%d", vm.Cpus, vm.Cpus),
		"-nic",
		fmt.Sprintf("user%s", hostfwdRules(vm)),
	}

	for i, storage := range vm.Images {
		if storage.Cdrom {
			args = append(args,
				"-drive",
				fmt.Sprintf("file=%s,media=cdrom,index=%d", storage.Path, i),
			)
		} else {
			args = append(args,
				"-drive",
				fmt.Sprintf(
					"driver=raw,if=virtio,file=%s,cache=none,media=disk,index=%d",
					storage.Path, i,
				))
		}
	}

	dir, err := homedir.Dir()
	if err != nil {
		return templateConfig{}, err
	}

	tc := templateConfig{
		ID:      id,
		Name:    vm.Name,
		Command: qemuPath,
		Args:    args,
		SpinQMP: spinQMPBin,
		Home:    dir,
	}

	return tc, nil
}

func hostfwdRules(vm *spinregistry.UpdatedVM) string {
	str := ""

	for guest, hostaddr := range vm.Ports {
		str += fmt.Sprintf(",hostfwd=tcp:%s-:%d", hostaddr, guest)
	}

	return str
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
