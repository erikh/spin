// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry HTTP client types
//
// Command:
// $ goa gen github.com/erikh/spin/design

package client

import (
	spinregistry "github.com/erikh/spin/gen/spin_registry"
	goa "goa.design/goa/v3/pkg"
)

// VMCreateRequestBody is the type of the "spin-registry" service "vm_create"
// endpoint HTTP request body.
type VMCreateRequestBody struct {
	// Image references
	Images []*ImageRequestBody `form:"images" json:"images" xml:"images"`
	// Name of VM; does not need to be unique
	Name string `form:"name" json:"name" xml:"name"`
	// CPU count
	Cpus uint `form:"cpus" json:"cpus" xml:"cpus"`
	// Memory (in megabytes)
	Memory uint `form:"memory" json:"memory" xml:"memory"`
}

// VMUpdateRequestBody is the type of the "spin-registry" service "vm_update"
// endpoint HTTP request body.
type VMUpdateRequestBody struct {
	// VM to publish
	VM *UpdatedVMRequestBody `form:"vm" json:"vm" xml:"vm"`
}

// StorageVolumesCreateRequestBody is the type of the "spin-registry" service
// "storage_volumes_create" endpoint HTTP request body.
type StorageVolumesCreateRequestBody struct {
	// name of volume
	Name string `form:"name" json:"name" xml:"name"`
	// path to volume
	Path string `form:"path" json:"path" xml:"path"`
}

// StorageVolumesDeleteRequestBody is the type of the "spin-registry" service
// "storage_volumes_delete" endpoint HTTP request body.
type StorageVolumesDeleteRequestBody struct {
	// name of volume
	Name string `form:"name" json:"name" xml:"name"`
}

// StorageImagesListRequestBody is the type of the "spin-registry" service
// "storage_images_list" endpoint HTTP request body.
type StorageImagesListRequestBody struct {
	// name of volume to list images for
	VolumeName string `form:"volume_name" json:"volume_name" xml:"volume_name"`
}

// StorageImagesCreateRequestBody is the type of the "spin-registry" service
// "storage_images_create" endpoint HTTP request body.
type StorageImagesCreateRequestBody struct {
	// Volume name; required if image is not a cdrom
	Volume *string `form:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
	// Image filename, no `/` characters
	Image string `form:"image" json:"image" xml:"image"`
	// Image size (in gigabytes); required if image is not a cdrom
	ImageSize *uint `form:"image_size,omitempty" json:"image_size,omitempty" xml:"image_size,omitempty"`
	// Is this image a cdrom?
	Cdrom bool `form:"cdrom" json:"cdrom" xml:"cdrom"`
}

// StorageImagesDeleteRequestBody is the type of the "spin-registry" service
// "storage_images_delete" endpoint HTTP request body.
type StorageImagesDeleteRequestBody struct {
	// name of volume
	VolumeName string `form:"volume_name" json:"volume_name" xml:"volume_name"`
	// name of image
	ImageName string `form:"image_name" json:"image_name" xml:"image_name"`
}

// StorageImagesGetRequestBody is the type of the "spin-registry" service
// "storage_images_get" endpoint HTTP request body.
type StorageImagesGetRequestBody struct {
	// name of volume
	VolumeName string `form:"volume_name" json:"volume_name" xml:"volume_name"`
	// name of image
	ImageName string `form:"image_name" json:"image_name" xml:"image_name"`
}

// VMGetResponseBody is the type of the "spin-registry" service "vm_get"
// endpoint HTTP response body.
type VMGetResponseBody struct {
	// Image references
	Images []*ImageResponseBody `form:"images,omitempty" json:"images,omitempty" xml:"images,omitempty"`
	// Name of VM; does not need to be unique
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// CPU count
	Cpus *uint `form:"cpus,omitempty" json:"cpus,omitempty" xml:"cpus,omitempty"`
	// Memory (in megabytes)
	Memory *uint `form:"memory,omitempty" json:"memory,omitempty" xml:"memory,omitempty"`
}

// StorageImagesCreateResponseBody is the type of the "spin-registry" service
// "storage_images_create" endpoint HTTP response body.
type StorageImagesCreateResponseBody struct {
	// Image path
	Path *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// Is this a cdrom image?
	Cdrom *bool `form:"cdrom,omitempty" json:"cdrom,omitempty" xml:"cdrom,omitempty"`
	// Volume name
	Volume *string `form:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
}

