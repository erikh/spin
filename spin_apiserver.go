package spin

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	brokerclient "code.hollensbe.org/erikh/spin/clients/broker"
	registryclient "code.hollensbe.org/erikh/spin/clients/registry"
	spinapiserver "code.hollensbe.org/erikh/spin/gen/spin_apiserver"
	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	goa "goa.design/goa/v3/pkg"
)

type spinApiserversrvc struct {
	logger   *log.Logger
	broker   *brokerclient.Client
	registry *registryclient.Client
}

// NewSpinApiserver returns the spin-apiserver service implementation.
func NewSpinApiserver(logger *log.Logger, broker *brokerclient.Client, registry *registryclient.Client) spinapiserver.Service {
	return &spinApiserversrvc{
		logger:   logger,
		broker:   broker,
		registry: registry,
	}
}

func toRegistryVM(p *spinapiserver.CreateVM, images []*spinregistry.Image) *spinregistry.UpdatedVM {
	return &spinregistry.UpdatedVM{
		Name:   p.Name,
		Cpus:   p.Cpus,
		Memory: p.Memory,
		Images: images,
	}
}

func (s *spinApiserversrvc) getStatus(ctx context.Context, pkg string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		status, err := s.broker.Status(ctx, pkg)
		if err != nil {
			if e, ok := err.(*goa.ServiceError); ok && e.ErrorName() == "record_not_found" {
				time.Sleep(100 * time.Millisecond)
				continue
			}

			return err
		}

		if !status.Status {
			return fmt.Errorf("[%v]: %q", *status.Causer, *status.Reason)
		}

		break
	}

	return nil
}

// this call executes entire packages as queued subroutines. it automatically
// overwrites the ID for each add so that the package is correct. No accounting
// for dependencies are made, so use this call wisely.
func (s *spinApiserversrvc) apiOneShot(ctx context.Context, adds ...*spinbroker.AddPayload) error {
	pkg, err := s.broker.New(ctx)
	if err != nil {
		return err
	}

	for _, add := range adds {
		add.ID = pkg
		_, err = s.broker.Add(ctx, add)
		if err != nil {
			return err
		}
	}

	if _, err := s.broker.Enqueue(ctx, pkg); err != nil {
		return err
	}

	return s.getStatus(ctx, pkg)
}

