package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"非法邮箱"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	//用户绑定邮箱，第一次是输入邮箱
	//后台会给这个邮箱发送验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	session := sessions.Default(c)
	if cr.Code == nil {
		//第一次，后台发验证码
		//生成4位验证码，将生成的验证码存入session
		code := random.Code(4)

		//写入session
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码是："+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}

	//第二次，用户输入邮箱、验证码和密码
	code := session.Get("valid_code")
	//检验验证码
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	//修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	if len(cr.Password) < 4 {
		res.FailWithMessage("用户密码强度太低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	//第一次的邮箱和第二次的邮箱也要做一致性校验
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":     cr.Email,
		"pass_word": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}

	//完成绑定
	res.OkWithMessage("邮箱绑定成功", c)
	return
}
