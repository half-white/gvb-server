package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		UserName: "xxx",
		NickName: "xxx",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Inh4eCIsIm5pY2tfbmFtZSI6Inh4eCIsInJvbGUiOjEsInVzZXJfaWQiOjEsImV4cCI6MTcyMTQ3MTk3NS4zNTQ3ODgsImlzcyI6IjEyMzQifQ.Jgrn5n0c8A8g13bMmbFovyaonsR0Qu3x0nS34qoUZHQ")
	fmt.Println(claims, err)
}
