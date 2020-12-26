// Code generated by goa v3.2.6, DO NOT EDIT.
//
// spin-apiserver HTTP client CLI support package
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package client

import (
	"encoding/json"
	"fmt"

	spinapiserver "code.hollensbe.org/erikh/spin/gen/spin_apiserver"
)

// BuildAddVolumePayload builds the payload for the spin-apiserver add_volume
// endpoint from CLI flags.
func BuildAddVolumePayload(spinApiserverAddVolumeBody string) (*spinapiserver.AddVolumePayload, error) {
	var err error
	var body AddVolumeRequestBody
	{
		err = json.Unmarshal([]byte(spinApiserverAddVolumeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"path\": \"Autem neque adipisci.\",\n      \"volume\": \"Quia soluta veniam adipisci quia.\"\n   }'")
		}
	}
	v := &spinapiserver.AddVolumePayload{
		Volume: body.Volume,
		Path:   body.Path,
	}

	return v, nil
}

// BuildRemoveVolumePayload builds the payload for the spin-apiserver
// remove_volume endpoint from CLI flags.
func BuildRemoveVolumePayload(spinApiserverRemoveVolumeVolume string) (*spinapiserver.RemoveVolumePayload, error) {
	var volume string
	{
		volume = spinApiserverRemoveVolumeVolume
	}
	v := &spinapiserver.RemoveVolumePayload{}
	v.Volume = volume

	return v, nil
}

// BuildLabelVolumePayload builds the payload for the spin-apiserver
// label_volume endpoint from CLI flags.
func BuildLabelVolumePayload(spinApiserverLabelVolumeVolume string, spinApiserverLabelVolumeLabel string) (*spinapiserver.LabelVolumePayload, error) {
	var volume string
	{
		volume = spinApiserverLabelVolumeVolume
	}
	var label string
	{
		label = spinApiserverLabelVolumeLabel
	}
	v := &spinapiserver.LabelVolumePayload{}
	v.Volume = volume
	v.Label = label

	return v, nil
}

// BuildInfoVolumePayload builds the payload for the spin-apiserver info_volume
// endpoint from CLI flags.
func BuildInfoVolumePayload(spinApiserverInfoVolumeVolume string) (*spinapiserver.InfoVolumePayload, error) {
	var volume string
	{
		volume = spinApiserverInfoVolumeVolume
	}
	v := &spinapiserver.InfoVolumePayload{}
	v.Volume = volume

	return v, nil
}

// BuildCreateImageOnVolumePayload builds the payload for the spin-apiserver
// create_image_on_volume endpoint from CLI flags.
func BuildCreateImageOnVolumePayload(spinApiserverCreateImageOnVolumeBody string) (*spinapiserver.CreateImageOnVolumePayload, error) {
	var err error
	var body CreateImageOnVolumeRequestBody
	{
		err = json.Unmarshal([]byte(spinApiserverCreateImageOnVolumeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Et et rerum ut voluptatem at fugit.\",\n      \"image_size\": 15419653682127034888,\n      \"volume\": \"Dolor veritatis quo non quae rerum officia.\"\n   }'")
		}
	}
	v := &spinapiserver.CreateImageOnVolumePayload{
		ImageName: body.ImageName,
		ImageSize: body.ImageSize,
		Volume:    body.Volume,
	}

	return v, nil
}

// BuildDeleteImageOnVolumePayload builds the payload for the spin-apiserver
// delete_image_on_volume endpoint from CLI flags.
func BuildDeleteImageOnVolumePayload(spinApiserverDeleteImageOnVolumeBody string) (*spinapiserver.DeleteImageOnVolumePayload, error) {
	var err error
	var body DeleteImageOnVolumeRequestBody
	{
		err = json.Unmarshal([]byte(spinApiserverDeleteImageOnVolumeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Minima et labore qui nobis.\",\n      \"volume\": \"Omnis dolor assumenda voluptatem dignissimos ut.\"\n   }'")
		}
	}
	v := &spinapiserver.DeleteImageOnVolumePayload{
		ImageName: body.ImageName,
		Volume:    body.Volume,
	}

	return v, nil
}

// BuildResizeImageOnVolumePayload builds the payload for the spin-apiserver
// resize_image_on_volume endpoint from CLI flags.
func BuildResizeImageOnVolumePayload(spinApiserverResizeImageOnVolumeBody string) (*spinapiserver.ResizeImageOnVolumePayload, error) {
	var err error
	var body ResizeImageOnVolumeRequestBody
	{
		err = json.Unmarshal([]byte(spinApiserverResizeImageOnVolumeBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Harum et tempora rem qui quia.\",\n      \"image_size\": 7533879699541094989,\n      \"volume\": \"Non dicta quisquam rem qui.\"\n   }'")
		}
	}
	v := &spinapiserver.ResizeImageOnVolumePayload{
		ImageName: body.ImageName,
		ImageSize: body.ImageSize,
		Volume:    body.Volume,
	}

	return v, nil
}

// BuildInfoImageOnVolumePayload builds the payload for the spin-apiserver
// info_image_on_volume endpoint from CLI flags.
func BuildInfoImageOnVolumePayload(spinApiserverInfoImageOnVolumeVolume string, spinApiserverInfoImageOnVolumeImageName string) (*spinapiserver.InfoImageOnVolumePayload, error) {
	var volume string
	{
		volume = spinApiserverInfoImageOnVolumeVolume
	}
	var imageName string
	{
		imageName = spinApiserverInfoImageOnVolumeImageName
	}
	v := &spinapiserver.InfoImageOnVolumePayload{}
	v.Volume = volume
	v.ImageName = imageName

	return v, nil
}

// BuildMoveImagePayload builds the payload for the spin-apiserver move_image
// endpoint from CLI flags.
func BuildMoveImagePayload(spinApiserverMoveImageBody string) (*spinapiserver.MoveImagePayload, error) {
	var err error
	var body MoveImageRequestBody
	{
		err = json.Unmarshal([]byte(spinApiserverMoveImageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Ullam odio tenetur aliquid consequatur.\",\n      \"target_volume\": \"Id et autem ut debitis.\",\n      \"volume\": \"Quam neque.\"\n   }'")
		}
	}
	v := &spinapiserver.MoveImagePayload{
		ImageName:    body.ImageName,
		Volume:       body.Volume,
		TargetVolume: body.TargetVolume,
	}

	return v, nil
}
