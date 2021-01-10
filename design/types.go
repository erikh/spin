package design

import . "goa.design/goa/v3/dsl" // nolint

// UpdateVM is a type encapsulating a VM and an ID to update.
var UpdateVM = Type("UpdateVM", func() {
	Attribute("id", UInt64, "ID of VM to update")
	Attribute("vm", Any, "vm to update", func() {
		Meta("struct:field:type", "*vm.Transient", "github.com/erikh/spin/pkg/vm")
	})
	Required("id", "vm")
})
