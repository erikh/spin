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
	return s.db.VMCreate(p)
}

func (s *spinRegistry) VMUpdate(ctx context.Context, p *spinregistry.UpdateVM) error {
	return s.db.VMUpdate(p.ID, p.VM)
}

func (s *spinRegistry) VMDelete(ctx context.Context, p *spinregistry.VMDeletePayload) error {
	return s.db.VMDelete(p.ID)
}

func (s *spinRegistry) VMGet(ctx context.Context, p *spinregistry.VMGetPayload) (*spinregistry.VM, error) {
	return s.db.VMGet(p.ID)
}

func (s *spinRegistry) VMList(ctx context.Context) ([]uint64, error) {
	return s.db.VMList()
}

func (s *spinRegistry) StorageVolumesCreate(ctx context.Context, p *spinregistry.StorageVolumesCreatePayload) error {
	return s.db.StorageVolumeCreate(p.Name)
}

func (s *spinRegistry) StorageVolumesDelete(ctx context.Context, p *spinregistry.StorageVolumesDeletePayload) error {
	return s.db.StorageVolumeDelete(p.Name)
}

func (s *spinRegistry) StorageVolumesList(ctx context.Context) ([]string, error) {
	return s.db.StorageVolumeList()
}

func (s *spinRegistry) StorageImagesList(ctx context.Context, p *spinregistry.StorageImagesListPayload) ([]string, error) {
	return s.db.StorageImageList(p.VolumeName)
}

func (s *spinRegistry) StorageImagesCreate(ctx context.Context, p *spinregistry.Storage) error {
	return s.db.StorageImageCreate(p)
}

func (s *spinRegistry) StorageImagesDelete(ctx context.Context, p *spinregistry.StorageImagesDeletePayload) error {
	return s.db.StorageImageDelete(p.VolumeName, p.ImageName)
}

func (s *spinRegistry) StorageImagesGet(ctx context.Context, p *spinregistry.StorageImagesGetPayload) (*spinregistry.Storage, error) {
	return s.db.StorageImageGet(p.VolumeName, p.ImageName)
}
