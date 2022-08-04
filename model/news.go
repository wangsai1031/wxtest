package model

type NewsEntity struct {
	Id         int64  `gorm:"primary_key;column:id" json:"id"`
	ArticleIds string `gorm:"column:article_ids" json:"article_ids"`
	MediaId    string `gorm:"column:media_id" json:"media_id"`       // 图文上传ID，长度不固定，但不会超过 128 字符
	MsgId      int64  `gorm:"column:msg_id" json:"msg_id"`           // 消息发送任务的ID
	MsgDataId  int64  `gorm:"column:msg_data_id" json:"msg_data_id"` // 消息的数据ID，用于在图文分析数据接口中
	BizStatus  int64  `gorm:"column:biz_status" json:"biz_status"`   // 状态：0默认，3上传图文消息素材，4发布中，5发布完成
	UploadTime int64  `gorm:"column:upload_time" json:"upload_time"` // 上传时间
	SendTime   int64  `gorm:"column:send_time" json:"send_time"`     // 群发消息时间
	ErrMsg     string `gorm:"column:err_msg" json:"err_msg"`         // 错误信息
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` // 创建时间
	ModifyTime int64  `gorm:"column:modify_time" json:"modify_time"` // 修改时间
}

func (NewsEntity) TableName() string {
	return "news"
}
