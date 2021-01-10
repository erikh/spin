package registryclient

import (
	"context"
	"net/http"
	"time"

	"github.com/erikh/spin/gen/http/spin_registry/client"
	spinregistry "github.com/erikh/spin/gen/spin_registry"
	"github.com/erikh/spin/pkg/vm"
	goahttp "goa.design/goa/v3/http"
)

const registryProto = "http"

// Config is the configuration struct for the HTTP client. Timeout is in
// seconds.  Host should contain a port if not port 80.
type Config struct {
	Host    string
	Timeout int
}

// Client is the outer handle for the registry HTTP client.
type Client struct {
	client *client.Client
}

// New constructs a new http client from the configuration provided.
func New(cc Config) *Client {
	return &Client{client.NewClient(
		registryProto,
		cc.Host,
		&http.Client{Timeout: time.Duration(cc.Timeout) * time.Second},
		goahttp.RequestEncoder,
		goahttp.ResponseDecoder,
		false,
	)}
}

// VMCreate creates a new vm.
func (c *Client) VMCreate(ctx context.Context, vm *vm.Transient) (uint64, error) {
	if err := vm.Validate(); err != nil {
		return 0, err
	}

	pkg, err := c.client.VMCreate()(ctx, vm)
	if err != nil {
		return 0, err
	}

	return pkg.(uint64), nil
}

// VMUpdate updates a vm by id.
func (c *Client) VMUpdate(ctx context.Context, id uint64, vm *vm.Transient) error {
	if err := vm.Validate(); err != nil {
		return err
	}

	_, err := c.client.VMUpdate()(ctx, &spinregistry.UpdateVM{ID: id, VM: vm})
	return err
}

// VMDelete deletes a vm by id.
func (c *Client) VMDelete(ctx context.Context, id uint64) error {
	_, err := c.client.VMDelete()(ctx, &spinregistry.VMDeletePayload{ID: id})
	return err
}

// VMGet retrieves a vm by id.
func (c *Client) VMGet(ctx context.Context, id uint64) (*vm.Transient, error) {
	ret, err := c.client.VMGet()(ctx, &spinregistry.VMGetPayload{ID: id})
	if err != nil {
		return nil, err
	}

	return ret.(*vm.Transient), nil
}

// VMList retrieves all IDs of all VMs.
func (c *Client) VMList(ctx context.Context) ([]uint64, error) {
	res, err := c.client.VMList()(ctx, nil)
	if err != nil {
		return nil, err
	}

	return res.([]uint64), nil
}

// StorageVolumeCreate creates a volume by name.
func (c *Client) StorageVolumeCreate(ctx context.Context, name, path string) error {
	_, err := c.client.StorageVolumesCreate()(ctx, &spinregistry.StorageVolumesCreatePayload{Name: name, Path: path})
	return err
}

// StorageVolumeDelete deletes a volume by name.
func (c *Client) StorageVolumeDelete(ctx context.Context, name string) error {
	_, err := c.client.StorageVolumesDelete()(ctx, &spinregistry.StorageVolumesDeletePayload{Name: name})
	return err
}

// StorageVolumeList lists all volumes by name.
func (c *Client) StorageVolumeList(ctx context.Context) ([]string, error) {
	list, err := c.client.StorageVolumesList()(ctx, nil)
	if err != nil {
		return nil, err
	}

	return list.([]string), nil
}

// StorageImageList lists images for a given volume.
func (c *Client) StorageImageList(ctx context.Context, name string) ([]string, error) {
	list, err := c.client.StorageImagesList()(ctx, &spinregistry.StorageImagesListPayload{VolumeName: name})
	if err != nil {
		return nil, err
	}

	return list.([]string), nil
}

// StorageImageGet retrieves an image by name w/ volume name.
func (c *Client) StorageImageGet(ctx context.Context, volumeName, imageName string) (*vm.Storage, error) {
	res, err := c.client.StorageImagesGet()(ctx, &spinregistry.StorageImagesGetPayload{
		VolumeName: volumeName,
		ImageName:  imageName,
	})
	if err != nil {
		return nil, err
	}

	return res.(*vm.Storage), nil
}

// StorageImageCreate creates an image
func (c *Client) StorageImageCreate(ctx context.Context, s *vm.Storage) (*vm.Image, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}

	img, err := c.client.StorageImagesCreate()(ctx, s)
	if err != nil {
		return nil, err
	}

	return img.(*vm.Image), nil
}

// StorageImageDelete deletes an image by name & volume name
func (c *Client) StorageImageDelete(ctx context.Context, volumeName, imageName string) error {
	_, err := c.client.StorageImagesDelete()(ctx, &spinregistry.StorageImagesDeletePayload{
		VolumeName: volumeName,
		ImageName:  imageName,
	})
	return err
}
