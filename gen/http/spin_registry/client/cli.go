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

// BuildCreatePayload builds the payload for the spin-registry create endpoint
// from CLI flags.
func BuildCreatePayload(spinRegistryCreateBody string) (*spinregistry.VM, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"cpus\": 9896149697376370528,\n      \"memory\": 12073838662949931649,\n      \"name\": \"Nihil officia aspernatur.\",\n      \"storage\": [\n         {\n            \"image\": \"In aut vero amet.\",\n            \"image_size\": 4230286153146778503,\n            \"volume\": \"Voluptatibus est minima cum delectus occaecati.\"\n         },\n         {\n            \"image\": \"In aut vero amet.\",\n            \"image_size\": 4230286153146778503,\n            \"volume\": \"Voluptatibus est minima cum delectus occaecati.\"\n         },\n         {\n            \"image\": \"In aut vero amet.\",\n            \"image_size\": 4230286153146778503,\n            \"volume\": \"Voluptatibus est minima cum delectus occaecati.\"\n         },\n         {\n            \"image\": \"In aut vero amet.\",\n            \"image_size\": 4230286153146778503,\n            \"volume\": \"Voluptatibus est minima cum delectus occaecati.\"\n         }\n      ]\n   }'")
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

// BuildUpdatePayload builds the payload for the spin-registry update endpoint
// from CLI flags.
func BuildUpdatePayload(spinRegistryUpdateBody string, spinRegistryUpdateID string) (*spinregistry.UpdateVM, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(spinRegistryUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"vm\": {\n         \"cpus\": 6421225864845588134,\n         \"memory\": 15540618699971399982,\n         \"name\": \"Et dicta.\",\n         \"storage\": [\n            {\n               \"image\": \"In aut vero amet.\",\n               \"image_size\": 4230286153146778503,\n               \"volume\": \"Voluptatibus est minima cum delectus occaecati.\"\n            },\n            {\n               \"image\": \"In aut vero amet.\",\n               \"image_size\": 4230286153146778503,\n               \"volume\": \"Voluptatibus est minima cum delectus occaecati.\"\n            },\n            {\n               \"image\": \"In aut vero amet.\",\n               \"image_size\": 4230286153146778503,\n               \"volume\": \"Voluptatibus est minima cum delectus occaecati.\"\n            },\n            {\n               \"image\": \"In aut vero amet.\",\n               \"image_size\": 4230286153146778503,\n               \"volume\": \"Voluptatibus est minima cum delectus occaecati.\"\n            }\n         ]\n      }\n   }'")
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
		id, err = strconv.ParseUint(spinRegistryUpdateID, 10, 64)
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

// BuildDeletePayload builds the payload for the spin-registry delete endpoint
// from CLI flags.
func BuildDeletePayload(spinRegistryDeleteID string) (*spinregistry.DeletePayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinRegistryDeleteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinregistry.DeletePayload{}
	v.ID = id

	return v, nil
}

// BuildGetPayload builds the payload for the spin-registry get endpoint from
// CLI flags.
func BuildGetPayload(spinRegistryGetID string) (*spinregistry.GetPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(spinRegistryGetID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &spinregistry.GetPayload{}
	v.ID = id

	return v, nil
}
