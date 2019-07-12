package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("service-parrot", func() {

	Description("The Parrot service serves a parrot payload to the cute animal service")
	Method("postparrot", func() {
		Payload(ParrotPayload)
		HTTP(func() {
			POST("/new-parrot")
			Body(Empty)
			Response(StatusCreated)
		})
	})

	Method("listaparrot", func() {
		HTTP(func() {
			Payload(func() {
				Attribute("id", String, "The ID of the Parrot", func() {
					Example("123")
					MinLength(1)
				})
				Required("id")
			})

			Error("not_found")
			Result(Parrot)

			HTTP(func() {
				GET("/{id}")
				Body(Empty)
				Response(StatusOK)
				Response("not_found", StatusNotFound)
			})
		})
	})

	Method("listallparrots", func() {
		HTTP(func() {
			Result(CollectionOf(Parrot))
			HTTP(func() {
				GET("/all-parrots")
				Body(Empty)
				Response(StatusOK)
			})
		})
	})
})
