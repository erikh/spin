package vm

// Storage is for creating and referencing new storage. Existing storage
// references, with the exception of adding cdroms should consult with the
// Image struct instead.
type Storage struct {
	Volume    *string `json:"volume"`
	Image     string  `json:"image"`
	ImageSize *uint   `json:"image_size"`
	CDROM     bool    `json:"cdrom"`
}

// Core is all the values that always apply to VMs in any situation.
type Core struct {
	Name   string          `json:"name"`
	CPUs   uint            `json:"cpus"`
	Memory uint            `json:"memory"`
	Ports  map[uint]string `json:"ports"`
}

// Create is for creating new VMs. Storage definitons are created as necessary
// and turned into image definitions (see Updated type)
type Create struct {
	Core    `json:",inline"`
	Storage []Storage `json:"storage"`
}

// Image is an image that is already created.
type Image struct {
	Path   string  `json:"path"`
	CDROM  bool    `json:"cdrom"`
	Volume *string `json:"volume"`
}

// Transient is for in-situ records that are usually being transformed.
type Transient struct {
	Core   `json:",inline"`
	Images []Image `json:"images"`
}
