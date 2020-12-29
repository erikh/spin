package spin

import (
	"context"
	"log"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
)

type spinRegistry struct {
	logger *log.Logger
}

func NewSpinRegistry(logger *log.Logger) spinregistry.Service {
	return &spinRegistry{logger}
}

func (s *spinRegistry) Create(ctx context.Context, p *spinregistry.VM) (uint64, error) {
	return 0, nil
}

func (s *spinRegistry) Update(ctx context.Context, p *spinregistry.UpdateVM) error {
	return nil
}

func (s *spinRegistry) Delete(ctx context.Context, p *spinregistry.DeletePayload) error {
	return nil
}

func (s *spinRegistry) Get(ctx context.Context, p *spinregistry.GetPayload) (*spinregistry.VM, error) {
	return nil, nil
}

func (s *spinRegistry) List(ctx context.Context) ([]uint64, error) {
	return nil, nil
}
