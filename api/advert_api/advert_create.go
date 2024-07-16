package advert_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type AdverRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`        //显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法"`    //跳转链接
	Image  string `json:"images" binding:"required,url" msg:"图片地址不正确"` //图片
	IsShow bool   `json:"is_show" binding:"required" msg:"请输入是否展示"`    //是否展示
}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdverRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Image:  cr.Image,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}

	res.OkWithMessage("添加广告成功", c)

}