// StorageImagesGetResponseBody is the type of the "spin-registry" service
// "storage_images_get" endpoint HTTP response body.
type StorageImagesGetResponseBody struct {
	// Image path
	Path *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// Is this a cdrom image?
	Cdrom *bool `form:"cdrom,omitempty" json:"cdrom,omitempty" xml:"cdrom,omitempty"`
	// Volume name
	Volume *string `form:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
}

// ImageRequestBody is used to define fields on request body types.
type ImageRequestBody struct {
	// Image path
	Path string `form:"path" json:"path" xml:"path"`
	// Is this a cdrom image?
	Cdrom bool `form:"cdrom" json:"cdrom" xml:"cdrom"`
	// Volume name
	Volume *string `form:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
}

// UpdatedVMRequestBody is used to define fields on request body types.
type UpdatedVMRequestBody struct {
	// Image references
	Images []*ImageRequestBody `form:"images" json:"images" xml:"images"`
	// Name of VM; does not need to be unique
	Name string `form:"name" json:"name" xml:"name"`
	// CPU count
	Cpus uint `form:"cpus" json:"cpus" xml:"cpus"`
	// Memory (in megabytes)
	Memory uint `form:"memory" json:"memory" xml:"memory"`
}

// ImageResponseBody is used to define fields on response body types.
type ImageResponseBody struct {
	// Image path
	Path *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	// Is this a cdrom image?
	Cdrom *bool `form:"cdrom,omitempty" json:"cdrom,omitempty" xml:"cdrom,omitempty"`
	// Volume name
	Volume *string `form:"volume,omitempty" json:"volume,omitempty" xml:"volume,omitempty"`
}

// NewVMCreateRequestBody builds the HTTP request body from the payload of the
// "vm_create" endpoint of the "spin-registry" service.
func NewVMCreateRequestBody(p *spinregistry.UpdatedVM) *VMCreateRequestBody {
	body := &VMCreateRequestBody{
		Name:   p.Name,
		Cpus:   p.Cpus,
		Memory: p.Memory,
	}
	if p.Images != nil {
		body.Images = make([]*ImageRequestBody, len(p.Images))
		for i, val := range p.Images {
			body.Images[i] = marshalSpinregistryImageToImageRequestBody(val)
		}
	}
	return body
}

// NewVMUpdateRequestBody builds the HTTP request body from the payload of the
// "vm_update" endpoint of the "spin-registry" service.
func NewVMUpdateRequestBody(p *spinregistry.UpdateVM) *VMUpdateRequestBody {
	body := &VMUpdateRequestBody{}
	if p.VM != nil {
		body.VM = marshalSpinregistryUpdatedVMToUpdatedVMRequestBody(p.VM)
	}
	return body
}

// NewStorageVolumesCreateRequestBody builds the HTTP request body from the
// payload of the "storage_volumes_create" endpoint of the "spin-registry"
// service.
func NewStorageVolumesCreateRequestBody(p *spinregistry.StorageVolumesCreatePayload) *StorageVolumesCreateRequestBody {
	body := &StorageVolumesCreateRequestBody{
		Name: p.Name,
		Path: p.Path,
	}
	return body
}

// NewStorageVolumesDeleteRequestBody builds the HTTP request body from the
// payload of the "storage_volumes_delete" endpoint of the "spin-registry"
// service.
func NewStorageVolumesDeleteRequestBody(p *spinregistry.StorageVolumesDeletePayload) *StorageVolumesDeleteRequestBody {
	body := &StorageVolumesDeleteRequestBody{
		Name: p.Name,
	}
	return body
}

// NewStorageImagesListRequestBody builds the HTTP request body from the
// payload of the "storage_images_list" endpoint of the "spin-registry" service.
func NewStorageImagesListRequestBody(p *spinregistry.StorageImagesListPayload) *StorageImagesListRequestBody {
	body := &StorageImagesListRequestBody{
		VolumeName: p.VolumeName,
	}
	return body
}

