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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"path\": \"Eum officiis eligendi.\",\n      \"volume\": \"Non deleniti consequuntur qui doloremque.\"\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Iusto minima et labore.\",\n      \"image_size\": 8231368540121560477,\n      \"volume\": \"Consequatur omnis dolor assumenda.\"\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Dignissimos ut doloremque.\",\n      \"volume\": \"Et tempora rem.\"\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Quia voluptas aut.\",\n      \"image_size\": 4442355973768769894,\n      \"volume\": \"Quisquam rem.\"\n   }'")
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"image_name\": \"Illum sint.\",\n      \"target_volume\": \"Quam neque.\",\n      \"volume\": \"Ullam odio tenetur aliquid consequatur.\"\n   }'")
		}
	}
	v := &spinapiserver.MoveImagePayload{
		ImageName:    body.ImageName,
		Volume:       body.Volume,
		TargetVolume: body.TargetVolume,
	}

	return v, nil
}
