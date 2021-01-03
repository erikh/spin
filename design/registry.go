package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = Service("spin-registry", func() {
	Description("Keeper of the VMs")

	/*******************************************
	*
	* vm registry functions
	*
	********************************************/

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

	/*******************************************
	*
	* storage registry functions
	*
	********************************************/

	Method("storage/volumes/list", func() {
		Description("list all volumes")

		Result(ArrayOf(String))

		HTTP(func() {
			GET("/storage/volumes/list")
			Response(StatusOK)
		})
	})

	Method("storage/volumes/create", func() {
		Description("create a new volume")

		Payload(func() {
			Attribute("name", String, "name of volume")
			Required("name")
		})

		HTTP(func() {
			POST("/storage/volumes/create")
			Response(StatusOK)
		})
	})

	Method("storage/volumes/delete", func() {
		Description("delete an existing volume")

		Payload(func() {
			Attribute("name", String, "name of volume")
			Required("name")
		})

		HTTP(func() {
			POST("/storage/volumes/delete")
			Response(StatusOK)
		})
	})

	Method("storage/images/list", func() {
		Description("list all images by volume")

		Payload(func() {
			Attribute("volume_name", String, "name of volume to list images for")
			Required("volume_name")
		})

		Result(ArrayOf(String))

		HTTP(func() {
			GET("/storage/images/list")
			Response(StatusOK)
		})
	})

	Method("storage/images/create", func() {
		Description("add an image definition to the registry")

		Payload(Storage)

		HTTP(func() {
			POST("/storage/images/create")
			Response(StatusOK)
		})
	})

	Method("storage/images/delete", func() {
		Description("remove an image definition from the registry")

		Payload(func() {
			Attribute("volume_name", String, "name of volume")
			Attribute("image_name", String, "name of image")
			Required("volume_name", "image_name")
		})

		HTTP(func() {
			POST("/storage/images/delete")
			Response(StatusOK)
		})
	})

	Method("storage/images/get", func() {
		Description("retrieves an image definition from the registry")

		Payload(func() {
			Attribute("volume_name", String, "name of volume")
			Attribute("image_name", String, "name of image")
			Required("volume_name", "image_name")
		})

		Result(Storage)

		HTTP(func() {
			GET("/storage/images/get")
			Response(StatusOK)
		})
	})
})
