package user_api

import (
	"gvb_server/models/ctype"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"` //昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"`
	PassWord string     `json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"` //角色权限
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

}
