package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi

	router.GET("settings", settingsApi.SettingsInfoView)
}
