package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id" structs:"-"` //主键
	CreatedAt time.Time `json:"created_at" structs:"-"`           //创建时间
	UpdatedAt time.Time `json:"-" structs:"-"`                    //更新时间
}

type RemoveRequest struct {
	IDList []uint `json:"id_list" binding:"required"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
