package vm

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
)

// StringPtr returns the string provided as a pointer.
func StringPtr(s string) *string {
	return &s
}

// UIntPtr returns the uint provided as a pointer.
func UIntPtr(u uint) *uint {
	return &u
}

var nameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)

// Storage is for creating and referencing new storage. Existing storage
// references, with the exception of adding cdroms should consult with the
// Image struct instead.
type Storage struct {
	Volume    *string `json:"volume"`
	Image     string  `json:"image"`
	ImageSize *uint   `json:"image_size"`
	CDROM     bool    `json:"cdrom"`
}

// PortMap is a mapping of guest port -> hostaddr:port. hostaddr must be an IP
// address.
type PortMap map[uint16]string

// Core is all the values that always apply to VMs in any situation.
type Core struct {
	Name   string  `json:"name"`
	CPUs   uint    `json:"cpus"`
	Memory uint    `json:"memory"`
	Ports  PortMap `json:"ports"`
}

// Image is an image that is already created.
type Image struct {
	Path   string  `json:"path"`
	CDROM  bool    `json:"cdrom"`
	Volume *string `json:"volume"`
}

// Create is for creating new VMs. Storage definitons are created as necessary
// and turned into image definitions (see Updated type)
type Create struct {
	Core    `json:",inline"`
	Storage []Storage `json:"storage"`
}

// Transient is for in-situ records that are usually being transformed.
type Transient struct {
	Core   `json:",inline"`
	Images []Image `json:"images"`
}

// Validate the Core structure
func (c Core) Validate() error {
	if c.Memory == 0 {
		return errors.New("memory must be non-zero")
	}

	if c.CPUs == 0 {
		return errors.New("cpus must be non-zero")
	}

	if c.Name == "" {
		return errors.New("name must not be empty")
	}

	if !nameRegex.MatchString(c.Name) {
		return fmt.Errorf("name does not match required format: %q", nameRegex)
	}

	if c.Ports == nil {
		return errors.New("ports were empty")
	}

	return c.Ports.Validate()
}

// Validate the Create structure
func (c Create) Validate() error {
	if err := c.Core.Validate(); err != nil {
		return err
	}

	if c.Storage == nil {
		return errors.New("storage was empty")
	}

	for x, storage := range c.Storage {
		if err := storage.Validate(); err != nil {
			return fmt.Errorf("Storage index %d: %v", x, err)
		}
	}

	return nil
}

// Validate the Transient structure
func (t Transient) Validate() error {
	if err := t.Core.Validate(); err != nil {
		return err
	}

	if t.Images == nil {
		return errors.New("images were empty")
	}

	for x, image := range t.Images {
		if err := image.Validate(); err != nil {
			return fmt.Errorf("Image index %d: %v", x, err)
		}
	}

	return nil
}

// Validate the PortMap structure
func (p PortMap) Validate() error {
	for guestport, hostaddr := range p {
		host, port, err := net.SplitHostPort(hostaddr)
		if err != nil {
			return fmt.Errorf("Port forwarding for guest port %d: %v", guestport, err)
		}

		if net.ParseIP(host) == nil {
			return errors.New("host portion must be an IP address")
		}

		if port == "" {
			return errors.New("invalid port")
		}

		if _, err := strconv.ParseUint(port, 10, 16); err != nil {
			return err
		}
	}

	return nil
}

// Validate the Storage structure
func (s Storage) Validate() error {
	// TODO
	// additional validations around image paths, specifically WRT use of whitespace

	if s.Image == "" {
		return errors.New("path to image is empty")
	}

	if s.CDROM {
		if s.Volume != nil {
			return errors.New("volumes aren't used with cdroms")
		}

		if s.ImageSize != nil {
			return errors.New("image size isn't used with cdroms")
		}
	} else {
		if s.Volume == nil {
			return errors.New("volume is unset")
		}

		if *s.Volume == "" {
			return errors.New("volume is empty")
		}

		if s.ImageSize == nil {
			return errors.New("image size is unset")
		}

		if *s.ImageSize == 0 {
			return errors.New("image size is 0")
		}
	}

	return nil
}

// Validate the Image structure
func (i Image) Validate() error {
	// TODO
	// additional validations around image paths, specifically WRT use of whitespace

	if i.Path == "" {
		return errors.New("path is empty")
	}

	if i.CDROM {
		if i.Volume != nil {
			return errors.New("volumes aren't used with cdroms")
		}
	} else {
		if i.Volume == nil {
			return errors.New("volume is unset")
		}

		if *i.Volume == "" {
			return errors.New("volume is empty")
		}
	}

	return nil
}
