package design

import (
	. "goa.design/goa/v3/dsl"
)

var Ha = Type("Ha", func() {
	Description("Ha describes a ha account information.")

	Attribute("theme", String, "theme of game", func() {
		Example("Ha")
	})
	Attribute("card", String, "card of abc", func() {
		Example("A")
	})

	Required("theme", "card")
})
