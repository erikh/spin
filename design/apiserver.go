package design

import . "goa.design/goa/v3/dsl" // nolint

var _ = Service("spin-apiserver", func() {
	Description("Bridge between the outer-facing UIs and the internals")
})
