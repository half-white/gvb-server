package models

import (
	"gvb_server/models/ctype"
)

// AuthModel 用户表
type UserModel struct {
	//gorm.Model 需要逻辑删除时可用
	MODEL
	NickName       string           `gorm:"size:36" json:"nick_name"` //昵称
	UserName       string           `gorm:"size:36" json:"user_name"`
	PassWord       string           `gorm:"size:128" json:"-"`
	Avatar         string           `gorm:"sign:256" json:"avatar_id"`
	Email          string           `gorm:"sign:128" json:"email"`
	Tel            string           `gorm:"size:18" json:"tel"`
	Addr           string           `gorm:"size:64" json:"addr"`
	Token          string           `gorm:"size:64" json:"token"`
	IP             string           `gorm:"size:20" json:"ip"`
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"` //角色权限
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`                                                            //发布的文章
	CollectsModels []ArticleModel   `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"` //收藏的文章
}
