package design

import . "goa.design/goa/v3/dsl" // nolint

var Storage = Type("Storage", func() {
	Attribute("volume", String, "Volume name, must not include `/`")
	Attribute("image", String, "Image filename, must not include `/`")
	Attribute("image_size", UInt, "Image size (in gigabytes)")
	Required("volume", "image", "image_size")
})

var VM = Type("VM", func() {
	Attribute("name", String, "Name of VM; does not need to be unique")
	Attribute("cpus", UInt, "CPU count")
	Attribute("memory", UInt, "Memory (in megabytes)")
	Attribute("storage", ArrayOf(Storage), "Storage references")
	Required("name", "cpus", "memory", "storage")
})

var UpdateVM = Type("UpdateVM", func() {
	Attribute("id", UInt64, "ID of VM to update")
	Attribute("vm", VM, "VM to publish")
	Required("id", "vm")
})
