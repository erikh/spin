package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = Service("spin-apiserver", func() {
	Description("Bridge between the outer-facing UIs and the internals")

	Method("vm_create", func() {
		Payload(Any, "vm type", func() {
			Meta("struct:field:type", "*vm.Create", "github.com/erikh/spin/pkg/vm")
		})
		Result(UInt64)

		HTTP(func() {
			POST("/vm/create")
			Response(StatusOK)
		})
	})

	Method("vm_delete", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to delete")
			Required("id")
		})

		HTTP(func() {
			POST("/vm/delete/{id}")
			Response(StatusOK)
		})
	})

	Method("vm_list", func() {
		Result(ArrayOf(UInt64))

		HTTP(func() {
			POST("/vm/list")
			Response(StatusOK)
		})
	})

	Method("vm_get", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to retrieve")
			Required("id")
		})
		Result(Any, "vm type", func() {
			Meta("struct:field:type", "*vm.Transient", "github.com/erikh/spin/pkg/vm")
		})

		HTTP(func() {
			GET("/vm/get/{id}")
			Response(StatusOK)
		})
	})

	Method("vm_update", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to Update")
			Attribute("vm", Any, "VM Manifest to Update", func() {
				Meta("struct:field:type", "*vm.Transient", "github.com/erikh/spin/pkg/vm")
			})
			Required("id", "vm")
		})

		HTTP(func() {
			POST("/vm/update/{id}")
			Response(StatusOK)
		})
	})

	Method("control_start", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to start")
			Required("id")
		})
		HTTP(func() {
			POST("/control/start/{id}")
			Response(StatusOK)
		})
	})

	Method("control_stop", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to stop")
			Required("id")
		})
		HTTP(func() {
			POST("/control/stop/{id}")
			Response(StatusOK)
		})
	})

	Method("control_shutdown", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to shutdown")
			Required("id")
		})
		HTTP(func() {
			POST("/control/shutdown/{id}")
			Response(StatusOK)
		})
	})
})
