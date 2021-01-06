package design

import . "goa.design/goa/v3/dsl" // nolint

// Image encapsulates an existing image on disk.
var Image = Type("Image", func() {
	Attribute("path", String, "Image path")
	Attribute("cdrom", Boolean, "Is this a cdrom image?")
	Attribute("volume", String, "Volume name")
	Required("path", "cdrom")
})

// Storage encapuslates the properties around VM storage
var Storage = Type("Storage", func() {
	Attribute("volume", String, "Volume name; required if image is not a cdrom")
	Attribute("image", String, "Image filename, no `/` characters")
	Attribute("image_size", UInt64, "Image size (in gigabytes); required if image is not a cdrom")
	Attribute("cdrom", Boolean, "Is this image a cdrom?")
	Required("image", "cdrom")
})

// VMBase is attributes that live similarly between cycles in the VM's lifetime.
var VMBase = Type("VMBase", func() {
	Attribute("name", String, "Name of VM; does not need to be unique")
	Attribute("cpus", UInt, "CPU count")
	Attribute("memory", UInt, "Memory (in megabytes)")
})

// CreateVM is a standalone VM for creation purposes.
var CreateVM = Type("CreateVM", func() {
	Extend(VMBase)
	Attribute("storage", ArrayOf(Storage), "Storage references")
	Required("name", "cpus", "memory", "storage")
})

// UpdatedVM is a standalone VM for purposes of updating.
var UpdatedVM = Type("UpdatedVM", func() {
	Extend(VMBase)
	Attribute("images", ArrayOf(Image), "Image references")
	Required("name", "cpus", "memory", "images")
})

// UpdateVM is a type encapsulating a VM and an ID to update.
var UpdateVM = Type("UpdateVM", func() {
	Attribute("id", UInt64, "ID of VM to update")
	Attribute("vm", UpdatedVM, "VM to publish")
	Required("id", "vm")
})
