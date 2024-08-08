package advert_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param data body models.RemoveRequest	true "广告id列表"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Debug().Select("id").Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.OkWithMessage(fmt.Sprintf("共删除%d条广告", count), c)
}
