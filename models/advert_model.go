package models

//广告表
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32" json:"title"` //显示的标题
	Href   string `json:"href"`                 //跳转链接
	Image  string `json:"images"`               //图片
	IsShow bool   `json:"is_show"`              //是否展示
}
