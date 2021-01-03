package spin

import (
	"context"
	"log"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	"code.hollensbe.org/erikh/spin/pkg/registry"
)

type spinRegistry struct {
	logger *log.Logger
	db     *registry.DB
}

// NewSpinRegistry creates a new spin registry
func NewSpinRegistry(logger *log.Logger, db *registry.DB) spinregistry.Service {
	return &spinRegistry{logger, db}
}

func (s *spinRegistry) VMCreate(ctx context.Context, p *spinregistry.VM) (uint64, error) {
	return s.db.Create(p)
}

func (s *spinRegistry) VMUpdate(ctx context.Context, p *spinregistry.UpdateVM) error {
	return s.db.Update(p.ID, p.VM)
}

func (s *spinRegistry) VMDelete(ctx context.Context, p *spinregistry.VMDeletePayload) error {
	return s.db.Delete(p.ID)
}

func (s *spinRegistry) VMGet(ctx context.Context, p *spinregistry.VMGetPayload) (*spinregistry.VM, error) {
	return s.db.Get(p.ID)
}

func (s *spinRegistry) VMList(ctx context.Context) ([]uint64, error) {
	return s.db.List()
}