// NewStorageImagesCreateRequestBody builds the HTTP request body from the
// payload of the "storage_images_create" endpoint of the "spin-registry"
// service.
func NewStorageImagesCreateRequestBody(p *spinregistry.Storage) *StorageImagesCreateRequestBody {
	body := &StorageImagesCreateRequestBody{
		Volume:    p.Volume,
		Image:     p.Image,
		ImageSize: p.ImageSize,
		Cdrom:     p.Cdrom,
	}
	return body
}

// NewStorageImagesDeleteRequestBody builds the HTTP request body from the
// payload of the "storage_images_delete" endpoint of the "spin-registry"
// service.
func NewStorageImagesDeleteRequestBody(p *spinregistry.StorageImagesDeletePayload) *StorageImagesDeleteRequestBody {
	body := &StorageImagesDeleteRequestBody{
		VolumeName: p.VolumeName,
		ImageName:  p.ImageName,
	}
	return body
}

// NewStorageImagesGetRequestBody builds the HTTP request body from the payload
// of the "storage_images_get" endpoint of the "spin-registry" service.
func NewStorageImagesGetRequestBody(p *spinregistry.StorageImagesGetPayload) *StorageImagesGetRequestBody {
	body := &StorageImagesGetRequestBody{
		VolumeName: p.VolumeName,
		ImageName:  p.ImageName,
	}
	return body
}

// NewVMGetUpdatedVMOK builds a "spin-registry" service "vm_get" endpoint
// result from a HTTP "OK" response.
func NewVMGetUpdatedVMOK(body *VMGetResponseBody) *spinregistry.UpdatedVM {
	v := &spinregistry.UpdatedVM{
		Name:   *body.Name,
		Cpus:   *body.Cpus,
		Memory: *body.Memory,
	}
	v.Images = make([]*spinregistry.Image, len(body.Images))
	for i, val := range body.Images {
		v.Images[i] = unmarshalImageResponseBodyToSpinregistryImage(val)
	}

	return v
}

// NewStorageImagesCreateImageOK builds a "spin-registry" service
// "storage_images_create" endpoint result from a HTTP "OK" response.
func NewStorageImagesCreateImageOK(body *StorageImagesCreateResponseBody) *spinregistry.Image {
	v := &spinregistry.Image{
		Path:   *body.Path,
		Cdrom:  *body.Cdrom,
		Volume: body.Volume,
	}

	return v
}

// NewStorageImagesGetImageOK builds a "spin-registry" service
// "storage_images_get" endpoint result from a HTTP "OK" response.
func NewStorageImagesGetImageOK(body *StorageImagesGetResponseBody) *spinregistry.Image {
	v := &spinregistry.Image{
		Path:   *body.Path,
		Cdrom:  *body.Cdrom,
		Volume: body.Volume,
	}

	return v
}

// ValidateVMGetResponseBody runs the validations defined on
// vm_get_response_body
func ValidateVMGetResponseBody(body *VMGetResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Cpus == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cpus", "body"))
	}
	if body.Memory == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("memory", "body"))
	}
	if body.Images == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("images", "body"))
	}
	for _, e := range body.Images {
		if e != nil {
			if err2 := ValidateImageResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStorageImagesCreateResponseBody runs the validations defined on
// storage_images_create_response_body
func ValidateStorageImagesCreateResponseBody(body *StorageImagesCreateResponseBody) (err error) {
	if body.Path == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("path", "body"))
	}
	if body.Cdrom == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cdrom", "body"))
	}
	return
}

// ValidateStorageImagesGetResponseBody runs the validations defined on
// storage_images_get_response_body
func ValidateStorageImagesGetResponseBody(body *StorageImagesGetResponseBody) (err error) {
	if body.Path == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("path", "body"))
	}
	if body.Cdrom == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cdrom", "body"))
	}
	return
}

// ValidateUpdatedVMRequestBody runs the validations defined on
// UpdatedVMRequestBody
func ValidateUpdatedVMRequestBody(body *UpdatedVMRequestBody) (err error) {
	if body.Images == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("images", "body"))
	}
	return
}

// ValidateImageResponseBody runs the validations defined on ImageResponseBody
func ValidateImageResponseBody(body *ImageResponseBody) (err error) {
	if body.Path == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("path", "body"))
	}
	if body.Cdrom == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cdrom", "body"))
	}
	return
}
