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

// VMCreateRequestBody is the type of the "spin-registry" service "vm/create"
// endpoint HTTP request body.
type VMCreateRequestBody struct {
	// Name of VM; does not need to be unique
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// CPU count
	Cpus *uint `form:"cpus,omitempty" json:"cpus,omitempty" xml:"cpus,omitempty"`
	// Memory (in megabytes)
	Memory *uint `form:"memory,omitempty" json:"memory,omitempty" xml:"memory,omitempty"`
	// Storage references
	Storage []*StorageRequestBody `form:"storage,omitempty" json:"storage,omitempty" xml:"storage,omitempty"`
}

// VMUpdateRequestBody is the type of the "spin-registry" service "vm/update"
// endpoint HTTP request body.
type VMUpdateRequestBody struct {
	// VM to publish
	VM *VMRequestBody `form:"vm,omitempty" json:"vm,omitempty" xml:"vm,omitempty"`
}

// StorageVolumesCreateRequestBody is the type of the "spin-registry" service
// "storage/volumes/create" endpoint HTTP request body.
type StorageVolumesCreateRequestBody struct {
	// name of volume
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// StorageVolumesDeleteRequestBody is the type of the "spin-registry" service
// "storage/volumes/delete" endpoint HTTP request body.
type StorageVolumesDeleteRequestBody struct {
	// name of volume
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// StorageImagesListRequestBody is the type of the "spin-registry" service
// "storage/images/list" endpoint HTTP request body.
type StorageImagesListRequestBody struct {
	// name of volume to list images for
	VolumeName *string `form:"volume_name,omitempty" json:"volume_name,omitempty" xml:"volume_name,omitempty"`
}

// StorageImagesCreateRequestBody is the type of the "spin-registry" service
// "storage/images/create" endpoint HTTP request body.
type StorageImagesCreateRequestBody struct {
	// Volume name, must not include `/`
	Volume *string `form:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
	// Image filename, must not include `/`
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// Image size (in gigabytes)
	ImageSize *uint `form:"image_size,omitempty" json:"image_size,omitempty" xml:"image_size,omitempty"`
	// Is this image a cdrom?
	Cdrom *bool `form:"cdrom,omitempty" json:"cdrom,omitempty" xml:"cdrom,omitempty"`
}

// StorageImagesDeleteRequestBody is the type of the "spin-registry" service
// "storage/images/delete" endpoint HTTP request body.
type StorageImagesDeleteRequestBody struct {
	// name of volume
	VolumeName *string `form:"volume_name,omitempty" json:"volume_name,omitempty" xml:"volume_name,omitempty"`
	// name of image
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
}

// StorageImagesGetRequestBody is the type of the "spin-registry" service
// "storage/images/get" endpoint HTTP request body.
type StorageImagesGetRequestBody struct {
	// name of volume
	VolumeName *string `form:"volume_name,omitempty" json:"volume_name,omitempty" xml:"volume_name,omitempty"`
	// name of image
	ImageName *string `form:"image_name,omitempty" json:"image_name,omitempty" xml:"image_name,omitempty"`
}

// VMGetResponseBody is the type of the "spin-registry" service "vm/get"
// endpoint HTTP response body.
type VMGetResponseBody struct {
	// Name of VM; does not need to be unique
	Name string `form:"name" json:"name" xml:"name"`
	// CPU count
	Cpus uint `form:"cpus" json:"cpus" xml:"cpus"`
	// Memory (in megabytes)
	Memory uint `form:"memory" json:"memory" xml:"memory"`
	// Storage references
	Storage []*StorageResponseBody `form:"storage" json:"storage" xml:"storage"`
}

// StorageImagesGetResponseBody is the type of the "spin-registry" service
// "storage/images/get" endpoint HTTP response body.
type StorageImagesGetResponseBody struct {
	// Volume name, must not include `/`
	Volume string `form:"volume" json:"volume" xml:"volume"`
	// Image filename, must not include `/`
	Image string `form:"image" json:"image" xml:"image"`
	// Image size (in gigabytes)
	ImageSize *uint `form:"image_size,omitempty" json:"image_size,omitempty" xml:"image_size,omitempty"`
	// Is this image a cdrom?
	Cdrom *bool `form:"cdrom,omitempty" json:"cdrom,omitempty" xml:"cdrom,omitempty"`
}

// StorageResponseBody is used to define fields on response body types.
type StorageResponseBody struct {
	// Volume name, must not include `/`
	Volume string `form:"volume" json:"volume" xml:"volume"`
	// Image filename, must not include `/`
	Image string `form:"image" json:"image" xml:"image"`
	// Image size (in gigabytes)
	ImageSize *uint `form:"image_size,omitempty" json:"image_size,omitempty" xml:"image_size,omitempty"`
	// Is this image a cdrom?
	Cdrom *bool `form:"cdrom,omitempty" json:"cdrom,omitempty" xml:"cdrom,omitempty"`
}

// StorageRequestBody is used to define fields on request body types.
type StorageRequestBody struct {
	// Volume name, must not include `/`
	Volume *string `form:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
	// Image filename, must not include `/`
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// Image size (in gigabytes)
	ImageSize *uint `form:"image_size,omitempty" json:"image_size,omitempty" xml:"image_size,omitempty"`
	// Is this image a cdrom?
	Cdrom *bool `form:"cdrom,omitempty" json:"cdrom,omitempty" xml:"cdrom,omitempty"`
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

// NewVMGetResponseBody builds the HTTP response body from the result of the
// "vm/get" endpoint of the "spin-registry" service.
func NewVMGetResponseBody(res *spinregistry.VM) *VMGetResponseBody {
	body := &VMGetResponseBody{
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

// NewStorageImagesGetResponseBody builds the HTTP response body from the
// result of the "storage/images/get" endpoint of the "spin-registry" service.
func NewStorageImagesGetResponseBody(res *spinregistry.Storage) *StorageImagesGetResponseBody {
	body := &StorageImagesGetResponseBody{
		Volume:    res.Volume,
		Image:     res.Image,
		ImageSize: res.ImageSize,
		Cdrom:     res.Cdrom,
	}
	return body
}

// NewVMCreateVM builds a spin-registry service vm/create endpoint payload.
func NewVMCreateVM(body *VMCreateRequestBody) *spinregistry.VM {
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

// NewVMUpdateUpdateVM builds a spin-registry service vm/update endpoint
// payload.
func NewVMUpdateUpdateVM(body *VMUpdateRequestBody, id uint64) *spinregistry.UpdateVM {
	v := &spinregistry.UpdateVM{}
	v.VM = unmarshalVMRequestBodyToSpinregistryVM(body.VM)
	v.ID = id

	return v
}

// NewVMDeletePayload builds a spin-registry service vm/delete endpoint payload.
func NewVMDeletePayload(id uint64) *spinregistry.VMDeletePayload {
	v := &spinregistry.VMDeletePayload{}
	v.ID = id

	return v
}

// NewVMGetPayload builds a spin-registry service vm/get endpoint payload.
func NewVMGetPayload(id uint64) *spinregistry.VMGetPayload {
	v := &spinregistry.VMGetPayload{}
	v.ID = id

	return v
}

// NewStorageVolumesCreatePayload builds a spin-registry service
// storage/volumes/create endpoint payload.
func NewStorageVolumesCreatePayload(body *StorageVolumesCreateRequestBody) *spinregistry.StorageVolumesCreatePayload {
	v := &spinregistry.StorageVolumesCreatePayload{
		Name: *body.Name,
	}

	return v
}

// NewStorageVolumesDeletePayload builds a spin-registry service
// storage/volumes/delete endpoint payload.
func NewStorageVolumesDeletePayload(body *StorageVolumesDeleteRequestBody) *spinregistry.StorageVolumesDeletePayload {
	v := &spinregistry.StorageVolumesDeletePayload{
		Name: *body.Name,
	}

	return v
}

// NewStorageImagesListPayload builds a spin-registry service
// storage/images/list endpoint payload.
func NewStorageImagesListPayload(body *StorageImagesListRequestBody) *spinregistry.StorageImagesListPayload {
	v := &spinregistry.StorageImagesListPayload{
		VolumeName: *body.VolumeName,
	}

	return v
}

// NewStorageImagesCreateStorage builds a spin-registry service
// storage/images/create endpoint payload.
func NewStorageImagesCreateStorage(body *StorageImagesCreateRequestBody) *spinregistry.Storage {
	v := &spinregistry.Storage{
		Volume:    *body.Volume,
		Image:     *body.Image,
		ImageSize: body.ImageSize,
		Cdrom:     body.Cdrom,
	}

	return v
}

// NewStorageImagesDeletePayload builds a spin-registry service
// storage/images/delete endpoint payload.
func NewStorageImagesDeletePayload(body *StorageImagesDeleteRequestBody) *spinregistry.StorageImagesDeletePayload {
	v := &spinregistry.StorageImagesDeletePayload{
		VolumeName: *body.VolumeName,
		ImageName:  *body.ImageName,
	}

	return v
}

// NewStorageImagesGetPayload builds a spin-registry service storage/images/get
// endpoint payload.
func NewStorageImagesGetPayload(body *StorageImagesGetRequestBody) *spinregistry.StorageImagesGetPayload {
	v := &spinregistry.StorageImagesGetPayload{
		VolumeName: *body.VolumeName,
		ImageName:  *body.ImageName,
	}

	return v
}

// ValidateVMCreateRequestBody runs the validations defined on
// Vm/CreateRequestBody
func ValidateVMCreateRequestBody(body *VMCreateRequestBody) (err error) {
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

// ValidateVMUpdateRequestBody runs the validations defined on
// Vm/UpdateRequestBody
func ValidateVMUpdateRequestBody(body *VMUpdateRequestBody) (err error) {
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

// ValidateStorageVolumesCreateRequestBody runs the validations defined on
// Storage/Volumes/CreateRequestBody
func ValidateStorageVolumesCreateRequestBody(body *StorageVolumesCreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateStorageVolumesDeleteRequestBody runs the validations defined on
// Storage/Volumes/DeleteRequestBody
func ValidateStorageVolumesDeleteRequestBody(body *StorageVolumesDeleteRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateStorageImagesListRequestBody runs the validations defined on
// Storage/Images/ListRequestBody
func ValidateStorageImagesListRequestBody(body *StorageImagesListRequestBody) (err error) {
	if body.VolumeName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("volume_name", "body"))
	}
	return
}

// ValidateStorageImagesCreateRequestBody runs the validations defined on
// Storage/Images/CreateRequestBody
func ValidateStorageImagesCreateRequestBody(body *StorageImagesCreateRequestBody) (err error) {
	if body.Volume == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("volume", "body"))
	}
	if body.Image == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image", "body"))
	}
	return
}

// ValidateStorageImagesDeleteRequestBody runs the validations defined on
// Storage/Images/DeleteRequestBody
func ValidateStorageImagesDeleteRequestBody(body *StorageImagesDeleteRequestBody) (err error) {
	if body.VolumeName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("volume_name", "body"))
	}
	if body.ImageName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image_name", "body"))
	}
	return
}

// ValidateStorageImagesGetRequestBody runs the validations defined on
// Storage/Images/GetRequestBody
func ValidateStorageImagesGetRequestBody(body *StorageImagesGetRequestBody) (err error) {
	if body.VolumeName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("volume_name", "body"))
	}
	if body.ImageName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image_name", "body"))
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
