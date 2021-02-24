package design

import (
	. "goa.design/goa/v3/dsl"
)

// レスポンスデータの定義
// MediaType of Users API.
var UserResponse = ResultType("application/vnd.goa3-sample.user+json", func() {
	Description("User Response")
	ContentType("application/json")

	Reference(User)
	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")

		Required("id", "name", "email")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
	})
})

// MediaType of Ha API.
var HaResponse = ResultType("application/vnd.goa3-sample.ha+json", func() {
	Description("Ha Response")
	ContentType("application/json")

	Reference(Ha)
	Attributes(func() {
		Attribute("theme")
		Attribute("card")

		Required("theme", "card")
	})

	View("default", func() {
		Attribute("theme")
		Attribute("card")
	})
})
