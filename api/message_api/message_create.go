package message_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` //发送人ID
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  //接收人ID
	Content    string `json:"content" binding:"required"`      //消息内容
}

// 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	//当前用户发送消息
	//SendUserId就是当前登陆人ID
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var sendUser, revUser models.UserModel
	err = global.DB.Select("id").Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Select("id").Take(&revUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接收人不存在", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:        cr.SendUserID,
		SendUserNicekName: sendUser.NickName,
		SendUserAvatar:    sendUser.Avatar,
		RevUserID:         cr.RevUserID,
		RevUserNicekName:  revUser.NickName,
		RevUserAvatar:     revUser.Avatar,
		IsRead:            false,
		Content:           cr.Content,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("消息发送失败", c)
		return
	}
	res.OkWithMessage("消息发送成功", c)
	return
}
