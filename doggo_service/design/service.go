package design

import . "goa.design/goa/v3/dsl"

var _ = Service("doggo", func() {
	Description("The Doggo service serves a doggo payload to the cute animal service")
	Method("postdoggo", func() {
		Payload(DoggoPayload)
		HTTP(func() {
			POST("/new-doggo")
			Body(Empty)
			Response(StatusCreated)
		})
	})

	Method("listadoggo", func() {
		HTTP(func() {
			Payload(func() {
				Attribute("id", String, "The ID of the Doggo", func() {
					Example("123")
					MinLength(1)
				})
				Required("id")
			})

			Error("not_found")
			Result(Doggo)

			HTTP(func() {
				GET("/{id}")
				Body(Empty)
				Response(StatusOK)
				Response("not_found", StatusNotFound)
			})
		})
	})

	Method("listalldoggos", func() {
		HTTP(func() {
			Result(CollectionOf(Doggo))
			HTTP(func() {
				GET("/all-doggos")
				Body(Empty)
				Response(StatusOK)
			})
		})
	})
})
