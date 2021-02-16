package design

import (
	. "goa.design/goa/v3/dsl"
)

// サービスの定義
var _ = Service("Ha", func() {
	Description("ha serves.")

	HTTP(func() {
		Path("/api/v1/ha")
	})

	Method("draw card", func() {
		Description("decied on a theme and card")

		Result(CollectionOf(HaResponse))
		HTTP(func() {
			GET("/draw")
			Response(StatusOK)
		})
	})
})
