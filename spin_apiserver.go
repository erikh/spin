package spin

import (
	"context"
	"log"

	spinapiserver "code.hollensbe.org/erikh/spin/gen/spin_apiserver"
)

// spin-apiserver service example implementation.
// The example methods log the requests and return zero values.
type spinApiserversrvc struct {
	logger *log.Logger
}

// NewSpinApiserver returns the spin-apiserver service implementation.
func NewSpinApiserver(logger *log.Logger) spinapiserver.Service {
	return &spinApiserversrvc{logger}
}

// Add a volume for image allocation with backing storage, and name it
func (s *spinApiserversrvc) AddVolume(ctx context.Context, p *spinapiserver.AddVolumePayload) (err error) {
	s.logger.Print("spinApiserver.add_volume")
	return
}

// Remove a volume. Requires all images to be removed.
func (s *spinApiserversrvc) RemoveVolume(ctx context.Context, p *spinapiserver.RemoveVolumePayload) (err error) {
	s.logger.Print("spinApiserver.remove_volume")
	return
}

// Apply a label to a volume.
func (s *spinApiserversrvc) LabelVolume(ctx context.Context, p *spinapiserver.LabelVolumePayload) (err error) {
	s.logger.Print("spinApiserver.label_volume")
	return
}

// Get information on a volume
func (s *spinApiserversrvc) InfoVolume(ctx context.Context, p *spinapiserver.InfoVolumePayload) (err error) {
	s.logger.Print("spinApiserver.info_volume")
	return
}

// Create an image on a volume
func (s *spinApiserversrvc) CreateImageOnVolume(ctx context.Context, p *spinapiserver.CreateImageOnVolumePayload) (err error) {
	s.logger.Print("spinApiserver.create_image_on_volume")
	return
}

// Delete an image on a volume
func (s *spinApiserversrvc) DeleteImageOnVolume(ctx context.Context, p *spinapiserver.DeleteImageOnVolumePayload) (err error) {
	s.logger.Print("spinApiserver.delete_image_on_volume")
	return
}

// Resize an image on a volume
func (s *spinApiserversrvc) ResizeImageOnVolume(ctx context.Context, p *spinapiserver.ResizeImageOnVolumePayload) (err error) {
	s.logger.Print("spinApiserver.resize_image_on_volume")
	return
}

// Obtain information on an image
func (s *spinApiserversrvc) InfoImageOnVolume(ctx context.Context, p *spinapiserver.InfoImageOnVolumePayload) (err error) {
	s.logger.Print("spinApiserver.info_image_on_volume")
	return
}

// Move an image from one volume to another
func (s *spinApiserversrvc) MoveImage(ctx context.Context, p *spinapiserver.MoveImagePayload) (err error) {
	s.logger.Print("spinApiserver.move_image")
	return
}
