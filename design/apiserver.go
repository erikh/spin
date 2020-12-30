package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = Service("spin-apiserver", func() {
	Description("Bridge between the outer-facing UIs and the internals")

	Method("add_volume", func() {
		Description("Add a volume for image allocation with backing storage, and name it")

		Payload(func() {
			Attribute("volume", String, "volume identifier")
			Attribute("path", String, "path to backing storage")

			Required("volume", "path")
		})

		HTTP(func() {
			POST("/storage/volume/add")
			Response(StatusOK)
		})
	})

	Method("remove_volume", func() {
		Description("Remove a volume. Requires all images to be removed.")

		Payload(func() {
			Attribute("volume", String, "volume identifier")
			Required("volume")
		})

		HTTP(func() {
			POST("/storage/volume/remove/{volume}")
			Response(StatusOK)
		})
	})

	Method("label_volume", func() {
		Description("Apply a label to a volume.")

		Payload(func() {
			Attribute("volume", String, "volume identifier")
			Attribute("label", String, "label identifier to apply to volume")
			Required("volume", "label")
		})

		HTTP(func() {
			POST("/storage/volume/label/{volume}/{label}")
			Response(StatusOK)
		})
	})

	Method("info_volume", func() {
		Description("Get information on a volume")

		Payload(func() {
			Attribute("volume", String, "volume identifier")
			Required("volume")
		})

		HTTP(func() {
			GET("/storage/volume/info/{volume}")
			Response(StatusOK)
		})
	})

	Method("create_image_on_volume", func() {
		Description("Create an image on a volume")

		Payload(func() {
			Attribute("image_name", String, "image name")
			Attribute("image_size", UInt, "image size in GB")
			Attribute("volume", String, "volume identifier")

			Required("image_name", "image_size", "volume")
		})

		HTTP(func() {
			POST("/storage/volume/image/create")
			Response(StatusOK)
		})
	})

	Method("delete_image_on_volume", func() {
		Description("Delete an image on a volume")

		Payload(func() {
			Attribute("image_name", String, "image name")
			Attribute("volume", String, "volume identifier")

			Required("image_name", "volume")
		})

		HTTP(func() {
			POST("/storage/volume/image/delete")
			Response(StatusOK)
		})
	})

	Method("resize_image_on_volume", func() {
		Description("Resize an image on a volume")

		Payload(func() {
			Attribute("image_name", String, "image name")
			Attribute("image_size", UInt, "new image size; must be larger than original size")
			Attribute("volume", String, "volume identifier")

			Required("image_name", "image_size", "volume")
		})

		HTTP(func() {
			POST("/storage/volume/image/resize")
			Response(StatusOK)
		})
	})

	Method("info_image_on_volume", func() {
		Description("Obtain information on an image")

		Payload(func() {
			Attribute("image_name", String, "image name")
			Attribute("volume", String, "volume identifier")

			Required("image_name", "volume")
		})

		HTTP(func() {
			GET("/storage/volume/image/info/{volume}/{image_name}")
			Response(StatusOK)
		})
	})

	Method("move_image", func() {
		Description("Move an image from one volume to another")

		Payload(func() {
			Attribute("image_name", String, "image name")
			Attribute("volume", String, "volume identifier")
			Attribute("target_volume", String, "volume identifier to move to")

			Required("image_name", "volume", "target_volume")
		})

		HTTP(func() {
			POST("/storage/volume/image/move")
			Response(StatusOK)
		})
	})
})
