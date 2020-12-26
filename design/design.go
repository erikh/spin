package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = API("spin", func() {
	Title("Spin API service")
	Description("HTTP service for controlling the spin service")
	Server("spin", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var _ = Service("spin-broker", func() {
	Description("The message broker for the other services")

	Method("new", func() {
		Result(String)
		HTTP(func() {
			POST("/new")
			Response(StatusOK)
		})
	})

	Method("add", func() {
		Payload(func() {
			Attribute("id", String, "Package ID")
			Attribute("resource", String, "Resource name")
			Attribute("action", String, "Action name")
			Attribute("parameters", ArrayOf(String), "Action parameters")

			Required("id", "resource", "action")
		})
		Result(String)
		HTTP(func() {
			POST("/add/{id}")
			Response(StatusOK)
		})
	})

	Method("enqueue", func() {
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
		Payload(func() {
			Attribute("id", String, "Package ID")
			Required("id")
		})
		Result(func() {
			Attribute("status", Boolean, "Pass/Fail status")
			Attribute("reason", String, "Failure reason (if any)")
			Required("status")
		})
		HTTP(func() {
			GET("/status/{id}")
			Response(StatusOK)
		})
	})

	Method("next", func() {
		Payload(func() {
			Attribute("resource", String, "resource type")
			Required("resource")
		})
		Result(func() {
			Attribute("uuid", String, "Command ID")
			Attribute("resource", String, "resource type")
			Attribute("action", String, "action name")
			Attribute("parameters", ArrayOf(String), "parameters for action")
			Required("uuid", "resource", "action")
		})
		HTTP(func() {
			GET("/next/{resource}")
			Response(StatusOK)
		})
	})

	Method("complete", func() {
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
