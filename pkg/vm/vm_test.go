package vm

import (
	"testing"
)

func TestValidation(t *testing.T) {
	table := map[string]struct {
		v interface {
			Validate() error
		}
		pass bool
	}{
		"core: empty struct": {
			v:    Core{},
			pass: false,
		},
		"core: basic": {
			v: Core{
				Name:   "foo",
				CPUs:   1,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: true,
		},
		"core: bad name 1": {
			v: Core{
				Name:   "foo\n",
				CPUs:   1,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: false,
		},
		"core: bad name 2": {
			v: Core{
				Name:   "",
				CPUs:   1,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: false,
		},
		"core: bad name 3": {
			v: Core{
				Name:   "-asdf",
				CPUs:   1,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: false,
		},
		"core: bad name 4": {
			v: Core{
				Name:   "0asdf",
				CPUs:   1,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: false,
		},
		"core: bad name 5": {
			v: Core{
				Name:   "name with spaces",
				CPUs:   1,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: false,
		},
		"core: good name 1": {
			v: Core{
				Name:   "this-name",
				CPUs:   1,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: true,
		},
		"core: good name 2": {
			v: Core{
				Name:   "THISNAME",
				CPUs:   1,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: true,
		},
		"core: bad cpus": {
			v: Core{
				Name:   "name",
				CPUs:   0,
				Memory: 1024,
				Ports:  PortMap{},
			},
			pass: false,
		},
		"core: bad memory": {
			v: Core{
				Name:   "name",
				CPUs:   1,
				Memory: 0,
				Ports:  PortMap{},
			},
			pass: false,
		},
		"portmap: good ports": {
			v:    PortMap{1234: "127.0.0.1:1234"},
			pass: true,
		},
		"portmap: bad ports 1": {
			v:    PortMap{1234: "localhost:1234"},
			pass: false,
		},
		"portmap: bad ports 2": {
			v:    PortMap{1234: "1234"},
			pass: false,
		},
		"portmap: bad ports 3": {
			v:    PortMap{1234: "127.0.0.1:"},
			pass: false,
		},
		"portmap: bad ports 4": {
			v:    PortMap{1234: "127.0.0.1:\n"},
			pass: false,
		},
		"storage: empty struct": {
			v:    Storage{},
			pass: false,
		},
		"storage: basic image": {
			v: Storage{
				Volume:    StringPtr("test"),
				Image:     "test.raw",
				ImageSize: UIntPtr(50),
			},
			pass: true,
		},
		"storage: basic cdrom": {
			v: Storage{
				Image: "test.img",
				CDROM: true,
			},
			pass: true,
		},
		"storage: cdrom: no image": {
			v: Storage{
				CDROM: true,
			},
			pass: false,
		},
		"storage: image: no volume": {
			v: Storage{
				Image:     "test.raw",
				ImageSize: UIntPtr(50),
			},
			pass: false,
		},
		"storage: image: no size": {
			v: Storage{
				Volume: StringPtr("test"),
				Image:  "test.raw",
			},
			pass: false,
		},
		"storage: cdrom: with volume": {
			v: Storage{
				Volume: StringPtr("test"),
				Image:  "test.raw",
				CDROM:  true,
			},
			pass: false,
		},
		"storage: cdrom: with image size": {
			v: Storage{
				Image:     "test.raw",
				CDROM:     true,
				ImageSize: UIntPtr(50),
			},
			pass: false,
		},
		"image: empty struct": {
			v:    Image{},
			pass: false,
		},
		"image: basic": {
			v: Image{
				Path:   "path/to/image.raw",
				Volume: StringPtr("test"),
			},
			pass: true,
		},
		"image: basic cdrom": {
			v: Image{
				Path:  "path/to/image.img",
				CDROM: true,
			},
			pass: true,
		},
		"image: cdrom: with volume": {
			v: Image{
				Path:   "path/to/image.img",
				CDROM:  true,
				Volume: StringPtr("test"),
			},
			pass: false,
		},
		"image: image: without volume": {
			v: Image{
				Path: "path/to/image.raw",
			},
			pass: false,
		},
		"image: image: without path": {
			v: Image{
				Volume: StringPtr("test"),
			},
			pass: false,
		},
		"image: cdrom: without path": {
			v: Image{
				CDROM: true,
			},
			pass: false,
		},
	}

	for name, test := range table {
		err := test.v.Validate()
		if test.pass && err != nil {
			t.Fatalf("Test %q failed with error: %v", name, err)
		}

		if !test.pass && err == nil {
			t.Fatalf("%q was not supposed to pass, and did.", name)
		}
	}
}
