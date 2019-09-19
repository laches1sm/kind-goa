package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("parrot", func() {
	Title("Service-Parrot")
	Description("HTTP service that serves parrots to a cute animal service")
	Server("parrot", func() {
		Host("localhost", func() {
			URI("http://localhost:8088")
		})
	})
})


