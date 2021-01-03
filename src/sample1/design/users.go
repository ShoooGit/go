package design

import (
	. "goa.design/goa/v3/dsl"
)

// サービスの定義
var _ = Service("Users", func() {
	Description("users serves user account relative information.")

	HTTP(func() {
		Path("/api/v1/users")
	})

	Method("list user", func() {
		Description("List all stored users")

		Result(CollectionOf(UserResponse))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("get user", func() {
		Description("Show user by ID")

		Payload(GetUserPayload)
		Result(UserResponse)

		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("create user", func() {
		Description("Add new user and return its ID.")

		Payload(CreateUserPayload)
		Result(String)

		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	Method("update user", func() {
		Description("Update user item.")
		Payload(UpdateUserPayload)
		Result(UserResponse)
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
	})

	Method("delete user", func() {
		Description("Delete user by id.")
		Payload(DeleteUserPayload)
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})
})
