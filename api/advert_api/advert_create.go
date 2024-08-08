package advert_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type AdverRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`         //显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`      //跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址不正确" structs:"images"` //图片
	IsShow bool   `json:"is_show" msg:"请输入是否展示" structs:"is_show"`                      //是否展示
}

// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdverRequest	true "表示多个参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdverRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//重复判断机制
	var advert models.AdvertModel
	err = global.DB.Select("title").Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该广告已经存在", c)
		return
	}

	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}

	res.OkWithMessage("添加广告成功", c)

}
