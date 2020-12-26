// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the spin-apiserver service.
//
// Command:
// $ goa gen code.hollensbe.org/erikh/spin/design

package server

import (
	"fmt"
)

// AddVolumeSpinApiserverPath returns the URL path to the spin-apiserver service add_volume HTTP endpoint.
func AddVolumeSpinApiserverPath() string {
	return "/storage/volume/add"
}

// RemoveVolumeSpinApiserverPath returns the URL path to the spin-apiserver service remove_volume HTTP endpoint.
func RemoveVolumeSpinApiserverPath(volume string) string {
	return fmt.Sprintf("/storage/volume/remove/%v", volume)
}

// LabelVolumeSpinApiserverPath returns the URL path to the spin-apiserver service label_volume HTTP endpoint.
func LabelVolumeSpinApiserverPath(volume string, label string) string {
	return fmt.Sprintf("/storage/volume/label/%v/%v", volume, label)
}

// InfoVolumeSpinApiserverPath returns the URL path to the spin-apiserver service info_volume HTTP endpoint.
func InfoVolumeSpinApiserverPath(volume string) string {
	return fmt.Sprintf("/storage/volume/info/%v", volume)
}

// CreateImageOnVolumeSpinApiserverPath returns the URL path to the spin-apiserver service create_image_on_volume HTTP endpoint.
func CreateImageOnVolumeSpinApiserverPath() string {
	return "/storage/volume/image/create"
}

// DeleteImageOnVolumeSpinApiserverPath returns the URL path to the spin-apiserver service delete_image_on_volume HTTP endpoint.
func DeleteImageOnVolumeSpinApiserverPath() string {
	return "/storage/volume/image/delete"
}

// ResizeImageOnVolumeSpinApiserverPath returns the URL path to the spin-apiserver service resize_image_on_volume HTTP endpoint.
func ResizeImageOnVolumeSpinApiserverPath() string {
	return "/storage/volume/image/resize"
}

// InfoImageOnVolumeSpinApiserverPath returns the URL path to the spin-apiserver service info_image_on_volume HTTP endpoint.
func InfoImageOnVolumeSpinApiserverPath(volume string, imageName string) string {
	return fmt.Sprintf("/storage/volume/image/info/%v/%v", volume, imageName)
}

// MoveImageSpinApiserverPath returns the URL path to the spin-apiserver service move_image HTTP endpoint.
func MoveImageSpinApiserverPath() string {
	return "/storage/volume/image/move"
}
