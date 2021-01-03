// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-registry HTTP client CLI support package
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	spinregistry "code.hollensbe.org/erikh/spin/gen/spin_registry"
	goa "goa.design/goa/v3/pkg"
)

// BuildVMCreatePayload builds the payload for the spin-registry vm/create
// endpoint from CLI flags.
func BuildVMCreatePayload(spinRegistryVMCreateBody string) (*spinregistry.VM, error) {
	var err error
	var body VMCreateRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryVMCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"cpus\": 6421225864845588134,\n      \"memory\": 15540618699971399982,\n      \"name\": \"Et dicta.\",\n      \"storage\": [\n         {\n            \"cdrom\": true,\n            \"image\": \"Labore voluptas perferendis ea iusto adipisci.\",\n            \"image_size\": 8351098711286704476,\n            \"volume\": \"Porro eius officiis.\"\n         },\n         {\n            \"cdrom\": true,\n            \"image\": \"Labore voluptas perferendis ea iusto adipisci.\",\n            \"image_size\": 8351098711286704476,\n            \"volume\": \"Porro eius officiis.\"\n         },\n         {\n            \"cdrom\": true,\n            \"image\": \"Labore voluptas perferendis ea iusto adipisci.\",\n            \"image_size\": 8351098711286704476,\n            \"volume\": \"Porro eius officiis.\"\n         },\n         {\n            \"cdrom\": true,\n            \"image\": \"Labore voluptas perferendis ea iusto adipisci.\",\n            \"image_size\": 8351098711286704476,\n            \"volume\": \"Porro eius officiis.\"\n         }\n      ]\n   }'")
		}
		if body.Storage == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("storage", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &spinregistry.VM{
		Name:   body.Name,
		Cpus:   body.Cpus,
		Memory: body.Memory,
	}
	if body.Storage != nil {
		v.Storage = make([]*spinregistry.Storage, len(body.Storage))
		for i, val := range body.Storage {
			v.Storage[i] = marshalStorageRequestBodyToSpinregistryStorage(val)
		}
	}

	return v, nil
}

// BuildVMUpdatePayload builds the payload for the spin-registry vm/update
// endpoint from CLI flags.
func BuildVMUpdatePayload(spinRegistryVMUpdateBody string, spinRegistryVMUpdateID string) (*spinregistry.UpdateVM, error) {
	var err error
	var body VMUpdateRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryVMUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"vm\": {\n         \"cpus\": 8415388083871949362,\n         \"memory\": 15948619849308310515,\n         \"name\": \"Et qui fugit quis dignissimos qui.\",\n         \"storage\": [\n            {\n               \"cdrom\": true,\n               \"image\": \"Labore voluptas perferendis ea iusto adipisci.\",\n               \"image_size\": 8351098711286704476,\n               \"volume\": \"Porro eius officiis.\"\n            },\n            {\n               \"cdrom\": true,\n               \"image\": \"Labore voluptas perferendis ea iusto adipisci.\",\n               \"image_size\": 8351098711286704476,\n               \"volume\": \"Porro eius officiis.\"\n            },\n            {\n               \"cdrom\": true,\n               \"image\": \"Labore voluptas perferendis ea iusto adipisci.\",\n               \"image_size\": 8351098711286704476,\n               \"volume\": \"Porro eius officiis.\"\n            }\n         ]\n      }\n   }'")
		}
		if body.VM == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("vm", "body"))
		}
		if body.VM != nil {
			if err2 := ValidateVMRequestBody(body.VM); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
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
	v := &spinregistry.UpdateVM{}
	if body.VM != nil {
		v.VM = marshalVMRequestBodyToSpinregistryVM(body.VM)
	}
	v.ID = id

	return v, nil
}

// BuildVMDeletePayload builds the payload for the spin-registry vm/delete
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

// BuildVMGetPayload builds the payload for the spin-registry vm/get endpoint
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
