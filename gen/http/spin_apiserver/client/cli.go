// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver HTTP client CLI support package
//
// Command:
// $ goa gen github.com/erikh/spin/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	spinapiserver "github.com/erikh/spin/gen/spin_apiserver"
	goa "goa.design/goa/v3/pkg"
)

// BuildVMCreatePayload builds the payload for the spin-apiserver vm_create
// endpoint from CLI flags.
func BuildVMCreatePayload(spinApiserverVMCreateBody string) (*spinapiserver.CreateVM, error) {
	var err error
	var body VMCreateRequestBody
	{
		err = json.Unmarshal([]byte(spinApiserverVMCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"cpus\": 17092856392700368829,\n      \"memory\": 497222928895641055,\n      \"name\": \"Aspernatur vel vel illum voluptatem voluptatibus est.\",\n      \"storage\": [\n         {\n            \"cdrom\": true,\n            \"image\": \"Esse reprehenderit qui molestias eum voluptatem.\",\n            \"image_size\": 11533824901793082147,\n            \"volume\": \"Occaecati deserunt qui praesentium.\"\n         },\n         {\n            \"cdrom\": true,\n            \"image\": \"Esse reprehenderit qui molestias eum voluptatem.\",\n            \"image_size\": 11533824901793082147,\n            \"volume\": \"Occaecati deserunt qui praesentium.\"\n         }\n      ]\n   }'")
		}
		if body.Storage == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("storage", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &spinapiserver.CreateVM{
		Name:   body.Name,
		Cpus:   body.Cpus,
		Memory: body.Memory,
	}
	if body.Storage != nil {
		v.Storage = make([]*spinapiserver.Storage, len(body.Storage))
		for i, val := range body.Storage {
			v.Storage[i] = marshalStorageRequestBodyToSpinapiserverStorage(val)
		}
	}

	return v, nil
}

// BuildVMDeletePayload builds the payload for the spin-apiserver vm_delete
// endpoint from CLI flags.
func BuildVMDeletePayload(spinApiserverVMDeleteID string) (*spinapiserver.VMDeletePayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinApiserverVMDeleteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinapiserver.VMDeletePayload{}
	v.ID = id

	return v, nil
}

// BuildVMGetPayload builds the payload for the spin-apiserver vm_get endpoint
// from CLI flags.
func BuildVMGetPayload(spinApiserverVMGetID string) (*spinapiserver.VMGetPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinApiserverVMGetID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinapiserver.VMGetPayload{}
	v.ID = id

	return v, nil
}

// BuildVMUpdatePayload builds the payload for the spin-apiserver vm_update
// endpoint from CLI flags.
func BuildVMUpdatePayload(spinApiserverVMUpdateBody string, spinApiserverVMUpdateID string) (*spinapiserver.VMUpdatePayload, error) {
	var err error
	var body VMUpdateRequestBody
	{
		err = json.Unmarshal([]byte(spinApiserverVMUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"vm\": {\n         \"cpus\": 3472743333644681302,\n         \"images\": [\n            {\n               \"cdrom\": true,\n               \"path\": \"Dignissimos qui error modi.\",\n               \"volume\": \"Corrupti et voluptatibus et et occaecati.\"\n            },\n            {\n               \"cdrom\": true,\n               \"path\": \"Dignissimos qui error modi.\",\n               \"volume\": \"Corrupti et voluptatibus et et occaecati.\"\n            },\n            {\n               \"cdrom\": true,\n               \"path\": \"Dignissimos qui error modi.\",\n               \"volume\": \"Corrupti et voluptatibus et et occaecati.\"\n            }\n         ],\n         \"memory\": 639108202290023137,\n         \"name\": \"Quo dolore soluta consectetur.\"\n      }\n   }'")
		}
		if body.VM == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("vm", "body"))
		}
		if body.VM != nil {
			if err2 := ValidateUpdatedVMRequestBody(body.VM); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id uint64
	{
		id, err = strconv.ParseUint(spinApiserverVMUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinapiserver.VMUpdatePayload{}
	if body.VM != nil {
		v.VM = marshalUpdatedVMRequestBodyToSpinapiserverUpdatedVM(body.VM)
	}
	v.ID = id

	return v, nil
}

// BuildControlStartPayload builds the payload for the spin-apiserver
// control_start endpoint from CLI flags.
func BuildControlStartPayload(spinApiserverControlStartID string) (*spinapiserver.ControlStartPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinApiserverControlStartID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinapiserver.ControlStartPayload{}
	v.ID = id

	return v, nil
}

// BuildControlStopPayload builds the payload for the spin-apiserver
// control_stop endpoint from CLI flags.
func BuildControlStopPayload(spinApiserverControlStopID string) (*spinapiserver.ControlStopPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinApiserverControlStopID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinapiserver.ControlStopPayload{}
	v.ID = id

	return v, nil
}

// BuildControlShutdownPayload builds the payload for the spin-apiserver
// control_shutdown endpoint from CLI flags.
func BuildControlShutdownPayload(spinApiserverControlShutdownID string) (*spinapiserver.ControlShutdownPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinApiserverControlShutdownID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinapiserver.ControlShutdownPayload{}
	v.ID = id

	return v, nil
}
