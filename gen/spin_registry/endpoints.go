// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry endpoints
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package spinregistry

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "spin-registry" service endpoints.
type Endpoints struct {
	VMCreate             goa.Endpoint
	VMUpdate             goa.Endpoint
	VMDelete             goa.Endpoint
	VMGet                goa.Endpoint
	VMList               goa.Endpoint
	StorageVolumesList   goa.Endpoint
	StorageVolumesCreate goa.Endpoint
	StorageVolumesDelete goa.Endpoint
	StorageImagesList    goa.Endpoint
	StorageImagesCreate  goa.Endpoint
	StorageImagesDelete  goa.Endpoint
	StorageImagesGet     goa.Endpoint
}

// NewEndpoints wraps the methods of the "spin-registry" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		VMCreate:             NewVMCreateEndpoint(s),
		VMUpdate:             NewVMUpdateEndpoint(s),
		VMDelete:             NewVMDeleteEndpoint(s),
		VMGet:                NewVMGetEndpoint(s),
		VMList:               NewVMListEndpoint(s),
		StorageVolumesList:   NewStorageVolumesListEndpoint(s),
		StorageVolumesCreate: NewStorageVolumesCreateEndpoint(s),
		StorageVolumesDelete: NewStorageVolumesDeleteEndpoint(s),
		StorageImagesList:    NewStorageImagesListEndpoint(s),
		StorageImagesCreate:  NewStorageImagesCreateEndpoint(s),
		StorageImagesDelete:  NewStorageImagesDeleteEndpoint(s),
		StorageImagesGet:     NewStorageImagesGetEndpoint(s),
	}
}

// Use applies the given middleware to all the "spin-registry" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.VMCreate = m(e.VMCreate)
	e.VMUpdate = m(e.VMUpdate)
	e.VMDelete = m(e.VMDelete)
	e.VMGet = m(e.VMGet)
	e.VMList = m(e.VMList)
	e.StorageVolumesList = m(e.StorageVolumesList)
	e.StorageVolumesCreate = m(e.StorageVolumesCreate)
	e.StorageVolumesDelete = m(e.StorageVolumesDelete)
	e.StorageImagesList = m(e.StorageImagesList)
	e.StorageImagesCreate = m(e.StorageImagesCreate)
	e.StorageImagesDelete = m(e.StorageImagesDelete)
	e.StorageImagesGet = m(e.StorageImagesGet)
}

// NewVMCreateEndpoint returns an endpoint function that calls the method
// "vm_create" of service "spin-registry".
func NewVMCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatedVM)
		return s.VMCreate(ctx, p)
	}
}

// NewVMUpdateEndpoint returns an endpoint function that calls the method
// "vm_update" of service "spin-registry".
func NewVMUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateVM)
		return nil, s.VMUpdate(ctx, p)
	}
}

// NewVMDeleteEndpoint returns an endpoint function that calls the method
// "vm_delete" of service "spin-registry".
func NewVMDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*VMDeletePayload)
		return nil, s.VMDelete(ctx, p)
	}
}

// NewVMGetEndpoint returns an endpoint function that calls the method "vm_get"
// of service "spin-registry".
func NewVMGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*VMGetPayload)
		return s.VMGet(ctx, p)
	}
}

// NewVMListEndpoint returns an endpoint function that calls the method
// "vm_list" of service "spin-registry".
func NewVMListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.VMList(ctx)
	}
}

// NewStorageVolumesListEndpoint returns an endpoint function that calls the
// method "storage_volumes_list" of service "spin-registry".
func NewStorageVolumesListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.StorageVolumesList(ctx)
	}
}

// NewStorageVolumesCreateEndpoint returns an endpoint function that calls the
// method "storage_volumes_create" of service "spin-registry".
func NewStorageVolumesCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StorageVolumesCreatePayload)
		return nil, s.StorageVolumesCreate(ctx, p)
	}
}

// NewStorageVolumesDeleteEndpoint returns an endpoint function that calls the
// method "storage_volumes_delete" of service "spin-registry".
func NewStorageVolumesDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StorageVolumesDeletePayload)
		return nil, s.StorageVolumesDelete(ctx, p)
	}
}

// NewStorageImagesListEndpoint returns an endpoint function that calls the
// method "storage_images_list" of service "spin-registry".
func NewStorageImagesListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StorageImagesListPayload)
		return s.StorageImagesList(ctx, p)
	}
}

// NewStorageImagesCreateEndpoint returns an endpoint function that calls the
// method "storage_images_create" of service "spin-registry".
func NewStorageImagesCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Storage)
		return s.StorageImagesCreate(ctx, p)
	}
}

// NewStorageImagesDeleteEndpoint returns an endpoint function that calls the
// method "storage_images_delete" of service "spin-registry".
func NewStorageImagesDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StorageImagesDeletePayload)
		return nil, s.StorageImagesDelete(ctx, p)
	}
}

// NewStorageImagesGetEndpoint returns an endpoint function that calls the
// method "storage_images_get" of service "spin-registry".
func NewStorageImagesGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StorageImagesGetPayload)
		return s.StorageImagesGet(ctx, p)
	}
}
