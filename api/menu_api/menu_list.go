package menu_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// 创建对外输出显示结构
type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	//先查询菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)

	//查询连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus []MenuResponse
	for _, model := range menuList {
		//model就是一个菜单
		//申明后需要开辟空间，解决nil值问题
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
	return
}
