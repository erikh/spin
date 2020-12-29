package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = Service("spin-registry", func() {
	Description("Keeper of the VMs")

	Method("create", func() {
		Description("Create a VM")
		Payload(VM)
		Result(UInt64)

		HTTP(func() {
			POST("/create")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Description("Update a VM")
		Payload(UpdateVM)

		HTTP(func() {
			POST("/update/{id}")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Description("Delete a VM by ID")
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to remove")
		})

		HTTP(func() {
			POST("/delete/{id}")
			Response(StatusOK)
		})
	})

	Method("get", func() {
		Description("Retrieve a VM by ID")
		Payload(func() {
			Attribute("id", UInt64, "ID of VM to remove")
		})
		Result(VM)

		HTTP(func() {
			GET("/get/{id}")
			Response(StatusOK)
		})
	})

	Method("list", func() {
		Description("Retrieve all VM IDs")
		Result(ArrayOf(UInt64))

		HTTP(func() {
			GET("/list")
			Response(StatusOK)
		})
	})
})
