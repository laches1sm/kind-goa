package design

import . "goa.design/goa/v3/dsl"

var _ = API("doggo", func() {
	Title("Doggo Service")
	Description("HTTP service that serves doggos to a cute animal service")
	Server("doggo", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})
