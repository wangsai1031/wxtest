package model

type ArticleMediaEntity struct {
	Id         int64  `gorm:"primary_key;column:id" json:"id"`
	ArticleId  int64  `gorm:"column:article_id" json:"article_id"`   // 文章ID
	Source     string `gorm:"column:source" json:"source"`           // 原始图片
	WeixinUrl  string `gorm:"column:weixin_url" json:"weixin_url"`   // 图片链接
	MediaId    string `gorm:"column:media_id" json:"media_id"`       // 图片上传ID，不一定都有
	BizStatus  int64  `gorm:"column:biz_status" json:"biz_status"`   // 状态：0默认，1上传成功，-1上传失败
	ErrMsg     string `gorm:"column:err_msg" json:"err_msg"`         // 失败信息
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` // 创建时间
	ModifyTime int64  `gorm:"column:modify_time" json:"modify_time"` // 修改时间
}

func (ArticleMediaEntity) TableName() string {
	return "article_media"
}
