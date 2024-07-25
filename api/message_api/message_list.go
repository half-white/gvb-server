package message_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var messageList []models.MessageModel
	global.DB.Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID)
	fmt.Println(messageList)
}
