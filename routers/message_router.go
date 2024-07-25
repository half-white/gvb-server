package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.POST("messages", middleware.JwtAuth(), app.MessageCreateView)
	router.GET("messages_all", middleware.JwtAdmin(), app.MessageListAllView)
	router.GET("messgaes", middleware.JwtAuth(), app.MessageListView)
}
