package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = Service("spin-apiserver", func() {
	Description("Bridge between the outer-facing UIs and the internals")

	Method("vm/create", func() {
		Payload(VM)
		Result(UInt64)

		HTTP(func() {
			POST("/vm/create")
			Response(StatusOK)
		})
	})

	Method("vm/delete", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to delete")
			Required("id")
		})

		HTTP(func() {
			POST("/vm/delete/{id}")
			Response(StatusOK)
		})
	})

	Method("control/start", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to start")
			Required("id")
		})
		HTTP(func() {
			POST("/control/start/{id}")
			Response(StatusOK)
		})
	})

	Method("control/stop", func() {
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to stop")
			Required("id")
		})
		HTTP(func() {
			POST("/control/stop/{id}")
			Response(StatusOK)
		})
	})

	Method("control/shutdown", func() {
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
