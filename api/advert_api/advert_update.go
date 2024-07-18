package advert_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// @Tags 广告管理
// @Summary 更新广告
// @Description 更新广告
// @Param data body AdverRequest	true "查询参数"
// @Router /api/adverts/:id [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdverRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}

	err = global.DB.Model(&advert).Updates(structs.Map(&cr)).Error

	//结构体转map的第三方包

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}

	res.OkWithMessage("修改广告成功", c)
}
