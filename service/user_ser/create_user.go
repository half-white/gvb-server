package user_ser

import (
	"errors"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/pwd"
)

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	//判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Select("user_name").Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}

	//对密码进行hash
	hashPwd := pwd.HashPwd(password)

	//头像问题
	//1.默认头像
	//2.随机选择头像
	avatar := "/uploads/avatar/default.png"

	//入库
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		PassWord:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         ip,
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}
	return nil
}
