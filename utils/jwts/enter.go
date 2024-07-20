package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad jwt中的payload数据
type JwtPayLoad struct {
	UserName string `json:"username"`  //用户名
	NickName string `json:"nick_name"` //昵称
	Role     int    `json:"role"`      //权限等级 1 管理员 2 普通用户 3 游客
	UserID   uint   `json:"user_id"`   //用户id
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

var MySecret []byte
