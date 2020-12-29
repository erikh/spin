package design

import . "goa.design/goa/v3/dsl" // nolint

var Storage = Type("Storage", func() {
	Attribute("volume", String, "Volume name, must not include `/`")
	Attribute("image", String, "Image filename, must not include `/`")
	Attribute("image_size", UInt, "Image size (in gigabytes)")
	Required("volume", "image", "image_size")
})

var vmAttrs = func() {
	Attribute("name", String, "Name of VM; does not need to be unique")
	Attribute("cpus", UInt, "CPU count")
	Attribute("memory", UInt, "Memory (in megabytes)")
	Attribute("storage", ArrayOf(Storage), "Storage references")
}

var VM = Type("VM", func() {
	vmAttrs()
	Required("name", "cpus", "memory", "storage")
})

var UpdateVM = Type("UpdateVM", func() {
	Attribute("id", UInt64, "ID of VM to update")
	vmAttrs()
	Required("id", "name", "cpus", "memory", "storage")
})
