package settings_api

import (
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(map[string]string{
		"id": "xxx",
	}, c)
}
