package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("getchat", func() {
	Title("Chat Service")
	Description("Service for chat app, a Goa teaser")
	Server("chat api", func() {
		Host("172.25.0.3", func() {
			URI("http://172.25.0.3:8000")
		})
	})
})

var _ = Service("chatapi", func() {
	Description("The service performs get chat.")
	cors.Origin("http://172.25.0.2", func() {
		cors.Headers("Access-Control-Allow-Origin", "Authorization")
		cors.Methods("GET")
		cors.MaxAge(600)
		cors.Credentials()
	})
	Method("getchat", func() {
		Security(APIKeyAuth)
		Payload(func() {
			APIKey("api_key", "key", String, "API key used to perform authorization")
			Attribute("id", Int, func() {
				Description("room id")
			})
			Required("key", "id")
		})
		Result(CollectionOf(Chat))
		Error("NotFound")
		Error("BadRequest")
		HTTP(func() {
			GET("/chatroom/{id}")
			Header("key:Authorization")
			Response(StatusOK)
		})
	})
	Method("postchat", func() {
		Security(APIKeyAuth)
		Payload(func() {
			APIKey("api_key", "key", String, "API key used to perform authorization")
			Attribute("Id", String, "room id")
			Attribute("UserId", String, "user id")
			Attribute("RoomName", String, "room name")
			Attribute("Member", String, "member")
			Attribute("Chat", String, "chat")
			Attribute("PostDt", String, func() { Format(FormatDateTime) })
			Attribute("Cookie", String, "cookie")
			Required("key", "Id", "UserId", "RoomName", "Member", "Chat", "PostDt", "Cookie")
		})
		Result(Boolean)
		Error("NotFound")
		Error("BadRequest")
		HTTP(func() {
			POST("/chatroom/chat")
			Header("key:Authorization")
			Response(StatusOK)
		})
	})
})

var Chat = ResultType("application/vnd.goa.chat", func() {
	Description("All chat")
	Attributes(func() {
		Attribute("Id", Int, "room id")
		Attribute("UserId", String, "user id")
		Attribute("RoomName", String, "room name")
		Attribute("Member", String, "member")
		Attribute("Chat", String, "chat")
		Attribute("PostDt", String, func() { Format(FormatDateTime) })
		Required("Id", "UserId", "RoomName", "Member", "Chat", "PostDt")
	})
})

var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("Secures endpoint by requiring an API key.")
})
