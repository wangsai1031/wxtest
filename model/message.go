package model

type MessageEntity struct {
	Id           int64  `gorm:"primary_key;column:id" json:"id"`
	ToUserName   string `gorm:"column:to_user_name" json:"to_user_name"`     // 开发者微信号
	FromUserName string `gorm:"column:from_user_name" json:"from_user_name"` // 发送方OpenID
	MsgType      string `gorm:"column:msg_type" json:"msg_type"`             // 消息类型
	Event        string `gorm:"column:event" json:"event"`                   // 时间类型
	Resource     string `gorm:"column:resource" json:"resource"`             // 消息原始内容
	CreateTime   int64  `gorm:"column:create_time" json:"create_time"`       // 消息创建时间
}

func (MessageEntity) TableName() string {
	return "message"
}
