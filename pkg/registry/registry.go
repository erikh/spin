package registry

import "errors"

type Registry struct {
	db *DB
}

func New(dbc DBConfig) (*Registry, error) {
	db, err := NewDB(dbc)
	if err != nil {
		return nil, err
	}

	return &Registry{db: db}, nil
}

func (r *Registry) Get(id uint64) (*VM, error) {
	return nil, errors.New("unimplemented")
}

func (r *Registry) Create(vm *VM) error {
	return errors.New("unimplemented")
}

func (r *Registry) Delete(id uint64) error {
	return errors.New("unimplemented")
}

func (r *Registry) Update(id uint64, vm *VM) error {
	return errors.New("unimplemented")
}

func (r *Registry) List() ([]uint64, error) {
	return nil, errors.New("unimplemented")
}
