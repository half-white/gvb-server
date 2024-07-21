package middleware

import (
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
		}
		//判断是否在redis中，在则已经注销
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已经失效", c)
			c.Abort()
			return
		}
		//登录的用户
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
		}
		//判断是否在redis中，在则已经注销
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("token已经失效", c)
			c.Abort()
			return
		}
		//登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
