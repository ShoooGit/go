package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

// API 定義
var _ = API("goa3-sample", func() {
	Title("SampleAPI")
	Description("goa3 sample code.")

	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("content-type")
		cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		cors.MaxAge(600)
	})

	Server("goa3-sample", func() {
		Services("Users", "Admin")
		Host("localhost", func() {
			Description("development host")
			URI("http://localhost:8080")
		})
	})
})
