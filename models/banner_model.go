package models

import "gorm.io/gorm"

type BannerModel struct {
	MODEL
	Path string `json:"path"`               //图片路径
	Hash string `json:"hash"`               //图片的hash值，用于判断重复图片
	Name string `gorm:"size:38"json:"name"` //图片名称
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {

	return
}
