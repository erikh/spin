package spin

import (
	"context"
	"fmt"
	"log"
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

func toRegistryVM(p *spinapiserver.VM) *spinregistry.VM {
	storage := []*spinregistry.Storage{}

	for _, stor := range p.Storage {
		storage = append(storage, (*spinregistry.Storage)(stor))
	}

	return &spinregistry.VM{
		Name:    p.Name,
		Cpus:    p.Cpus,
		Memory:  p.Memory,
		Storage: storage,
	}
}

func (s *spinApiserversrvc) GetStatus(ctx context.Context, pkg string) error {
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

func (s *spinApiserversrvc) VMCreate(ctx context.Context, p *spinapiserver.VM) (uint64, error) {
	id, err := s.registry.VMCreate(ctx, toRegistryVM(p))
	if err != nil {
		return id, err
	}

	pkg, err := s.broker.New(ctx)
	if err != nil {
		return 0, err
	}

	storids := []string{}

	for _, stor := range p.Storage {
		if stor.Cdrom == nil || !*stor.Cdrom {
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
			"vm": p, // NOTE loose typing through JSON lets us get away with this, but this expects *spinregistry.VM.
		},
		Dependencies: storids,
	})
	if err != nil {
		return id, err
	}

	if _, err := s.broker.Enqueue(ctx, pkg); err != nil {
		return id, err
	}

	if err := s.GetStatus(ctx, pkg); err != nil {
		return id, err
	}

	for _, storage := range p.Storage {
		if storage.Cdrom == nil || !*storage.Cdrom {
			if err := s.registry.StorageImageCreate(ctx, (*spinregistry.Storage)(storage)); err != nil {
				return id, err
			}
		}
	}

	return id, nil
}

func (s *spinApiserversrvc) VMDelete(ctx context.Context, p *spinapiserver.VMDeletePayload) error {
	return nil
}

// ControlStart implements control/start.
func (s *spinApiserversrvc) ControlStart(ctx context.Context, p *spinapiserver.ControlStartPayload) (err error) {
	pkg, err := s.broker.New(ctx)
	if err != nil {
		return err
	}

	_, err = s.broker.Add(ctx, &spinbroker.AddPayload{
		ID:       pkg,
		Resource: "emulation",
		Action:   "start",
		Parameters: map[string]interface{}{
			"id": p.ID,
		},
	})
	if err != nil {
		return err
	}

	if _, err := s.broker.Enqueue(ctx, pkg); err != nil {
		return err
	}

	return s.GetStatus(ctx, pkg)
}

// ControlStop implements control/stop.
func (s *spinApiserversrvc) ControlStop(ctx context.Context, p *spinapiserver.ControlStopPayload) (err error) {
	return nil
}

// ControlShutdown implements control/shutdown.
func (s *spinApiserversrvc) ControlShutdown(ctx context.Context, p *spinapiserver.ControlShutdownPayload) (err error) {
	return nil
}
