// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry HTTP server types
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package server

import (
	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "spin-registry" service "create"
// endpoint HTTP request body.
type CreateRequestBody struct {
	// Name of VM; does not need to be unique
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// CPU count
	Cpus *uint `form:"cpus,omitempty" json:"cpus,omitempty" xml:"cpus,omitempty"`
	// Memory (in megabytes)
	Memory *uint `form:"memory,omitempty" json:"memory,omitempty" xml:"memory,omitempty"`
	// Storage references
	Storage []*StorageRequestBody `form:"storage,omitempty" json:"storage,omitempty" xml:"storage,omitempty"`
}

// UpdateRequestBody is the type of the "spin-registry" service "update"
// endpoint HTTP request body.
type UpdateRequestBody struct {
	// VM to publish
	VM *VMRequestBody `form:"vm,omitempty" json:"vm,omitempty" xml:"vm,omitempty"`
}

// GetResponseBody is the type of the "spin-registry" service "get" endpoint
// HTTP response body.
type GetResponseBody struct {
	// Name of VM; does not need to be unique
	Name string `form:"name" json:"name" xml:"name"`
	// CPU count
	Cpus uint `form:"cpus" json:"cpus" xml:"cpus"`
	// Memory (in megabytes)
	Memory uint `form:"memory" json:"memory" xml:"memory"`
	// Storage references
	Storage []*StorageResponseBody `form:"storage" json:"storage" xml:"storage"`
}

// StorageResponseBody is used to define fields on response body types.
type StorageResponseBody struct {
	// Volume name, must not include `/`
	Volume string `form:"volume" json:"volume" xml:"volume"`
	// Image filename, must not include `/`
	Image string `form:"image" json:"image" xml:"image"`
	// Image size (in gigabytes)
	ImageSize uint `form:"image_size" json:"image_size" xml:"image_size"`
}

// StorageRequestBody is used to define fields on request body types.
type StorageRequestBody struct {
	// Volume name, must not include `/`
	Volume *string `form:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
	// Image filename, must not include `/`
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// Image size (in gigabytes)
	ImageSize *uint `form:"image_size,omitempty" json:"image_size,omitempty" xml:"image_size,omitempty"`
}

// VMRequestBody is used to define fields on request body types.
type VMRequestBody struct {
	// Name of VM; does not need to be unique
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// CPU count
	Cpus *uint `form:"cpus,omitempty" json:"cpus,omitempty" xml:"cpus,omitempty"`
	// Memory (in megabytes)
	Memory *uint `form:"memory,omitempty" json:"memory,omitempty" xml:"memory,omitempty"`
	// Storage references
	Storage []*StorageRequestBody `form:"storage,omitempty" json:"storage,omitempty" xml:"storage,omitempty"`
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "spin-registry" service.
func NewGetResponseBody(res *spinregistry.VM) *GetResponseBody {
	body := &GetResponseBody{
		Name:   res.Name,
		Cpus:   res.Cpus,
		Memory: res.Memory,
	}
	if res.Storage != nil {
		body.Storage = make([]*StorageResponseBody, len(res.Storage))
		for i, val := range res.Storage {
			body.Storage[i] = marshalSpinregistryStorageToStorageResponseBody(val)
		}
	}
	return body
}

// NewCreateVM builds a spin-registry service create endpoint payload.
func NewCreateVM(body *CreateRequestBody) *spinregistry.VM {
	v := &spinregistry.VM{
		Name:   *body.Name,
		Cpus:   *body.Cpus,
		Memory: *body.Memory,
	}
	v.Storage = make([]*spinregistry.Storage, len(body.Storage))
	for i, val := range body.Storage {
		v.Storage[i] = unmarshalStorageRequestBodyToSpinregistryStorage(val)
	}

	return v
}

// NewUpdateVM builds a spin-registry service update endpoint payload.
func NewUpdateVM(body *UpdateRequestBody, id uint64) *spinregistry.UpdateVM {
	v := &spinregistry.UpdateVM{}
	v.VM = unmarshalVMRequestBodyToSpinregistryVM(body.VM)
	v.ID = id

	return v
}

// NewDeletePayload builds a spin-registry service delete endpoint payload.
func NewDeletePayload(id uint64) *spinregistry.DeletePayload {
	v := &spinregistry.DeletePayload{}
	v.ID = id

	return v
}

// NewGetPayload builds a spin-registry service get endpoint payload.
func NewGetPayload(id uint64) *spinregistry.GetPayload {
	v := &spinregistry.GetPayload{}
	v.ID = id

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Cpus == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cpus", "body"))
	}
	if body.Memory == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("memory", "body"))
	}
	if body.Storage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("storage", "body"))
	}
	for _, e := range body.Storage {
		if e != nil {
			if err2 := ValidateStorageRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.VM == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vm", "body"))
	}
	if body.VM != nil {
		if err2 := ValidateVMRequestBody(body.VM); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStorageRequestBody runs the validations defined on StorageRequestBody
func ValidateStorageRequestBody(body *StorageRequestBody) (err error) {
	if body.Volume == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("volume", "body"))
	}
	if body.Image == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image", "body"))
	}
	if body.ImageSize == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image_size", "body"))
	}
	return
}

// ValidateVMRequestBody runs the validations defined on VMRequestBody
func ValidateVMRequestBody(body *VMRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Cpus == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cpus", "body"))
	}
	if body.Memory == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("memory", "body"))
	}
	if body.Storage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("storage", "body"))
	}
	for _, e := range body.Storage {
		if e != nil {
			if err2 := ValidateStorageRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
