package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = API("spin", func() {
	Title("Spin API service")
	Description("HTTP service for controlling the spin service")
	Server("spin", func() {
		Host("localhost", func() { URI("http://localhost:8080") })
	})
})
