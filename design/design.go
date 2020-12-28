package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = API("spin", func() {
	Title("Spin API service")
	Description("HTTP service for controlling the spin service")
	Server("spin", func() {
		Host("localhost", func() { URI("http://localhost:8080") })
	})
})

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

var _ = Service("spin-broker", func() {
	Description("The message broker for the other services")

	Method("new", func() {
		Description("Create a new package; a collection of items to join into the queue simultaneously")

		Result(String)

		HTTP(func() {
			POST("/new")
			Response(StatusOK)
		})
	})

	Method("add", func() {
		Description("Add a command to the package")

		Payload(func() {
			Attribute("id", String, "Package ID")
			Attribute("resource", String, "Resource name")
			Attribute("action", String, "Action name")
			Attribute("parameters", MapOf(String, Any), "Action parameters")

			Required("id", "resource", "action")
		})

		Result(String)

		HTTP(func() {
			POST("/add/{id}")
			Response(StatusOK)
		})
	})

	Method("enqueue", func() {
		Description("Enqueue the package into the various resource queues")

		Payload(func() {
			Attribute("id", String, "Package ID")
			Required("id")
		})

		Result(ArrayOf(String))

		HTTP(func() {
			POST("/enqueue/{id}")
			Response(StatusOK)
		})
	})

	Method("status", func() {
		Description("Get the status for a package")

		Payload(func() {
			Attribute("id", String, "Package ID")
			Required("id")
		})

		Result(func() {
			Attribute("status", Boolean, "Pass/Fail status")
			Attribute("reason", String, "Failure reason (if any)")
			Required("status")
		})

		Error("record_not_found")

		HTTP(func() {
			GET("/status/{id}")
			Response(StatusOK)

			Response("record_not_found", StatusBadRequest, func() {
				// Use HTTP status code 400 (BadRequest) to write "div_by_zero" errors
				Description("No records were found")
			})
		})
	})

	Method("next", func() {
		Description("Get the next command for a given resource")

		Payload(func() {
			Attribute("resource", String, "resource type")
			Required("resource")
		})

		Result(func() {
			Attribute("uuid", String, "Command ID")
			Attribute("resource", String, "resource type")
			Attribute("action", String, "action name")
			Attribute("parameters", MapOf(String, Any), "parameters for action")
			Required("uuid", "resource", "action")
		})

		Error("record_not_found")

		HTTP(func() {
			GET("/next/{resource}")
			Response(StatusOK)
			Response("record_not_found", StatusBadRequest, func() {
				// Use HTTP status code 400 (BadRequest) to write "div_by_zero" errors
				Description("No records were found")
			})
		})
	})

	Method("complete", func() {
		Description("Mark a command as completed with a result status")

		Payload(func() {
			Attribute("id", String, "Command ID")
			Attribute("status", Boolean, "status of work")
			Attribute("status_reason", String, "reason of success/failure")
			Required("id", "status")
		})

		HTTP(func() {
			POST("/complete")
			Response(StatusOK)
		})
	})
})
