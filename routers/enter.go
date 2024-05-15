package routers

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)

	router := gin.Default()

	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//系统配置api
	routerGroupApp.SettingsRouter()

	return router
}
