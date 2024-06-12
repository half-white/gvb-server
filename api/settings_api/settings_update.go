package settings_api

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Println("before", global.Config)
	global.Config.SiteInfo = cr
	fmt.Println("after", global.Config)
}
