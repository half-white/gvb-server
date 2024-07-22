package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("ASHdkAJSKldAsjdas123sd"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), app.UserBindEmailView)
}
