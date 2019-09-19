package design

import (
	. "goa.design/goa/v3/dsl"
)

var Parrot = ResultType("parrot", func() {
	Description(`It's a parrot.`)

	Attribute("id", String, "ID for the parrot", func() {
		Example("party-1234")
	})
	Attribute("name", String, "the name of the parrot", func() {
		Example("Party Parrot")
	})
	Attribute("colour", String, "the colour of the parrot", func() {
		Example("green")
	})
	Attribute("special_trick", String, "the special trick of the parrot", func() {
		Example("party")
	})
	Required("id", "name", "colour", "special_trick")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("colour")
		Attribute("special_trick")
	})
})

var ParrotPayload = Type("parrotpayload", func() {
	Description(`It's a parrot payload.`)

	Attribute("id", String, "ID for the parrot", func() {
		Example("party-1234")
	})
	Attribute("name", String, "the name of the parrot", func() {
		Example("Party Parrot")
	})
	Attribute("colour", String, "the colour of the parrot", func() {
		Example("green")
	})
	Attribute("special_trick", String, "the special trick of the parrot", func() {
		Example("party")
	})
	Required("id", "name", "colour", "special_trick")
})
