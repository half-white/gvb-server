package main

import (
	"fmt"
	"gvb_server/models"

	"github.com/fatih/structs"
)

type AdverRequest struct {
	models.MODEL `structs:"-"`
	Title        string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`      //显示的标题
	Href         string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"-"`      //跳转链接
	Images       string `json:"images" binding:"required,url" msg:"图片地址不正确" structs:"-"`   //图片
	IsShow       bool   `json:"is_show" binding:"required" msg:"请输入是否展示"structs:"is_show"` //是否展示
}

func main() {
	u1 := AdverRequest{
		Title:  "xx",
		Href:   "xxx",
		Images: "xxxx",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
