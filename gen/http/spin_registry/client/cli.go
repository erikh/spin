// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry HTTP client CLI support package
//
// Command:
// $ goa gen github.com/erikh/spin/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	spinregistry "github.com/erikh/spin/gen/spin_registry"
	goa "goa.design/goa/v3/pkg"
)

// BuildVMUpdatePayload builds the payload for the spin-registry vm_update
// endpoint from CLI flags.
func BuildVMUpdatePayload(spinRegistryVMUpdateBody string, spinRegistryVMUpdateID string) (*spinregistry.UpdateVM, error) {
	var err error
	var body VMUpdateRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryVMUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"vm\": \"Nulla numquam rerum asperiores corporis aut enim.\"\n   }'")
		}
		if body.VM == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("vm", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id uint64
	{
		id, err = strconv.ParseUint(spinRegistryVMUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinregistry.UpdateVM{
		VM: body.VM,
	}
	v.ID = id

	return v, nil
}

// BuildVMDeletePayload builds the payload for the spin-registry vm_delete
// endpoint from CLI flags.
func BuildVMDeletePayload(spinRegistryVMDeleteID string) (*spinregistry.VMDeletePayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinRegistryVMDeleteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinregistry.VMDeletePayload{}
	v.ID = id

	return v, nil
}

// BuildVMGetPayload builds the payload for the spin-registry vm_get endpoint
// from CLI flags.
func BuildVMGetPayload(spinRegistryVMGetID string) (*spinregistry.VMGetPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinRegistryVMGetID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinregistry.VMGetPayload{}
	v.ID = id

	return v, nil
}

// BuildStorageVolumesCreatePayload builds the payload for the spin-registry
// storage_volumes_create endpoint from CLI flags.
func BuildStorageVolumesCreatePayload(spinRegistryStorageVolumesCreateBody string) (*spinregistry.StorageVolumesCreatePayload, error) {
	var err error
	var body StorageVolumesCreateRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryStorageVolumesCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Fugit incidunt tempora.\",\n      \"path\": \"Sed quae ea.\"\n   }'")
		}
	}
	v := &spinregistry.StorageVolumesCreatePayload{
		Name: body.Name,
		Path: body.Path,
	}

	return v, nil
}

// BuildStorageVolumesDeletePayload builds the payload for the spin-registry
// storage_volumes_delete endpoint from CLI flags.
func BuildStorageVolumesDeletePayload(spinRegistryStorageVolumesDeleteBody string) (*spinregistry.StorageVolumesDeletePayload, error) {
	var err error
	var body StorageVolumesDeleteRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryStorageVolumesDeleteBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Tempore voluptate voluptas.\"\n   }'")
		}
	}
	v := &spinregistry.StorageVolumesDeletePayload{
		Name: body.Name,
	}

	return v, nil
}

// BuildStorageImagesListPayload builds the payload for the spin-registry
// storage_images_list endpoint from CLI flags.
func BuildStorageImagesListPayload(spinRegistryStorageImagesListBody string) (*spinregistry.StorageImagesListPayload, error) {
	var err error
	var body StorageImagesListRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryStorageImagesListBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"volume_name\": \"Mollitia eius hic.\"\n   }'")
		}
	}
	v := &spinregistry.StorageImagesListPayload{
		VolumeName: body.VolumeName,
	}

	return v, nil
}

// BuildStorageImagesDeletePayload builds the payload for the spin-registry
// storage_images_delete endpoint from CLI flags.
func BuildStorageImagesDeletePayload(spinRegistryStorageImagesDeleteBody string) (*spinregistry.StorageImagesDeletePayload, error) {
	var err error
	var body StorageImagesDeleteRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryStorageImagesDeleteBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Nihil eius molestiae.\",\n      \"volume_name\": \"Assumenda soluta repellendus eaque neque vel accusamus.\"\n   }'")
		}
	}
	v := &spinregistry.StorageImagesDeletePayload{
		VolumeName: body.VolumeName,
		ImageName:  body.ImageName,
	}

	return v, nil
}

// BuildStorageImagesGetPayload builds the payload for the spin-registry
// storage_images_get endpoint from CLI flags.
func BuildStorageImagesGetPayload(spinRegistryStorageImagesGetBody string) (*spinregistry.StorageImagesGetPayload, error) {
	var err error
	var body StorageImagesGetRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryStorageImagesGetBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Ex natus dicta aliquid sint provident sint.\",\n      \"volume_name\": \"Consequatur nemo autem ab delectus amet.\"\n   }'")
		}
	}
	v := &spinregistry.StorageImagesGetPayload{
		VolumeName: body.VolumeName,
		ImageName:  body.ImageName,
	}

	return v, nil
}
