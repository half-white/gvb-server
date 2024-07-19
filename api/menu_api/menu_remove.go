package menu_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest	true "广告id列表"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error()
			return err
		}
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error()
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error()
		res.FailWithMessage("删除菜单失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除%d个菜单", count), c)
}
