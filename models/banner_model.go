package models

import (
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"

	"gorm.io/gorm"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        //图片路径
	Hash      string          `json:"hash"`                        //图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38"json:"name"`          //图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片存储类型（本地、七牛）
}

// 钩子函数
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//本地图片，删除，还要删除本地存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
