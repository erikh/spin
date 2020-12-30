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

func (s *spinRegistry) Create(ctx context.Context, p *spinregistry.VM) (uint64, error) {
	return s.db.Create(p)
}

func (s *spinRegistry) Update(ctx context.Context, p *spinregistry.UpdateVM) error {
	return s.db.Update(p.ID, p.VM)
}

func (s *spinRegistry) Delete(ctx context.Context, p *spinregistry.DeletePayload) error {
	return s.db.Delete(p.ID)
}

func (s *spinRegistry) Get(ctx context.Context, p *spinregistry.GetPayload) (*spinregistry.VM, error) {
	return s.db.Get(p.ID)
}

func (s *spinRegistry) List(ctx context.Context) ([]uint64, error) {
	return s.db.List()
}
