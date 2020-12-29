package registry

type VM struct {
	Name    string
	CPUs    uint
	Memory  uint // in megabytes
	Volumes []Storage
}

type Storage struct {
	Volume   string
	Filename string
	Size     uint // in gigabytes
}
