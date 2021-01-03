package design

import . "goa.design/goa/v3/dsl" // nolint

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
			Attribute("causer", String, "Failure causer as UUID (if any)")
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
			Attribute("parameters", MapOf(String, Any, func() {
				Elem(func() {
					Meta("struct:field:type", "json.RawMessage", "encoding/json")
				})
			}), "Action parameters")
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
