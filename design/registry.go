package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = Service("spin-registry", func() {
	Description("Keeper of the VMs")

	Method("vm/create", func() {
		Description("Create a VM")
		Payload(VM)
		Result(UInt64)

		HTTP(func() {
			POST("/vm/create")
			Response(StatusOK)
		})
	})

	Method("vm/update", func() {
		Description("Update a VM")
		Payload(UpdateVM)

		HTTP(func() {
			POST("/vm/update/{id}")
			Response(StatusOK)
		})
	})

	Method("vm/delete", func() {
		Description("Delete a VM by ID")
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to remove")
			Required("id")
		})

		HTTP(func() {
			POST("/vm/delete/{id}")
			Response(StatusOK)
		})
	})

	Method("vm/get", func() {
		Description("Retrieve a VM by ID")
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to remove")
			Required("id")
		})
		Result(VM)

		HTTP(func() {
			GET("/vm/get/{id}")
			Response(StatusOK)
		})
	})

	Method("vm/list", func() {
		Description("Retrieve all VM IDs")
		Result(ArrayOf(UInt64))

		HTTP(func() {
			GET("/vm/list")
			Response(StatusOK)
		})
	})
})
