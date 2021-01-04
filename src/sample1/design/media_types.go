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

// MediaType of Admin API.
var AdminUserStatsResponse = ResultType("application/vnd.goa3-sample.admin.user_number+json", func() {
	Description("statistic of users")
	ContentType("application/json")

	Attributes(func() {
		Attribute("data", ArrayOf(Data), "グラフデータ")
		Attribute("x", String, "X軸に使用するkey")
		Attribute("y", String, "Y軸に使用するkey")
		Attribute("size", String, "ドットの大きさに使用するkey")
		Attribute("color", String, "ドットの色分けに使用するkey")
		Attribute("guide", GuideType)

		Required("data", "x", "y", "guide")
	})
})

var AdminUserResponse = ResultType("application/vnd.goa3-sample.admin.user+json", func() {
	Description("User Response")
	ContentType("application/json")

	Reference(User)
	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
		Attribute("created_at", String)
		Attribute("updated_at", String)
		Attribute("deleted_at", String)

		Required("id", "name", "email", "created_at", "updated_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("deleted_at")
	})
})
