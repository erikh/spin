// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-broker HTTP client CLI support package
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package client

import (
	"encoding/json"
	"fmt"

	spinbroker "code.hollensbe.org/erikh/spin/gen/spin_broker"
)

// BuildAddPayload builds the payload for the spin-broker add endpoint from CLI
// flags.
func BuildAddPayload(spinBrokerAddBody string, spinBrokerAddID string) (*spinbroker.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(spinBrokerAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"action\": \"Dolore saepe voluptatum dolorem rerum et.\",\n      \"parameters\": [\n         \"Maxime corrupti molestias est sunt ab repellat.\",\n         \"Quia eum temporibus.\",\n         \"Odit eligendi.\"\n      ],\n      \"resource\": \"Qui est officia laborum est esse similique.\"\n   }'")
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
		v.Parameters = make([]string, len(body.Parameters))
		for i, val := range body.Parameters {
			v.Parameters[i] = val
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Nobis consequatur omnis.\",\n      \"status\": false,\n      \"status_reason\": \"Voluptatem dignissimos ut doloremque harum et tempora.\"\n   }'")
		}
	}
	v := &spinbroker.CompletePayload{
		ID:           body.ID,
		Status:       body.Status,
		StatusReason: body.StatusReason,
	}

	return v, nil
}
