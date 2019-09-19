package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("service-parrot", func() {

	Description("The Parrot service serves a parrot payload")

	Error("not_found")
	Method("postparrot", func() {
		Payload(ParrotPayload)
		HTTP(func() {
			POST("/new-parrot")
			Body(Empty)
			Response(StatusCreated)
		})
	})

	Method("listaparrot", func() {
		Payload(func() {
			Attribute("id", String, "The ID of the Parrot", func() {
				Example("123")
				MinLength(1)
			})
			Required("id")
		})

		HTTP(func() {
			GET("/one-parrot")
		})

	})

	Method("listallparrots", func() {
		HTTP(func() {
			GET("/all-parrots")
			Response(StatusOK)
		})
	})

})
