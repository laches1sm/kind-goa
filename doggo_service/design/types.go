package design

import . "goa.design/goa/v3/dsl"

var Doggo = ResultType("types.doggos", func() {
	TypeName("Doggo")
	Attributes(func() {
		Attribute("id", String, "ID of a particular doggo")
		Attribute("name", String, "Name of a doggo", func() {
			Example("Rosie")
		})
		Attribute("colour", String, "Colour of a doggo", func() {
			Example("red")
		})
		Attribute("special_trick", String, "The special move that a doggo can do", func() {
			Example("buttwiggle")
		})
		Attribute("breed", String, "The breed of a doggo", func() {
			Example("shibe")
		})
		Attribute("good_dog", Bool, "ALL DOGS ARE GOOD DOGS")
	})
})

var DoggoPayload = Type("types.doggo_payload", func() {
	TypeName("DoggoPayload")
	Attributes(func() {
		Attribute("name", String, "Name of a doggo", func() {
			Example("Rosie")
		})
		Attribute("colour", String, "Colour of a doggo", func() {
			Example("red")
		})
		Attribute("special_trick", String, "The special move that a doggo can do", func() {
			Example("buttwiggle")
		})
		Attribute("breed", String, "The breed of a doggo", func() {
			Example("shibe")
		})
		Attribute("good_dog", Bool, "ALL DOGS ARE GOOD DOGS")
	})
})