func (s *spinApiserversrvc) VMGet(ctx context.Context, p *spinapiserver.VMGetPayload) (*spinapiserver.UpdatedVM, error) {
	vm, err := s.registry.VMGet(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	ret := &spinapiserver.UpdatedVM{
		Name:   vm.Name,
		Cpus:   vm.Cpus,
		Memory: vm.Memory,
	}

	for _, image := range vm.Images {
		ret.Images = append(ret.Images, (*spinapiserver.Image)(image))
	}

	return ret, nil
}
func (s *spinApiserversrvc) VMList(ctx context.Context) ([]uint64, error) {
	return s.registry.VMList(ctx)
}

func (s *spinApiserversrvc) VMCreate(ctx context.Context, p *spinapiserver.CreateVM) (uint64, error) {
	images := []*spinregistry.Image{}

	for _, storage := range p.Storage {
		if !storage.Cdrom && strings.Contains(storage.Image, "/") {
			return 0, errors.New("images cannot contain path components")
		}
	}

	for _, storage := range p.Storage {
		if storage.Cdrom {
			images = append(images, &spinregistry.Image{
				Path:  storage.Image,
				Cdrom: true,
			})
		} else {
			img, err := s.registry.StorageImageCreate(ctx, (*spinregistry.Storage)(storage))
			if err != nil {
				return 0, err
			}

			images = append(images, img)
		}
	}

	vm := toRegistryVM(p, images)
	id, err := s.registry.VMCreate(ctx, vm)
	if err != nil {
		return id, err
	}

	pkg, err := s.broker.New(ctx)
	if err != nil {
		return 0, err
	}

	storids := []string{}

	for _, stor := range p.Storage {
		if !stor.Cdrom {
			storid, err := s.broker.Add(ctx, &spinbroker.AddPayload{
				ID:       pkg,
				Resource: "storage",
				Action:   "create_image",
				Parameters: map[string]interface{}{
					"volume_path": stor.Volume,
					"image_name":  stor.Image,
					"image_size":  *stor.ImageSize,
				},
			})
			if err != nil {
				return id, err
			}

			storids = append(storids, storid)
		}
	}

	_, err = s.broker.Add(ctx, &spinbroker.AddPayload{
		ID:       pkg,
		Resource: "emulation",
		Action:   "write_config",
		Parameters: map[string]interface{}{
			"id": id,
			"vm": vm,
		},
		Dependencies: storids,
	})
	if err != nil {
		return id, err
	}

	if _, err := s.broker.Enqueue(ctx, pkg); err != nil {
		return id, err
	}

	if err := s.getStatus(ctx, pkg); err != nil {
		return id, err
	}

	return id, nil
}

func (s *spinApiserversrvc) VMDelete(ctx context.Context, p *spinapiserver.VMDeletePayload) error {
	vm, err := s.registry.VMGet(ctx, p.ID)
	if err != nil {
		return err
	}

	pkg, err := s.broker.New(ctx)
	if err != nil {
		return err
	}

	vmid, err := s.broker.Add(ctx, &spinbroker.AddPayload{
		ID:       pkg,
		Resource: "emulation",
		Action:   "stop",
		Parameters: map[string]interface{}{
			"id": p.ID,
		},
	})
	if err != nil {
		return err
	}

	storids := []string{}

	for _, stor := range vm.Images {
		if !stor.Cdrom {
			uuid, err := s.broker.Add(ctx, &spinbroker.AddPayload{
				ID:       pkg,
				Resource: "storage",
				Action:   "delete_image",
				Parameters: map[string]interface{}{
					"image_path": stor.Path,
				},
				Dependencies: []string{vmid},
			})
			if err != nil {
				return err
			}

			storids = append(storids, uuid)
		}
	}

	_, err = s.broker.Add(ctx, &spinbroker.AddPayload{
		ID:       pkg,
		Resource: "emulation",
		Action:   "remove_config",
		Parameters: map[string]interface{}{
			"id": p.ID,
		},
		Dependencies: append([]string{vmid}, storids...),
	})
	if err != nil {
		return err
	}

	if _, err := s.broker.Enqueue(ctx, pkg); err != nil {
		return err
	}

	if err := s.getStatus(ctx, pkg); err != nil {
		return err
	}

	for _, stor := range vm.Images {
		if !stor.Cdrom {
			if err := s.registry.StorageImageDelete(ctx, *stor.Volume, filepath.Base(stor.Path)); err != nil {
				return err
			}
		}
	}

	return s.registry.VMDelete(ctx, p.ID)
}

// ControlStart implements control/start.
func (s *spinApiserversrvc) ControlStart(ctx context.Context, p *spinapiserver.ControlStartPayload) (err error) {
	return s.apiOneShot(ctx, &spinbroker.AddPayload{
		Resource: "emulation",
		Action:   "start",
		Parameters: map[string]interface{}{
			"id": p.ID,
		},
	})
}

// ControlStop implements control/stop.
func (s *spinApiserversrvc) ControlStop(ctx context.Context, p *spinapiserver.ControlStopPayload) (err error) {
	return s.apiOneShot(ctx, &spinbroker.AddPayload{
		Resource: "emulation",
		Action:   "stop",
		Parameters: map[string]interface{}{
			"id": p.ID,
		},
	})
}

// ControlShutdown implements control/shutdown.
func (s *spinApiserversrvc) ControlShutdown(ctx context.Context, p *spinapiserver.ControlShutdownPayload) (err error) {
	return s.apiOneShot(ctx, &spinbroker.AddPayload{
		Resource: "emulation",
		Action:   "shutdown",
		Parameters: map[string]interface{}{
			"id": p.ID,
		},
	})
}
