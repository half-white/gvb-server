package settings_api

import "github.com/gin-gonic/gin"

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "xxx"})
}
