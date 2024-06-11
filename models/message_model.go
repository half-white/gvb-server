package models

//记录消息
type MessageModel struct {
	MODEL
	SendUserID        uint      `gorm:"primaryKey" json:"send_user_id"` //发送人id
	SendUserModel     UserModel `gorm:"foreignKey" json:"-"`
	SendUserNicekname string    `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvaatar   string    `json:"send_user_avater"`

	RevvUserID       uint      `gorm:"primaryKey" json:"rev_user_id"` //接收人id
	RevUserModel     UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevUserNicekname string    `gorm:"size:42" json:"rev_user_nick_name"`
	RevUserAvaatar   string    `json:"send_user_avater"`
	IsRead           bool      `gorm:"default:false" json:"is_read"` //接收方是否查看
	Content          string    `json:"content"`                      //消息内容
}
