package message_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	SendUserID        uint      `json:"send_user_id"` //发送人id
	SendUserNicekName string    `json:"send_user_nick_name"`
	SendUserAvatar    string    `json:"send_user_avater"`
	RevUserID         uint      `json:"rev_user_id"` //接收人id
	RevUserNicekName  string    `json:"rev_user_nick_name"`
	RevUserAvatar     string    `json:"rev_user_avater"`
	Content           string    `json:"content"`       //消息内容
	CreatedAt         time.Time `json:"created_at"`    //最新消息时间
	MessageCount      int       `json:"message_count"` //消息数量
}

type MessageGroup map[uint]*Message

func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages []Message

	global.DB.Order("created_at asc").Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range messageList {
		//判断是一个组的条件
		//send_user_id 和 rev_user_id 其中一个
		// 1 2 和 2 1；1 3 和 3 1 是一组
		message := Message{
			SendUserID:        model.SendUserID,
			SendUserNicekName: model.SendUserNicekName,
			SendUserAvatar:    model.SendUserAvatar,
			RevUserID:         model.RevUserID,
			RevUserNicekName:  model.RevUserNicekName,
			RevUserAvatar:     model.RevUserAvatar,
			Content:           model.Content,
			CreatedAt:         model.CreatedAt,
			MessageCount:      1,
		}
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			//不存在
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}

	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	res.OkWithData(messages, c)
	return
}
