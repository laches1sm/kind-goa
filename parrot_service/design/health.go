package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Health", func() {

	Error("service_unavailable")
	Method("check", func() {
		Description("Health check")
		HTTP(func() {
			GET("//health-check")
			Body(Empty)
			Response(StatusOK)
			Response(StatusServiceUnavailable)
		})
	})

})
