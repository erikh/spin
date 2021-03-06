// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker HTTP client CLI support package
//
// Command:
// $ goa gen github.com/erikh/spin/design

package client

import (
	"encoding/json"
	"fmt"

	spinbroker "github.com/erikh/spin/gen/spin_broker"
)

// BuildAddPayload builds the payload for the spin-broker add endpoint from CLI
// flags.
func BuildAddPayload(spinBrokerAddBody string, spinBrokerAddID string) (*spinbroker.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(spinBrokerAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"action\": \"Ratione et ex tempore ipsum in.\",\n      \"dependencies\": [\n         \"Itaque tenetur labore quia distinctio aperiam.\",\n         \"Soluta voluptatem minima natus.\",\n         \"Consequatur voluptatem.\",\n         \"Eos iste illum omnis suscipit.\"\n      ],\n      \"parameters\": {\n         \"Doloribus impedit.\": \"Qui dolor ex consequatur.\",\n         \"Maiores ut voluptatem sed et consectetur.\": \"Repellat vero.\",\n         \"Ut hic assumenda assumenda ullam qui.\": \"Quidem qui.\"\n      },\n      \"resource\": \"Voluptatibus nostrum commodi error omnis quis quia.\"\n   }'")
		}
	}
	var id string
	{
		id = spinBrokerAddID
	}
	v := &spinbroker.AddPayload{
		Resource: body.Resource,
		Action:   body.Action,
	}
	if body.Parameters != nil {
		v.Parameters = make(map[string]interface{}, len(body.Parameters))
		for key, val := range body.Parameters {
			tk := key
			tv := val
			v.Parameters[tk] = tv
		}
	}
	if body.Dependencies != nil {
		v.Dependencies = make([]string, len(body.Dependencies))
		for i, val := range body.Dependencies {
			v.Dependencies[i] = val
		}
	}
	v.ID = id

	return v, nil
}

// BuildEnqueuePayload builds the payload for the spin-broker enqueue endpoint
// from CLI flags.
func BuildEnqueuePayload(spinBrokerEnqueueID string) (*spinbroker.EnqueuePayload, error) {
	var id string
	{
		id = spinBrokerEnqueueID
	}
	v := &spinbroker.EnqueuePayload{}
	v.ID = id

	return v, nil
}

// BuildStatusPayload builds the payload for the spin-broker status endpoint
// from CLI flags.
func BuildStatusPayload(spinBrokerStatusID string) (*spinbroker.StatusPayload, error) {
	var id string
	{
		id = spinBrokerStatusID
	}
	v := &spinbroker.StatusPayload{}
	v.ID = id

	return v, nil
}

// BuildNextPayload builds the payload for the spin-broker next endpoint from
// CLI flags.
func BuildNextPayload(spinBrokerNextResource string) (*spinbroker.NextPayload, error) {
	var resource string
	{
		resource = spinBrokerNextResource
	}
	v := &spinbroker.NextPayload{}
	v.Resource = resource

	return v, nil
}

// BuildCompletePayload builds the payload for the spin-broker complete
// endpoint from CLI flags.
func BuildCompletePayload(spinBrokerCompleteBody string) (*spinbroker.CompletePayload, error) {
	var err error
	var body CompleteRequestBody
	{
		err = json.Unmarshal([]byte(spinBrokerCompleteBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Similique corrupti ab soluta vel qui vel.\",\n      \"status\": true,\n      \"status_reason\": \"Nihil autem dolorem soluta quisquam.\"\n   }'")
		}
	}
	v := &spinbroker.CompletePayload{
		ID:           body.ID,
		Status:       body.Status,
		StatusReason: body.StatusReason,
	}

	return v, nil
}
